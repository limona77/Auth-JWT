package repository

import (
	"auth/pkg/postgres"
)

type User interface {
	CreateUser()
}

type Repositories struct {
	User
}

func NewRepositories(db *postgres.DB) *Repositories {
	return &Repositories{&UserRepository{db}}
}
