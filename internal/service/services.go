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
	Register(ctx context.Context, params AuthParams) (Tokens, model.User, error)
	GenerateTokens(ctx context.Context, params AuthParams) (Tokens, model.User, error)
	SaveToken(ctx context.Context, token model.Token) (model.Token, error)
	Refresh(ctx context.Context, token string) (Tokens, model.User, error)
	Login(ctx context.Context, params AuthParams) (Tokens, model.User, error)
}
type Client interface {
	VerifyToken(token string) (TokenClaims, error)
	GetUserByEmail(ctx context.Context, params AuthParams) (model.User, error)
}
type Services struct {
	Auth
	Client
}

type ServicesDeps struct {
	Repository       *repository.Repositories
	SecretKeyAccess  []byte
	SecretKeyRefresh []byte
}

func NewServices(deps ServicesDeps) *Services {
	return &Services{
		Auth: NewAuthService(
			deps.Repository.User,
			deps.Repository.Token,
			deps.SecretKeyAccess,
			deps.SecretKeyRefresh),
		Client: NewClientService(deps.Repository.User),
	}
}
