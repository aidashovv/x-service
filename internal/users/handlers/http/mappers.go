package http

import (
	"pdd/internal/users/models"
)

func ToUser(userDTO UserDTO) (*models.User, error) {
	user, err := models.NewUser(
		userDTO.Username,
		userDTO.Password,
		userDTO.Age,
	)
	if err != nil {
		return &models.User{}, err
	}

	return user, nil
}

func toResponse(user *models.User) UserDTO {
	return NewUserDTO(
		user.Username.GetName(),
		user.Password.GetContent(),
		user.Age.GetYears(),
	)
}
