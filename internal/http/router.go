package http

import (
	"github.com/gorilla/mux"
	"mission-data-challenge/internal/service"
	"net/http"
)

func SetupRouter(s service.Service) *mux.Router {
	router := mux.NewRouter()

	registerHandlers(router, s)

	return router
}

func registerHandlers(r *mux.Router, s service.Service) {
	r.StrictSlash(true)
	r.Use(commonMiddleware)

	journals := r.PathPrefix("/journals").Subrouter()
	journals.HandleFunc("", getJournals(s)).Methods(http.MethodGet)
	journals.HandleFunc("", addJournal(s)).Methods(http.MethodPost)

	entries := r.PathPrefix("/journals/{journalID}/entries").Subrouter()
	entries.HandleFunc("", getEntries(s)).Methods(http.MethodGet)
	entries.HandleFunc("", addEntry(s)).Methods(http.MethodPost)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
