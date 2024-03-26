package custom_errors

import "errors"

var (
	ErrAlreadyExists    = errors.New("такой пользователь уже существует")
	ErrUserNotFound     = errors.New("такого пользователья не существует")
	ErrWrongCredetianls = errors.New("неверная почта или пароль,попробуйте ещё раз")
	ErrUserUnauthorized = errors.New("пользователь не авторизован")
)
