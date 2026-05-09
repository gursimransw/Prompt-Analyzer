package logic

import (
	"regexp"
	"strings"

	"github.com/gursimransw/prompt-analyzer/internal/types"
)

//This package will contain the log for our prompt analyzer so that it can match the patterns against
//Our Library of malicious prompts using simple regex

func MatchPromptPattern(config *types.PatternConfig, target_string string) (bool, []string) {
	//This function requires the config patterns for detecting against, and the prompt string itself
	//It returns whether the prompt matches ant patters, if yes, then of which category ?

	var matchedCategories []string //Empty slice for all matched categories
	patternMatched := false        //Patter match status  variable

	target_string = strings.ToLower(target_string)

	for category, patterns := range config.Categories {
		//Looping for each category in the config
		matchCount := 0
		//Inital matchCount is set to zero
		for _, pattern := range patterns {
			//Looping over all the patterns within the category
			matched, err := regexp.MatchString(pattern, target_string)

			if err != nil {
				panic(err)

			} else if matched == true {
				matchCount++ //Incrementing the matchcount counter on successful hit.
				patternMatched = true

			} else {
				continue
			}

		}

		if matchCount > 0 {
			matchedCategories = append(matchedCategories, category) //Appending results to the slice
		}
	}

	return patternMatched, matchedCategories

}
