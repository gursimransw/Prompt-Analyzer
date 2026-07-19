package rules

import (
	"testing"

	"github.com/gursimransw/BearBreach/services/trust-guard/internal/types"
)

func TestGetRuleByID(t *testing.T) {

	rule1 := types.DetectionRule{

		ID:          "SD-002",
		Name:        "Generic API Key Assignment",
		Category:    "sensitive_data",
		FindingType: "api_key",
		AppliesTo:   []string{"input", "output"},
		Pattern:     `(?i)(api[_-]?key)\s*(?::=|[:=])\s*["']?[^\s"']{8,}["']?`,
		Weight:      0.30,
		Severity:    "critical",
		Reason:      "Detected generic API key or secret key assignment",
	}
	rule2 := types.DetectionRule{

		ID:          "SD-001",
		Name:        "Test Rule",
		Category:    "sensitive_data",
		FindingType: "api_key",
		AppliesTo:   []string{"input", "output"},
		Pattern:     `Testing Strin, no logic needed`,
		Weight:      0.86,
		Severity:    "critical",
		Reason:      "Detected generic API key or secret key assignment",
	}

	testRuleID := "SD-001"

	testRulesList := []types.DetectionRule{rule1, rule2}

	rule, error := GetRuleByID(testRuleID, testRulesList)

	if error != nil {
		t.Errorf("Test ran into an error %s", error.Error())
	}

	if rule.Name != "Test Rule" {

		t.Errorf("Expected rule name to be - Test Rule, got %s", rule.Name)

	}

}
