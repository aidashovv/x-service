package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"x-service/internal/users/adapters/dtos"
	"x-service/internal/users/models"

	myerr "x-service/internal/core/errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) Add(ctx context.Context, user *models.User) error {
	q := `
		INSERT INTO users (id, username, password, age)
		VALUES ($1, $2, $3, $4)
	`

	dbUser := dtos.ToDBUser(user)
	if _, err := s.db.ExecContext(ctx, q,
		dbUser.ID,
		dbUser.Username,
		dbUser.Password,
		dbUser.Age,
	); err != nil {
		return fmt.Errorf("add user: %w", err)
	}

	return nil
}

func (s *Storage) Get(ctx context.Context, id uuid.UUID) (*models.User, error) {
	q := `
		SELECT id, username, password, age
		FROM users
		WHERE id = $1
	`

	var dbUser dtos.DBUser
	if err := s.db.GetContext(ctx, &dbUser, q, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, myerr.ErrUserNotFound
		}

		return nil, fmt.Errorf("get user: %w", err)
	}

	return dtos.ToUser(dbUser), nil
}

func (s *Storage) Update(ctx context.Context, user *models.User) error {
	q := `
		UPDATE users
		SET username = $1, password = $2, age = $3
		WHERE id = $4
	`

	dbUser := dtos.ToDBUser(user)
	result, err := s.db.ExecContext(ctx, q,
		dbUser.Username,
		dbUser.Password,
		dbUser.Age,
		dbUser.ID,
	)

	if err != nil {
		return fmt.Errorf("update user: %w", err)
	}

	if rowsAffected, err := result.RowsAffected(); err != nil {
		return fmt.Errorf("rows affected: %w", err)
	} else if rowsAffected == 0 {
		return myerr.ErrUserNotFound
	}

	return nil
}

func (s *Storage) Delete(ctx context.Context, id uuid.UUID) error {
	q := `
		DELETE FROM users
		WHERE id = $1
	`

	result, err := s.db.ExecContext(ctx, q, id)
	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	if rowsAffected, err := result.RowsAffected(); err != nil {
		return fmt.Errorf("rows affected: %w", err)
	} else if rowsAffected == 0 {
		return myerr.ErrUserNotFound
	}

	return nil
}
