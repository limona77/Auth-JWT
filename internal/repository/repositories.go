package repository

import (
	"auth/internal/model"
	"auth/pkg/postgres"
	"context"
)

type User interface {
	CreateUser(context.Context, model.User) (model.User, error)
	GetUserByEmail(context.Context, string) (model.User, error)
}

type Token interface {
	SaveToken(context.Context, model.Token) (model.Token, error)
}
type Repositories struct {
	User
	Token
}

func NewRepositories(db *postgres.DB) *Repositories {
	return &Repositories{NewUserRepository(db), NewTokenRepository(db)}
}
