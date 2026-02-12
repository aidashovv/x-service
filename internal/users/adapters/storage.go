package adapters

import (
	"sync"
	"x-service/internal/users/models"

	myerr "x-service/internal/core/errors"
)

type Repository interface {
	Add(user *models.User) error
	Get(username string) (*models.User, error)
	Update(username string) error
	Delete(username string) error
}

type Storage struct {
	users map[string]*models.User
	mx    sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		users: make(map[string]*models.User),
	}
}

func (s *Storage) Add(user *models.User) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	if _, exists := s.users[user.Username.GetName()]; exists {
		return myerr.ErrUserAlreadyExists
	}

	s.users[user.Username.GetName()] = user
	return nil
}

func (s *Storage) Get(username string) (*models.User, error) {
	s.mx.RLock()
	defer s.mx.RUnlock()

	user, exists := s.users[username]
	if !exists {
		return nil, myerr.ErrUserNotFound
	}

	return user, nil
}

func (s *Storage) Update(username string) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	user, exists := s.users[username]
	if !exists {
		return myerr.ErrUserNotFound
	}

	s.users[username] = user
	return nil
}

func (s *Storage) Delete(username string) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	if _, exists := s.users[username]; !exists {
		return myerr.ErrUserNotFound
	}

	delete(s.users, username)
	return nil
}
