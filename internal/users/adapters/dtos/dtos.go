package dtos

import (
	"github.com/google/uuid"
)

type DBUser struct {
	ID       uuid.UUID `db:"id"`
	Username string    `db:"username"`
	Password string    `db:"password"`
	Age      int       `db:"age"`
}

func NewDBUser(id uuid.UUID, username, password string, age int) DBUser {
	return DBUser{
		ID:       id,
		Username: username,
		Password: password,
		Age:      age,
	}

}
