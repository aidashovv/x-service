package adapters

import (
	"context"
	"x-service/internal/users/models"

	"github.com/google/uuid"
)

type Repository interface {
	Add(ctx context.Context, user *models.User) error
	Get(ctx context.Context, id uuid.UUID) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}
