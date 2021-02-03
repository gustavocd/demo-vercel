// Package web knows how to manage different responses for the clients.
package web

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // The database driver in use.
)

// Respond converts a Go value to JSON and sends it to the client.
func Respond(w http.ResponseWriter, data interface{}, statusCode int) error {
	// Convert the response value to JSON.
	res, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Respond with the provided JSON.
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if _, err := w.Write(res); err != nil {
		return err
	}

	return nil
}

// OpenConn ...
func OpenConn() (*sqlx.DB, error) {
	// Query parameters.
	q := make(url.Values)
	q.Set("sslmode", "require")
	q.Set("timezone", "utc")

	// Construct url.
	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD")),
		Host:     os.Getenv("DB_HOST"),
		Path:     os.Getenv("DB_PATH"),
		RawQuery: q.Encode(),
	}

	db, err := sqlx.Connect("postgres", u.String())
	if err != nil {
		return nil, err
	}

	return db, nil
}
