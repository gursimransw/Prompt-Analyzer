package logic

import (
	"regexp"
	"strings"

	"github.com/gursimransw/prompt-analyzer/internal/types"
)

//This package will contain the log for our prompt analyzer so that it can match the patterns against
//Our Library of malicious prompts using simple regex

func MatchPromptPattern(config *types.PatternConfig, target_string string) (bool, []string) {

	var matchedCategories []string
	patternMatched := false

	target_string = strings.ToLower(target_string)

	for category, patterns := range config.Categories {
		matchCount := 0
		for _, pattern := range patterns {
			matched, err := regexp.MatchString(pattern, target_string)

			if err != nil {
				panic(err)

			} else if matched == true {
				matchCount++
				patternMatched = true

			} else {
				continue
			}

		}

		if matchCount > 0 {
			matchedCategories = append(matchedCategories, category)
		}
	}

	return patternMatched, matchedCategories

}
