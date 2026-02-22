package dtos

import "x-service/internal/users/models"

func ToDBUser(user *models.User) DBUser {
	return NewDBUser(
		user.ID,
		user.Username.GetName(),
		user.Password.GetContent(),
		user.Age.GetYears(),
	)
}

func ToUser(dbUser DBUser) *models.User {
	return models.NewUserFromDB(
		dbUser.ID,
		dbUser.Username,
		dbUser.Password,
		dbUser.Age,
	)
}
