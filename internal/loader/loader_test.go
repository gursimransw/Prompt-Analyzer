package loader

import (
	"testing"
)

func TestLoaderFunctions(t *testing.T) {

	policyConfigFilePath := "../../configs/policy/policy.json"
	detectionRuleLibraryPath := "../../configs/rules/rules.json"

	policyConfig, err := LoadPolicyConfig(policyConfigFilePath)

	if err != nil {
		t.Fatalf("Failed to load policy confgurations from path %s , ERROR - %v. **TEST FAILED**", policyConfigFilePath, err)
	} else {
		t.Logf("Policy Loaded from path %s successfully. **TEST PASSED**", policyConfigFilePath)
	}

	policyConfigActions := policyConfig.DefaultActionsConfig
	policyConfigSeverityMapping := policyConfig.SeverityConfig

	if len(policyConfigActions) > 0 {
		t.Logf("Successfully loaded policy action contents - %+v **TEST PASSED** ", policyConfigActions)
	} else {
		t.Fatalf("Failed to load policy action contents. **TEST FAILED**")

	}

	if len(policyConfigSeverityMapping) > 0 {
		t.Logf("Successfully loaded policy severity mapping contents - %+v **TEST PASSED** ", policyConfigSeverityMapping)
	} else {
		t.Fatalf("Failed to load policy severity mapping contents. **TEST FAILED**")

	}

	detectionRulesLibrary, err := LoadDetectionRules(detectionRuleLibraryPath)
	if err != nil {
		t.Fatalf("Failed to load detection rules from path %s , ERROR - %v. **TEST FAILED**", detectionRuleLibraryPath, err)
	} else {
		t.Logf("Detection Rules loaded from path %s successfully. **TEST PASSED**", detectionRuleLibraryPath)
	}

	if len(*detectionRulesLibrary) > 0 {

		rules := *detectionRulesLibrary
		rule := rules[0]
		t.Logf("Successfully loaded rules, EXAMPLE RULE => %+v **TEST PASSED**", rule)

	} else {
		t.Fatalf("Failed to load detection rules. **TEST FAILED**")

	}

}
