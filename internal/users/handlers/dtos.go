package handlers

import (
	"encoding/json"
)

type UserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

func NewUserDTO(username, password string, age int) UserDTO {
	return UserDTO{
		Username: username,
		Password: password,
		Age:      age,
	}
}

func (u UserDTO) ToBytes() []byte {
	b, err := json.MarshalIndent(u, "", "	")
	if err != nil {
		panic(err)
	}

	return b
}

type PasswordUserDTO struct {
	Content string `json:"password"`
}

func NewPasswordUserDTO(content string) PasswordUserDTO {
	return PasswordUserDTO{
		Content: content,
	}
}

func (p PasswordUserDTO) GetContent() string {
	return p.Content
}

type ErrorDTO struct {
	Message string
}

func NewErrorDTO(message string) ErrorDTO {
	return ErrorDTO{
		Message: message,
	}
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "	")
	if err != nil {
		panic(err) // if basic thing, like "2 + 2 = 4", wrong, app should panic
	}

	return string(b)
}
