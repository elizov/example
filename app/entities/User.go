package entities

import (
	"net/http"

	"github.com/mholt/binding"
)

// User - user entity
type User struct {
	ID    int
	Email string
	Name  string
}

// FieldMap - binding
func (u *User) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&u.ID:    "id",
		&u.Email: "email",
		&u.Name:  "name",
	}
}
