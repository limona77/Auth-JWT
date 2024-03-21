package custom_errors

import "errors"

var (
	ErrAlreadyExists    = errors.New("такой пользователь уже существует")
	ErrUserNotFound     = errors.New("такого пользователья не существует")
	ErrWrongCredetianls = errors.New("попробуйте ещё раз")
)
