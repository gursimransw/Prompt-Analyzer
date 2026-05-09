package types

type PatternConfig struct {
	Categories map[string][]string `json:"categories"`
}

//This is the struct of Prompt Pattern Configuration that is present in the patterns library.
//The json from patterns.json gets loaded into this struct using the LoadPatterns function in internal/utils/loader.go

type InputPrompt struct {
	Prompt string `json:"prompt" validate:"required"`
}

//This is a struct for the input prompt that we recieve from our API endpoint
