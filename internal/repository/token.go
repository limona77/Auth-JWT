package repository

import (
	custom_errros "auth/internal/custom-errors"
	"auth/internal/model"
	"auth/pkg/postgres"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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

func (tR *TokenRepository) GetToken(ctx context.Context, userId int) (model.Token, error) {
	path := "internal.repository.token.RefreshToken"
	sql, args, err := tR.Builder.Select("id", "refresh_token", "user_id").
		From("public.tokens").
		Where("user_id = ?", userId).
		ToSql()
	if err != nil {
		return model.Token{}, fmt.Errorf(path+".ToSql, error: {%w}", err)
	}

	var modelToken model.Token

	err = tR.Pool.QueryRow(ctx, sql, args...).Scan(&modelToken.ID, &modelToken.RefreshToken, &modelToken.UserID)
	if err != nil {
		if err != nil {
			var pgErr *pgconn.PgError
			if ok := errors.As(err, &pgErr); ok {
				return model.Token{}, err
			}
			if errors.Is(err, pgx.ErrNoRows) {
				return model.Token{}, custom_errros.ErrUserNotFound
			}
			return model.Token{}, fmt.Errorf(path+".QueryRow, error: {%w}", err)
		}
	}
	return modelToken, nil
}
