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
	GenerateTokens(context.Context, AuthParams) (Tokens, error)
}

type Services struct {
	Auth
}

type ServicesDeps struct {
	Repository       *repository.Repositories
	SecretKeyAccess  []byte
	SecretKeyRefresh []byte
}

func NewServices(deps ServicesDeps) *Services {
	return &Services{Auth: NewAuthService(deps.Repository.User, deps.SecretKeyAccess, deps.SecretKeyRefresh)}

}
