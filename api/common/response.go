package common

import (
	"encoding/json"
	"net/http"
)

// JSONError writes a JSON error response with the given message and status code.
func JSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// JSONSuccess writes a JSON success response with the given status code and data.
func JSONSuccess(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// WriteError writes a JSON error response. Alias for JSONError with reordered params.
func WriteError(w http.ResponseWriter, statusCode int, message string) {
	JSONError(w, message, statusCode)
}

// WriteJSON writes a JSON success response. Alias for JSONSuccess.
func WriteJSON(w http.ResponseWriter, statusCode int, data any) {
	JSONSuccess(w, statusCode, data)
}
