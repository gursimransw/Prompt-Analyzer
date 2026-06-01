package prompts

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gursimransw/prompt-analyzer/internal/logic"
	"github.com/gursimransw/prompt-analyzer/internal/types"
	"github.com/gursimransw/prompt-analyzer/utils/response"
)

//This is our API endpoint function, the http.HandlerFunc basically allow us to expose normal functions
//As an API endpoint. Here the PromptAnalyzer function will take detectionRules  library and policy configuration as an input and analyze the prompt given by the user.
//And give a verdict

func PromptAnalyzer(detectionRules *[]types.DetectionRule, policyConfig *types.PolicyConfig) http.HandlerFunc { //This is the API layer for the function

	//This is where we are handling all the logic, the PromptAnalyzer returns a function , a function that is of type http.HandlerFunc
	//This is what that function looks like
	return func(w http.ResponseWriter, r *http.Request) {
		//The return function takes 2 inputs, the reponse writer and the request body.

		//Initialize a unique requestId as soon as an API request is received.
		requestId := uuid.NewString()

		slog.Info("API Request received", slog.String("method", r.Method), slog.String("endpoint", r.RequestURI), slog.String("remoteAddress", r.RemoteAddr), slog.String("userAgent", r.UserAgent()), slog.String("requestId", requestId))

		var prompt types.InputPrompt
		//Initializing a prompt struct of type InputPrompt, this will be used to access the response body of the
		//Request submitted by the user , i.e a prompt. This will allow us to access variables inside the json request

		err := json.NewDecoder(r.Body).Decode(&prompt)
		//Here we are decoding the response r and writing it into the prompt variable pointer location that we declared in the previous line

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body"), requestId))
			slog.Warn("Received an empty request body.", slog.String("requestId", requestId))
			return

		} //Check for empty body

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err, requestId))
			slog.Warn("Received a malformed request.", slog.String("Error", err.Error()), slog.String("requestId", requestId))
			return
		} //Check for any other error

		//Request Validation & returning validation errors accordingly

		if err := validator.New().Struct(prompt); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs, requestId))
			slog.Warn("Failed to validate the request, please follow proper schema for the request", slog.String("Error", validateErrs.Error()), slog.String("requestId", requestId))
			return
		}

		slog.Info("Analyzing the Prompt....", slog.String("prompt", prompt.Prompt), slog.String("requestId", requestId))

		detectionRuleMatched, matchedRules, matchedRulescategories, matchedRulesReasons, matchedRulesEffectiveWeight, effectiveSeverity, effectiveActions := logic.MatchPromptPattern(detectionRules, policyConfig, prompt.Prompt)

		slog.Info("Prompt Analysis Completed",
			slog.Bool("matched", detectionRuleMatched),
			slog.Any("rules", matchedRules),
			slog.String("input", prompt.Prompt),
			slog.Float64("riskScore", matchedRulesEffectiveWeight),
			slog.String("severity", effectiveSeverity),
			slog.String("verdict", effectiveActions),
			slog.Any("categories", matchedRulescategories),
			slog.Any("reasons", matchedRulesReasons),
			slog.String("requestId", requestId))
		//Analyzing the prompt

		//Using WriteJson function to return the following status on success
		response.WriteJson(w, http.StatusOK, map[string]interface{}{
			"requestId":  requestId,
			"matched":    detectionRuleMatched,
			"rules":      matchedRules,
			"input":      prompt.Prompt,
			"riskScore":  matchedRulesEffectiveWeight,
			"severity":   effectiveSeverity,
			"verdict":    effectiveActions,
			"categories": matchedRulescategories,
			"reasons":    matchedRulesReasons,
		})

	}
}
