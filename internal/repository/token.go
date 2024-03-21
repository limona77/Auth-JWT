package repository

import (
	"auth/internal/model"
	"auth/pkg/postgres"
	"context"
	"fmt"
)

type TokenRepository struct {
	*postgres.DB
}

func NewTokenRepository(db *postgres.DB) *TokenRepository {
	return &TokenRepository{db}
}

func (tR *TokenRepository) SaveToken(ctx context.Context, token model.Token) (model.Token, error) {
	path := "internal.repository.token.SaveToken"

	sql, args, err := tR.Builder.
		Insert("public.tokens").
		Columns("user_id", "refresh_token").
		Values(token.UserID, token.RefreshToken).
		Suffix(`
		ON CONFLICT (user_id) DO UPDATE
    SET refresh_token = excluded.refresh_token
		RETURNING id,refresh_token,user_id`).
		ToSql()
	if err != nil {
		return model.Token{}, fmt.Errorf(path+".ToSql, error: {%w}", err)
	}
	var t model.Token
	err = tR.Pool.QueryRow(ctx, sql, args...).
		Scan(&t.ID, &t.RefreshToken, &t.UserID)
	if err != nil {
		return model.Token{}, fmt.Errorf(path+".Scan, error: {%w}", err)
	}

	return t, nil
}
