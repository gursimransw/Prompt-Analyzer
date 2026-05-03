package prompts

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gursimransw/prompt-analyzer/internal/logic"
	"github.com/gursimransw/prompt-analyzer/internal/types"
	"github.com/gursimransw/prompt-analyzer/internal/utils/response"
)

func PromptAnalyzer(config *types.PatternConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var prompt types.InputPrompt

		err := json.NewDecoder(r.Body).Decode(&prompt)

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return

		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		//Request Validation

		if err := validator.New().Struct(prompt); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		slog.Info("Analyzing the Prompt", slog.String("prompt", prompt.Prompt))

		matched, category := logic.MatchPromptPattern(config, prompt.Prompt)

		// slog.Info("User created successfully", slog.String("userId", fmt.Sprint(lastId)))

		// if err != nil {
		// 	response.WriteJson(w, http.StatusInternalServerError, err)
		// }

		response.WriteJson(w, http.StatusOK, map[string]interface{}{
			"matched":  matched,
			"category": category,
			"input":    prompt.Prompt,
		})
	}
}
