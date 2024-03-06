package custom_errors

import "errors"

var (
	ErrAlreadyExists = errors.New("user already exists")
	ErrUserNotFound  = errors.New("user not found")
)
