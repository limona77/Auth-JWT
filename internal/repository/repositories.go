package repository

import (
	"auth/internal/model"
	"auth/pkg/postgres"
	"context"
)

type User interface {
	CreateUser(context.Context, model.User) (model.User, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
}

//	type Token interface {
//		SaveToken(ctx context.Context, token string) error
//	}
type Repositories struct {
	User
	//Token
}

func NewRepositories(db *postgres.DB) *Repositories {
	return &Repositories{NewUserRepository(db)}
}
