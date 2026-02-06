package services

import (
	"errors"
	"pdd/internal/users/adapters"
	"pdd/internal/users/models"

	myerr "pdd/internal/users/errors"
)

type Service interface {
	Add(user *models.User) error
	Get(username string) (*models.User, error)
	UpdatePassword(username, newPassword string) error
	Delete(username string) error
}

type UserService struct {
	repository adapters.Repository
}

func NewUserService(storage *adapters.Storage) *UserService {
	return &UserService{
		repository: storage,
	}
}

func (us *UserService) Add(user *models.User) error {
	return us.repository.Add(user)
}

func (us UserService) Get(username string) (*models.User, error) {
	return us.repository.Get(username)
}

func (us *UserService) UpdatePassword(username, newPassword string) error {
	user, err := us.Get(username)
	if err != nil {
		if errors.Is(err, myerr.ErrUserNotFound) {
			return myerr.ErrUserNotFound
		}

		return myerr.ErrSmthIsWrong
	}

	if err := user.SetPassword(newPassword); err != nil {
		switch {
		case errors.Is(err, myerr.ErrPasswordIsEmpty):
			return myerr.ErrPasswordIsEmpty
		case errors.Is(err, myerr.ErrPasswordTooShort):
			return myerr.ErrPasswordTooShort
		default:
			return myerr.ErrSmthIsWrong
		}
	}

	return us.repository.Update(username)
}

func (us *UserService) Delete(username string) error {
	return us.repository.Delete(username)
}
