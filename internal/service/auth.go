package service

import "auth/internal/repository"

type AuthService struct {
	authRepository repository.User
}

func NewAuthService(aR repository.User) *AuthService {
	return &AuthService{aR}
}

func (aS *AuthService) CreateUser() {}
