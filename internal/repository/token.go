package repository

import (
	"auth/pkg/postgres"
)

type TokenRepository struct {
	*postgres.DB
}

func NewTokenRepository(db *postgres.DB) *UserRepository {
	return &UserRepository{db}
}

//func (tR *TokenRepository) SaveToken() error {
//	return nil
//}
