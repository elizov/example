package entity

import (
	"net/http"

	"github.com/mholt/binding"
)

type User struct {
	ID    int
	Email string
	Name  string
}

func (u *User) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&u.ID:    "id",
		&u.Email: "email",
		&u.Name:  "name",
	}
}
