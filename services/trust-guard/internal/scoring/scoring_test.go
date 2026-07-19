package scoring

import (
	"testing"

	"github.com/gursimransw/BearBreach/services/trust-guard/internal/types"
)

func TestGetEffectiveRiskScoreByMatchedRules(t *testing.T) {

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

		ID:          "SD-003",
		Name:        "Test Rule",
		Category:    "sensitive_data",
		FindingType: "api_key",
		AppliesTo:   []string{"input", "output"},
		Pattern:     `Testing Strin, no logic needed`,
		Weight:      0.86,
		Severity:    "critical",
		Reason:      "Detected generic API key or secret key assignment",
	}

	matchedRules := types.DetectionRuleList{
		RuleList: []types.DetectionRule{rule1, rule2},
	}

	effectiveWeight := GetEffectiveRiskScoreByMatchedRules(matchedRules)

	if effectiveWeight != 1 {
		t.Errorf("Expected Effective Risk Score as 1, got %f", effectiveWeight)
	}

}
