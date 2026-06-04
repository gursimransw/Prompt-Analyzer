package types

type InputPrompt struct {
	Prompt string `json:"prompt" validate:"required"`
}

//This is a struct for the input prompt that we recieve from our API endpoint

// This is the struct for the detection rules for analyzing prompts, here we are not creating it like a config
// Rather we are creating a struct for the rule itself as now we are gonna use the list of rules directly into our code.
type DetectionRule struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Pattern  string  `json:"pattern"`
	Weight   float64 `json:"weight"`
	Severity string  `json:"severity"`
	Reason   string  `json:"reason"`
}

// This is the struct for severity configurations mappings as described inside /config/policy/PolicyConfig.json file.
type SeverityConfig struct {
	Severity string  `json:"severity"`
	MinScore float64 `json:"min_score"`
	MaxScore float64 `json:"max_score"`
}

// PolicyConfig struct will allow us to reference the policy itself anywhere in the code and we can use both severity config and
// Default actions config wherever we need.
type PolicyConfig struct {
	SeverityConfig       []SeverityConfig  `json:"severity_thresholds"`
	DefaultActionsConfig map[string]string `json:"default_actions"`
}
