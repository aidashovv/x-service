package usecases

import (
	"context"
	"errors"
	"fmt"
	"x-service/internal/users/adapters"
	"x-service/internal/users/models"

	myerr "x-service/internal/core/errors"

	"github.com/google/uuid"
)

type UserService struct {
	repository adapters.Repository
}

func NewUserService(storage adapters.Repository) *UserService {
	return &UserService{
		repository: storage,
	}
}

func (us *UserService) Add(ctx context.Context, user *models.User) error {
	return us.repository.Add(ctx, user)
}

func (us *UserService) Get(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return us.repository.Get(ctx, id)
}

func (us *UserService) UpdatePassword(ctx context.Context, id uuid.UUID, newPassword string) error {
	user, err := us.Get(ctx, id)
	if err != nil {
		if errors.Is(err, myerr.ErrUserNotFound) {
			return myerr.ErrUserNotFound
		}

		return fmt.Errorf("get user: %w", err)
	}

	if err := user.SetPassword(newPassword); err != nil {
		switch {
		case errors.Is(err, myerr.ErrPasswordIsEmpty):
			return myerr.ErrPasswordIsEmpty
		case errors.Is(err, myerr.ErrPasswordTooShort):
			return myerr.ErrPasswordTooShort
		default:
			return fmt.Errorf("set password: %w", err)
		}
	}

	return us.repository.Update(ctx, user)
}

func (us *UserService) Delete(ctx context.Context, id uuid.UUID) error {
	return us.repository.Delete(ctx, id)
}
