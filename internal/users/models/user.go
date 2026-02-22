package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `json:"id" db:"id"`

	Username Username `json:"username" db:"username"`
	Password Password `json:"password" db:"password"`
	Age      Age      `json:"age" db:"age"`
}

func NewUser(username, password string, age int) (*User, error) {
	checkedUsername, err := NewUsername(username)
	if err != nil {
		return nil, err
	}

	checkedAge, err := NewAge(age)
	if err != nil {
		return nil, err
	}

	checkedPassword, err := NewPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       uuid.New(),
		Username: checkedUsername,
		Age:      checkedAge,
		Password: checkedPassword,
	}, nil
}

func NewUserFromDB(id uuid.UUID, username, password string, age int) *User {
	dbUser, _ := NewUser(username, password, age) // невозможно получить инвалидную модель из бд
	dbUser.ID = id

	return dbUser
}
