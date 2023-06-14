package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getJournals(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode("hello world"); err != nil {
			generateErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("error encoding response"))
			return
		}
	}
}

func generateErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)

	errResp := struct {
		message string
	}{
		message: err.Error(),
	}

	if err = json.NewEncoder(w).Encode(errResp); err != nil {
		fmt.Printf("error generating response: %v", err)
	}
}
