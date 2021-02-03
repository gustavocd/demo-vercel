package users

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gustavocd/demo-vercel/web"
)

// Create adds a new User to our database.
func Create(w http.ResponseWriter, r *http.Request) {
	var user NewUser

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		web.Respond(w, nil, http.StatusBadRequest)
		return
	}

	user.ID = uuid.New().String()

	const q = `
			INSERT INTO users (id, name, email)
			VALUES ($1, $2, $3);
		`

	db, _ := web.OpenConn()
	if _, err := db.ExecContext(r.Context(), q, user.ID, user.Name, user.Email); err != nil {
		web.Respond(w, nil, http.StatusInternalServerError)
		return
	}

	web.Respond(w, user, http.StatusCreated)
}
