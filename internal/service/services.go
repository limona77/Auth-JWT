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
type IAuth interface {
	Register(ctx context.Context, params AuthParams) (Tokens, model.User, error)
	GenerateTokens(ctx context.Context, params AuthParams) (Tokens, model.User, error)
	Refresh(ctx context.Context, token string) (Tokens, model.User, error)
	Login(ctx context.Context, params AuthParams) (Tokens, model.User, error)
	Logout(ctx context.Context, token string) (int, error)
}
type IClient interface {
	VerifyToken(token string) (TokenClaims, error)
	GetUserByEmail(ctx context.Context, params AuthParams) (model.User, error)
}
type Services struct {
	IAuth
	IClient
}

type ServicesDeps struct {
	Repository       *repository.Repositories
	SecretKeyAccess  []byte
	SecretKeyRefresh []byte
}

func NewServices(deps ServicesDeps) *Services {
	return &Services{
		IAuth: NewAuthService(
			deps.Repository.IUser,
			deps.Repository.IToken,
			deps.SecretKeyAccess,
			deps.SecretKeyRefresh),
		IClient: NewClientService(deps.Repository.IUser),
	}
}
