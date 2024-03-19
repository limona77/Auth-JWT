package service

import (
	"auth/internal/model"
	"auth/internal/repository"
	"context"
)

type AuthParams struct {
	Email    string
	Password string
}

//go:generate mockgen -source=services.go -destination=mocks/mock.go
type Auth interface {
	CreateUser(context.Context, AuthParams) (model.User, error)
	GenerateTokens(context.Context, AuthParams) (Tokens, error)
	GetUserByEmail(context.Context, AuthParams) (model.User, error)
	ParseToken(string) error
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
