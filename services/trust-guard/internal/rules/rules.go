package rules

import (
	"fmt"

	"github.com/gursimransw/BearBreach/services/trust-guard/internal/types"
)

func GetRuleByID(ruleID string, rules []types.DetectionRule) (types.DetectionRule, error) {

	// blankRule := types.DetectionRule{

	// 	ID:          "",
	// 	Name:        "",
	// 	Category:    "",
	// 	FindingType: "",
	// 	AppliesTo:   make([]string, 0),
	// 	Pattern:     "",
	// 	Weight:      0.00,
	// 	Severity:    "",
	// 	Reason:      "",
	// }

	for _, rule := range rules {
		if rule.ID == ruleID {
			return rule, nil
		}

	}

	return types.DetectionRule{}, fmt.Errorf("rule %s not found", ruleID)

}
