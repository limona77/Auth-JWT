package repository

import (
	"auth/internal/model"
	"auth/pkg/postgres"
	"context"
)

type IUser interface {
	CreateUser(context.Context, model.User) (model.User, error)
	GetUserByEmail(context.Context, string) (model.User, error)
}

type IToken interface {
	SaveToken(context.Context, model.Token) (model.Token, error)
	GetToken(ctx context.Context, id int) (model.Token, error)
	RemoveToken(ctx context.Context, token string) (int, error)
}
type Repositories struct {
	IUser
	IToken
}

func NewRepositories(db *postgres.DB) *Repositories {
	return &Repositories{NewUserRepository(db), NewTokenRepository(db)}
}
