package internal

import (
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRouter(s Service) *mux.Router {
	router := mux.NewRouter()

	registerHandlers(router, s)

	return router
}

func registerHandlers(r *mux.Router, s Service) {
	r.StrictSlash(true)

	transactions := r.PathPrefix("/journal").Subrouter()
	transactions.HandleFunc("", getJournals(s)).Methods(http.MethodGet)
}
