package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const jsonContentType = "application/json"
const contentType = "Content-Type"

// Writes json formatted response.
func JsonResponse(w http.ResponseWriter, response interface{}) {
	result, err := json.Marshal(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while parsing json response: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set(contentType, jsonContentType)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

// Writes json formatted error.
func JsonErrorResponse(w http.ResponseWriter, err error, code int) {
	result, err := json.Marshal(err.Error())
	if err != nil {
		http.Error(w, fmt.Sprintf("error while parsing error response: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set(contentType, jsonContentType)
	w.WriteHeader(code)
	w.Write(result)
}

// Sets content type json and status OK
func JsonNullHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set(contentType, jsonContentType)
	w.WriteHeader(http.StatusOK)
}
