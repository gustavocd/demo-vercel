package users

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gustavocd/demo-vercel/web"
)

// Update updates an existing User in our database.
func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user UpdateUser

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		web.Respond(w, nil, http.StatusBadRequest)
		return
	}

	const q = `
		UPDATE users
		SET name = $1, email = $2
		WHERE id = $3;
	`

	db, _ := web.OpenConn()
	if _, err := db.ExecContext(r.Context(), q, user.Name, user.Email, id); err != nil {
		web.Respond(w, nil, http.StatusInternalServerError)
		return
	}

	web.Respond(w, user, http.StatusCreated)
}
