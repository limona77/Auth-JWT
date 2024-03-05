package custom_errros

import "errors"

var (
	ErrAlreadyExists = errors.New("user already exists")
)
