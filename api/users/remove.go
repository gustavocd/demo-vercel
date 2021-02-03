package users

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gustavocd/demo-vercel/web"
)

// Remove removes an existing User from our database.
func Remove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	const q = `DELETE FROM users WHERE id = $1;`

	db, _ := web.OpenConn()
	if _, err := db.ExecContext(r.Context(), q, id); err != nil {
		web.Respond(w, nil, http.StatusInternalServerError)
		return
	}

	var response struct {
		Message string `json:"message"`
	}

	response.Message = "User deleted!"

	web.Respond(w, response, http.StatusOK)
}
