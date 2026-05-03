package types

type PatternConfig struct {
	Categories map[string][]string `json:"categories"`
}

type InputPrompt struct {
	Prompt string `json:"prompt" validate:"required"`
}
