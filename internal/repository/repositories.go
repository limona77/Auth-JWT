package repository

import (
	"auth/internal/model"
	"auth/pkg/postgres"
	"context"
)

type User interface {
	CreateUser(context.Context, model.User) (int, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
}

type Repositories struct {
	User
}

func NewRepositories(db *postgres.DB) *Repositories {
	return &Repositories{NewUserRepository(db)}
}
