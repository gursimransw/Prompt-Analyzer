package logic

import (
	"regexp"

	"github.com/gursimransw/prompt-analyzer/internal/types"
)

// This function will take policy configuration pointer and weight (rule weight) and it will match them with the
// Policy Configuration library present in the location defined in the config.yaml file and will return
// severity and respective action for that specific calculated weight.
func getSeverityAndActionFromWeight(PolicyConfig *types.PolicyConfig, weight float64) (string, string) {

	var severity string //severity that the function will return
	var action string   //action that the function will return

	//Here we are looping over the severity thresholds array that is present inside of the policy configuration object
	for _, severityThreshold := range PolicyConfig.SeverityConfig {
		if weight >= severityThreshold.MinScore && weight <= severityThreshold.MaxScore {
			severity = severityThreshold.Severity

			//Here we are getting the action that is mapped to that severity in the DefaultActions attribute in the policy configuration
			action = PolicyConfig.DefaultActionsConfig[severity]
			return severity, action

		}
	}
	return "UNDEFINED", "LOG" //Safe return if the logic fails somewhere
}

func MaskRedactedValues(inputString string) string {
	runes := []rune(inputString)
	maskedRune := make([]rune, len(runes))

	for k, v := range runes {
		if k == 0 || k == len(runes)-1 {
			maskedRune[k] = v

		} else {

			maskedRune[k] = '*'

		}
	}
	return string(maskedRune)
}

func AnalyzeContent(Rules *[]types.DetectionRule, PolicyConfig *types.PolicyConfig, targetString string, scanContext string) (bool, []string, []string, []string, float64, string, string, []types.Finding, string) {
	//This function excpects an array pointer of the type DetectionRule and pointer of Policy Configuration
	//The function takes the input prompt as a paramter and matches it against the Detection Rule Logic using regex
	//This function will calculate the matched categroes, reason, siskscre, severity and verdict too.

	var matchedRules []string //Empty slice for all matched Rules
	var matchedRulesReasons []string
	var matchedRulesCategories []string
	matchedRulesEffectiveWeight := 0.00
	categorySeen := make(map[string]bool)
	detectionRuleMatched := false //Patter match status  variable
	var effectiveSeverity string
	var effectiveActions string
	var findingsList []types.Finding

	for _, rule := range *Rules {
		for _, v := range rule.AppliesTo {
			if scanContext == v {

				ruleRegexPattern := regexp.MustCompile(rule.Pattern)

				matchedIndexSlices := (ruleRegexPattern.FindAllStringIndex(targetString, -1))

				if len(matchedIndexSlices) > 0 {
					detectionRuleMatched = true                                             //If mathed once, we will just set it to true
					matchedRules = append(matchedRules, rule.ID)                            //Every rule matched will be added to the array of matched rules
					matchedRulesReasons = append(matchedRulesReasons, rule.Reason)          //For every rule matched, the reason will be added to te reason array
					matchedRulesEffectiveWeight = matchedRulesEffectiveWeight + rule.Weight //For every matched rule we are incrementing the weight
					if matchedRulesEffectiveWeight >= 1.00 {
						matchedRulesEffectiveWeight = 1.00
					} //Here we are capping the max value of weight to 1 only

					//Here we are getting the effective severity and action from the calculated weight
					effectiveSeverity, effectiveActions = getSeverityAndActionFromWeight(PolicyConfig, matchedRulesEffectiveWeight)

					if !categorySeen[rule.Category] {
						matchedRulesCategories = append(matchedRulesCategories, rule.Category)
						categorySeen[rule.Category] = true

					} //Here we are making sure that sometimes rules can have same category even if they are different. This logic will
					//prevent same category appearing for multiple times in our response

					if rule.Category == "sensitive_data" {

						for _, indexSlice := range matchedIndexSlices {

							startIndex := indexSlice[0]
							endIndex := indexSlice[1]
							matchedValue := targetString[startIndex:endIndex]

							finding := types.Finding{
								RuleID:        rule.ID,
								Category:      rule.Category,
								StartIndex:    int(startIndex),
								EndIndex:      int(endIndex),
								RedactedValue: MaskRedactedValues(matchedValue),
								FindingType:   rule.FindingType,
							}

							findingsList = append(findingsList, finding)

						}

					}

				}

			}

		}
	}

	//Adding applies to context to categories -

	//Returning them all
	return detectionRuleMatched, matchedRules, matchedRulesCategories, matchedRulesReasons, matchedRulesEffectiveWeight, effectiveSeverity, effectiveActions, findingsList, scanContext

}
