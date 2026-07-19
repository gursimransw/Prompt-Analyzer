package policy

import (
	"github.com/gursimransw/BearBreach/services/trust-guard/internal/types"
)

// This function will take policy configuration pointer and weight (rule weight) and it will match them with the
// Policy Configuration library present in the location defined in the config.yaml file and will return
// severity and respective action for that specific calculated weight.
func GetSeverityAndActionFromWeight(PolicyConfig *types.PolicyConfig, weight float64) (string, string) {

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
