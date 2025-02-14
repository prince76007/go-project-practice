package utils

import (
	"log"
	"net/http"
)

// RespondWithError sends a JSON response with an error message and status code.
func RespondWithError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"error": "` + message + `"}`))
}

// RespondWithJSON sends a JSON response with the given data and status code.
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
	}
}

// CheckError logs the error and sends a response if the error is not nil.
func CheckError(err error, w http.ResponseWriter) {
	if err != nil {
		log.Printf("Error: %v", err)
		RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
	}
}