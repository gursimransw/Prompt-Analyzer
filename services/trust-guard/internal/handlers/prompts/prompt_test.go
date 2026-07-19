package prompts

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gursimransw/BearBreach/services/trust-guard/internal/types"
)

func testDetectionRules() []types.DetectionRule {
	return []types.DetectionRule{
		{
			ID:          "SD-015",
			Name:        "Generic API Key Assignment",
			Category:    "sensitive_data",
			FindingType: "api_key",
			AppliesTo:   []string{"input", "output"},
			Pattern:     "(?i)(api[_\\-]?key|apikey|api[_\\-]?token|access[_\\-]?key|secret[_\\-]?key)\\s*[:=]\\s*[\"']?[A-Za-z0-9\\-_]{20,80}[\"']?",
			Weight:      0.92,
			Severity:    "critical",
			Reason:      "Detected generic API key or secret key assignment",
		},
	}
}

func testPolicyConfig() types.PolicyConfig {
	return types.PolicyConfig{
		SeverityConfig: []types.SeverityConfig{
			{
				Severity: "LOW",
				MinScore: 0.00,
				MaxScore: 0.30,
			},
			{
				Severity: "MEDIUM",
				MinScore: 0.31,
				MaxScore: 0.60,
			},
			{
				Severity: "HIGH",
				MinScore: 0.61,
				MaxScore: 0.89,
			},
			{
				Severity: "CRITICAL",
				MinScore: 0.90,
				MaxScore: 1.00,
			},
		},
		DefaultActionsConfig: map[string]string{
			"LOW":      "ALLOW",
			"MEDIUM":   "LOG",
			"HIGH":     "ALERT",
			"CRITICAL": "BLOCK",
		},
	}
}

type promptAnalyzerResponse struct {
	Matched     bool            `json:"matched"`
	Rules       []string        `json:"rules"`
	RiskScore   float64         `json:"riskScore"`
	Severity    string          `json:"severity"`
	Verdict     string          `json:"verdict"`
	Categories  []string        `json:"categories"`
	Reasons     []string        `json:"reasons"`
	RequestId   string          `json:"requestId"`
	ScanContext string          `json:"scanContext"`
	Findings    []types.Finding `json:"findings"`
}

func TestPromptAnalyzerMaliciousPrompt(t *testing.T) {
	rules := testDetectionRules()
	policyConfig := testPolicyConfig()

	handler := PromptAnalyzer(&rules, &policyConfig)

	requestBody := []byte(`{
		"content": "api_key = 'sk-test-1234567890abcdef' password := 'Security2024!!!'"
	}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/analyze-input",
		bytes.NewReader(requestBody),
	)

	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status code 200, got %d", recorder.Code)
	}

	var response promptAnalyzerResponse

	err := json.NewDecoder(recorder.Body).Decode(&response)
	if err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if !response.Matched {
		t.Fatal("expected matched=true, got matched=false")
	}

	if len(response.Rules) != 1 {
		t.Fatalf("expected 1 matched rule, got %d", len(response.Rules))
	}

	if response.Rules[0] != "SD-015" {
		t.Fatalf("expected matched rule SD-015, got %s", response.Rules[0])
	}

	if response.RiskScore != 0.92 {
		t.Fatalf("expected risk score 0.92, got %f", response.RiskScore)
	}

	if response.Severity != "CRITICAL" {
		t.Fatalf("expected severity CRITICAL, got %s", response.Severity)
	}

	if response.Verdict != "BLOCK" {
		t.Fatalf("expected verdict BLOCK, got %s", response.Verdict)
	}

	if len(response.Categories) != 1 || response.Categories[0] != "sensitive_data" {
		t.Fatalf("expected category prompt_injection, got %+v", response.Categories)
	}

	if len(response.Reasons) != 1 {
		t.Fatalf("expected 1 reason, got %d", len(response.Reasons))
	}

	if len(response.RequestId) != 36 {

		t.Fatalf("Invalid or no requestId was generated, got %s", response.RequestId)

	}

	if len(response.Findings) <= 0 {
		t.Fatalf("expected atleast one findings, got %d", len(response.Findings))

	}

	if response.Findings[0].FindingType != "api_key" {
		t.Fatalf("expected finding type = api_key, got %s", response.Findings[0].FindingType)

	}
}
