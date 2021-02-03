// Package users manages all CRUD related operations.
package users

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gustavocd/demo-vercel/web"
)

// FetchByID returns a single User.
func FetchByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user Info
	const q = `
		SELECT id, name, email, created_at, updated_at
		FROM users
		WHERE id = $1;
	`

	db, _ := web.OpenConn()
	if err := db.GetContext(r.Context(), &user, q, id); err != nil {
		web.Respond(w, nil, http.StatusInternalServerError)
		return
	}

	web.Respond(w, user, http.StatusOK)
}
