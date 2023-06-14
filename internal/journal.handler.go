package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func addJournal(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newJournal Journal
		if err := json.NewDecoder(r.Body).Decode(&newJournal); err != nil {
			generateErrorResponse(w, http.StatusBadRequest, fmt.Errorf("error parsing request"))
			return
		}

		resp, err := s.AddJournal(r.Context(), newJournal)
		if err != nil {
			generateErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("error saving journal"))
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err = json.NewEncoder(w).Encode(resp); err != nil {
			generateErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("error encoding response"))
			return
		}
	}
}

func getJournals(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list, err := s.GetAllJournals(r.Context())
		if err != nil {
			generateErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("error retrieving object"))
			return
		}

		w.WriteHeader(http.StatusOK)
		if err = json.NewEncoder(w).Encode(list); err != nil {
			generateErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("error encoding response"))
			return
		}
	}
}

func generateErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)

	errResp := struct {
		Message string
	}{
		Message: err.Error(),
	}

	if err = json.NewEncoder(w).Encode(errResp); err != nil {
		fmt.Printf("error generating response: %v", err)
	}
}
