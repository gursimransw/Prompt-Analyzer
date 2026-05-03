package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}

}

func ValidationError(errs validator.ValidationErrors) Response {
	var errMsgs []string

	for _, err := range errs {

		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("fields %s is a required field", err.Field()))

		default:
			errMsgs = append(errMsgs, fmt.Sprintf("fields %s is invalid", err.Field()))
		}

	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ","),
	}

}
