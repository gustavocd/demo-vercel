package users

import (
	"net/http"
	"time"

	"github.com/gustavocd/demo-vercel/web"
)

// Info represents an individual user.
type Info struct {
	ID        string    `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// NewUser represents the data needed to create a new User.
type NewUser struct {
	ID    string `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

// UpdateUser represents the data needed to update an existing User.
type UpdateUser struct {
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

// FetchAll returns a list of Users.
func FetchAll(w http.ResponseWriter, r *http.Request) {
	const q = `
			SELECT id, name, email, created_at, updated_at
			FROM users
			ORDER BY created_at DESC;
		`

	db, _ := web.OpenConn()
	var users []Info
	if err := db.SelectContext(r.Context(), &users, q); err != nil {
		web.Respond(w, users, http.StatusInternalServerError)
		return
	}

	web.Respond(w, users, http.StatusOK)
}
