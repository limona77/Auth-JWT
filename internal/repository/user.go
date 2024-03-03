package repository

import (
	"auth/pkg/postgres"
)

type UserRepository struct {
	*postgres.DB
}

//func NewAuthRepository(db *postgres.DB) *UserRepository {
//	return &UserRepository{db}
//}

func (uR *UserRepository) CreateUser() {}
