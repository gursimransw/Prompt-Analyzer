package outputs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gursimransw/BearBreach/services/trust-guard/internal/detector"
	"github.com/gursimransw/BearBreach/services/trust-guard/internal/models"
	"github.com/gursimransw/BearBreach/services/trust-guard/internal/types"
)

//This is our API endpoint function, the http.HandlerFunc basically allow us to expose normal functions
//As an API endpoint. Here the PromptAnalyzer function will take detectionRules  library and policy configuration as an input and analyze the prompt given by the user.
//And give a verdict

func OutputAnalyzer(detectionRules *[]types.DetectionRule, policyConfig *types.PolicyConfig) http.HandlerFunc { //This is the API layer for the function

	//This is where we are handling all the logic, the PromptAnalyzer returns a function , a function that is of type http.HandlerFunc
	//This is what that function looks like
	return func(w http.ResponseWriter, r *http.Request) {
		//The return function takes 2 inputs, the reponse writer and the request body.

		//Initialize a unique requestId as soon as an API request is received.
		requestId := uuid.NewString()

		slog.Info("API Request received", slog.String("method", r.Method), slog.String("endpoint", r.RequestURI), slog.String("remoteAddress", r.RemoteAddr), slog.String("userAgent", r.UserAgent()), slog.String("requestId", requestId))

		var output types.Content
		//Initializing a prompt struct of type InputPrompt, this will be used to access the response body of the
		//Request submitted by the user , i.e a prompt. This will allow us to access variables inside the json request

		err := json.NewDecoder(r.Body).Decode(&output)
		//Here we are decoding the response r and writing it into the prompt variable pointer location that we declared in the previous line

		if errors.Is(err, io.EOF) {
			models.WriteJson(w, http.StatusBadRequest, models.GeneralError(fmt.Errorf("empty body"), requestId))
			slog.Warn("Received an empty request body.", slog.String("requestId", requestId))
			return

		} //Check for empty body

		if err != nil {
			models.WriteJson(w, http.StatusBadRequest, models.GeneralError(err, requestId))
			slog.Warn("Received a malformed request.", slog.String("Error", err.Error()), slog.String("requestId", requestId))
			return
		} //Check for any other error

		//Request Validation & returning validation errors accordingly

		if err := validator.New().Struct(output); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			models.WriteJson(w, http.StatusBadRequest, models.ValidationError(validateErrs, requestId))
			slog.Warn("Failed to validate the request, please follow proper schema for the request", slog.String("Error", validateErrs.Error()), slog.String("requestId", requestId))
			return
		}

		slog.Info("Analyzing the Prompt....", slog.String("prompt", detector.MaskRedactedValues(output.Content)), slog.String("requestId", requestId))

		detectionRuleMatched, matchedRules, matchedRulesCategories, matchedRulesReasons, matchedRulesEffectiveWeight, effectiveSeverity, effectiveActions, findings, scanContext := detector.AnalyzeContent(detectionRules, policyConfig, output.Content, "output")

		slog.Info("Prompt Analysis Completed",
			slog.Bool("matched", detectionRuleMatched),
			slog.Any("rules", matchedRules),
			//slog.String("input", prompt.Prompt),
			slog.Float64("riskScore", matchedRulesEffectiveWeight),
			slog.String("severity", effectiveSeverity),
			slog.String("verdict", effectiveActions),
			slog.Any("categories", matchedRulesCategories),
			slog.Any("reasons", matchedRulesReasons),
			slog.Any("findings", findings),
			slog.String("requestId", requestId))
		//Analyzing the output

		//Using WriteJson function to return the following status on success
		models.WriteJson(w, http.StatusOK, map[string]interface{}{
			"requestId": requestId,
			"matched":   detectionRuleMatched,
			"rules":     matchedRules,
			"findings":  findings,
			//"input":      logic.MaskRedactedValues(prompt.Prompt),
			"riskScore":   matchedRulesEffectiveWeight,
			"severity":    effectiveSeverity,
			"verdict":     effectiveActions,
			"categories":  matchedRulesCategories,
			"reasons":     matchedRulesReasons,
			"scanContext": scanContext,
		})

	}
}
