// Package api ...
package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gustavocd/demo-vercel/api/users"
)

// API represents the entry point for all our serverless functions.
func API(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()

	router.HandleFunc("/api/users", users.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/users", users.FetchAll).Methods(http.MethodGet)
	router.HandleFunc("/api/users/{id}", users.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/users/{id}", users.Remove).Methods(http.MethodDelete)
	router.HandleFunc("/api/users/{id}", users.FetchByID).Methods(http.MethodGet)

	router.ServeHTTP(w, r)
}
