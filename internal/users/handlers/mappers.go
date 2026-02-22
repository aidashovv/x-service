package handlers

import (
	"x-service/internal/users/models"
)

func toUser(userDTO UserDTO) (*models.User, error) {
	user, err := models.NewUser(
		userDTO.Username,
		userDTO.Password,
		userDTO.Age,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func toResponse(user *models.User) UserDTO {
	return NewUserDTO(
		user.ID,
		user.Username.GetName(),
		user.Password.GetContent(),
		user.Age.GetYears(),
	)
}
