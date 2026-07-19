package policy

import (
	"testing"

	"github.com/gursimransw/BearBreach/services/trust-guard/internal/types"
)

func TestGetSeverityAndActionFromWeight(t *testing.T) {

	weight := 0.92
	testSeverityConfig := types.SeverityConfig{
		Severity: "HIGH",
		MinScore: 0.80,
		MaxScore: 0.95,
	}

	policyConfig := types.PolicyConfig{
		SeverityConfig: []types.SeverityConfig{testSeverityConfig},
		DefaultActionsConfig: map[string]string{
			"LOW":      "ALLOW",
			"MEDIUM":   "LOG",
			"HIGH":     "ALERT",
			"CRITICAL": "BLOCK",
		},
	}

	severity, action := GetSeverityAndActionFromWeight(&policyConfig, weight)

	if action != "ALERT" {
		t.Errorf("Expected ALERT, got  %s", action)
	}

	if severity != "HIGH" {
		t.Errorf("Expected HIGH, got  %s", severity)
	}

}
