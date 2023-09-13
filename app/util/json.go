package util

import (
	"encoding/json"
	"net/http"
	"strings"
)

type JsonError struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

type JsonResponseError struct {
	Errors []JsonError `json:"errors"`
}

func WriteJsonError(w http.ResponseWriter, statusCode int, message ...string) {
	response := JsonResponseError{
		Errors: []JsonError{
			{
				Reason:  http.StatusText(statusCode),
				Message: strings.Join(message, "\n"),
			},
		},
	}

	WriteJson(w, response, statusCode)
}

func WriteJson(w http.ResponseWriter, v interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(v)
	w.Header().Set("Content-Type", "application/json")
}
