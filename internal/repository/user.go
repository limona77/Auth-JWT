package repository

import postgres "auth/pkg"

type UserRepository struct {
	*postgres.DB
}

func NewAuthRepository(db *postgres.DB) *UserRepository {
	return &UserRepository{db}
}

func (uR *UserRepository) CreateUser() {}
