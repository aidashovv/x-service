package usecases

import (
	"context"
	"x-service/internal/users/models"

	"github.com/google/uuid"
)

type Service interface {
	Add(ctx context.Context, user *models.User) error
	Get(ctx context.Context, id uuid.UUID) (*models.User, error)
	UpdatePassword(ctx context.Context, id uuid.UUID, newPassword string) error
	Delete(ctx context.Context, id uuid.UUID) error
}
