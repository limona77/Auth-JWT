package service

import (
	"auth/internal/hashPassword"
	"auth/internal/model"
	"auth/internal/repository"
	"context"
	"fmt"
)

type AuthService struct {
	authRepository repository.User
}

func NewAuthService(aR repository.User) *AuthService {
	return &AuthService{aR}
}

func (aS *AuthService) CreateUser(ctx context.Context, params AuthParams) error {
	path := "internal.service.auth.CreateUser"
	password, err := hashPassword.HashPassword(params.Password)
	if err != nil {
		return fmt.Errorf(path+".HashPassword, error: {%w}", err)
	}

	_, err = aS.authRepository.CreateUser(ctx, model.User{params.Email, password})
	if err != nil {
		return fmt.Errorf(path+".CreateUser, error: {%w}", err)
	}
	return nil
}
