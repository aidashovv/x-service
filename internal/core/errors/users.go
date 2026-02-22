package errors

import "errors"

var (
	// for handlers
	ErrPasswordIsEmpty = errors.New("password is empty")
	ErrInvalidAge      = errors.New("age is invalid")

	// for domain
	ErrUsernameIsEmpty  = errors.New("username is empty")
	ErrUsernameTooLong  = errors.New("password too long")
	ErrPasswordTooShort = errors.New("password too short")
	ErrMinAgeThreshold  = errors.New("age under min threshold")

	// for storage & service
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")

	// for config
	ErrRequiredPassword = errors.New("db password is required")
)
