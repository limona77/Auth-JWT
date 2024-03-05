package service

import (
	"auth/internal/repository"
	"context"
)

type AuthParams struct {
	Email    string
	Password string
}

type Auth interface {
	CreateUser(context.Context, AuthParams) error
}

type Services struct {
	Auth
}

type ServicesDeps struct {
	Repository *repository.Repositories
}

func NewServices(deps ServicesDeps) *Services {
	return &Services{Auth: NewAuthService(deps.Repository.User)}

}
