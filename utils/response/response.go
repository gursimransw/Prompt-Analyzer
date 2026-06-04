package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status    string `json:"status"`
	Error     string `json:"error"`
	RequestId string `json:"requestId"`
} //This is  a struct that we have created for client side, the tags specify how they will be displayed to the end client

const (
	StatusOK    = "OK"
	StatusError = "Error"
) //Declared statuses as constants for re-usability

// WriteJson function is responsible writing the json response for the client,
func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	//w is instance of http.ResponseWriter struct interface
	//Status sets the status code like 200, 201, 300 ...
	//data interface{} means that it will allow all data forms to send the response data via this function and not just of a
	//specific kind, this makes this function more re-usable, it might occur to one that it should be restricted to Response struct type only
	//But then that restricts it usage in other places, if we ever wish to create more functions for more API endpoints
	w.Header().Set("Content-Type", "application/json")
	//Here we are writing reponse header as JSON , this tells that the reponse is in JSON
	w.WriteHeader(status)
	//Sets status

	return json.NewEncoder(w).Encode(data)
	//Converts Go object → JSON Then writes it into: w

}

// This is a General Error function that will be used to cover a broader class of errors during runtime.
// This will be the reponse body that will be given to client for all errors exceptn validation errors
func GeneralError(err error, requestId string) Response {
	return Response{
		RequestId: requestId,
		Status:    StatusError,
		Error:     err.Error(),
	}

}

// This function is specifically meant to for any validation errors like - Missing / Malformed fields
// If there's a validation error due to malformed / missing fields, this response body will be returned to the client
func ValidationError(errs validator.ValidationErrors, requestId string) Response {
	var errMsgs []string
	//Initialize an empty slice for all error messages in errs

	for _, err := range errs {
		//loop over errs

		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("fields %s is a required field", err.Field()))
			//For the fields that had "required" tag

		default:
			errMsgs = append(errMsgs, fmt.Sprintf("fields %s is invalid", err.Field()))
			//All other validation errors
		}

	}

	return Response{
		RequestId: requestId,
		Status:    StatusError,
		Error:     strings.Join(errMsgs, ","),
	} //Returning response with the exact error

}
