package models

import (
	"time"
	"unicode/utf8"

	"github.com/google/uuid"

	myerr "x-service/internal/core/errors"
)

const (
	minPasswordLength = 8
	minAgeThreshold   = 16
	maxUsernameLength = 32
)

type User struct {
	uuid string

	Username Username `json:"username"`
	Password Password `json:"password"`
	Age      Age      `json:"age"`
}

func NewUser(username, password string, age int) (*User, error) {
	checkedUsername, err := NewUsername(username)
	if err != nil {
		return &User{}, err
	}

	checkedAge, err := NewAge(age)
	if err != nil {
		return &User{}, err
	}

	checkedPassword, err := NewPassword(password)
	if err != nil {
		return &User{}, err
	}

	return &User{
		uuid:     uuid.NewString(),
		Username: checkedUsername,
		Age:      checkedAge,
		Password: checkedPassword,
	}, nil
}

type Username struct {
	name string
}

func NewUsername(name string) (Username, error) {
	if name == "" {
		return Username{}, myerr.ErrUsernameIsEmpty
	}

	if utf8.RuneCountInString(name) > maxUsernameLength {
		return Username{}, myerr.ErrUsernameTooLong
	}

	return Username{
		name: name,
	}, nil
}

func (u Username) GetName() string {
	return u.name
}

type Age struct {
	years int
}

func NewAge(years int) (Age, error) {
	if years < 1 {
		return Age{}, myerr.ErrInvalidAge
	}

	if years < minAgeThreshold {
		return Age{}, myerr.ErrMinAgeThreshold
	}

	return Age{
		years: years,
	}, nil
}

func (a Age) GetYears() int {
	return a.years
}

type Password struct {
	content   string
	updatedAt time.Time
}

func NewPassword(content string) (Password, error) {
	if content == "" {
		return Password{}, myerr.ErrPasswordIsEmpty
	}

	if utf8.RuneCountInString(content) < minPasswordLength {
		return Password{}, myerr.ErrPasswordTooShort
	}

	return Password{
		content:   content,
		updatedAt: time.Now(),
	}, nil
}

func (u *User) SetPassword(content string) error {
	newPassword, err := NewPassword(content)
	if err != nil {
		return err
	}

	u.Password = newPassword
	return nil
}

func (p Password) GetContent() string {
	return p.content
}
