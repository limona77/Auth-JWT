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
	//token, err := tR.getTokenByUserID(ctx, token.UserID)
	//if token.RefreshToken != "" || !errors.Is(err, pgx.ErrNoRows) {
	//	_, err := tR.deleteTokenByUserID(ctx, token.UserID)
	//	if err != nil {
	//		return model.Token{}, fmt.Errorf(path+".deleteTokenByUserID, error: {%w}", err)
	//	}
	//} else if err != nil {
	//	return model.Token{}, fmt.Errorf(path+".getTokenByUserID, error: {%w}", err)
	//}

	sql, args, err := tR.Builder.
		Insert("public.tokens").
		Columns("refresh_token", "user_id").
		Values(token.RefreshToken, token.UserID).
		Suffix("RETURNING id,refresh_token,user_id").
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
func (tR *TokenRepository) getTokenByUserID(ctx context.Context, userID int) (model.Token, error) {
	path := "internal.repository.token.getTokenByUserID"

	sql, args, err := tR.Builder.
		Select("id,refresh_token, user_id").
		From("public.tokens").
		Where("user_id = ?", userID).
		ToSql()
	if err != nil {
		return model.Token{}, fmt.Errorf(path+".ToSql, error: {%w}", err)
	}
	var t model.Token
	err = tR.Pool.QueryRow(ctx, sql, args...).
		Scan(&t.ID, &t.RefreshToken, &t.UserID)
	if err != nil {
		return model.Token{}, fmt.Errorf(path+".QueryRow, error: {%w}", err)
	}
	return t, nil
}

func (tR *TokenRepository) deleteTokenByUserID(ctx context.Context, userID int) (int, error) {
	path := "internal.repository.token.deleteTokenByUserID"
	sql, args, err := tR.Builder.
		Delete("public.tokens").
		Where("user_id = ?", userID).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf(path+".ToSql, error: {%w}", err)
	}

	var t int
	err = tR.Pool.QueryRow(ctx, sql, args...).Scan(&t)
	if err != nil {
		return 0, fmt.Errorf(path+".QueryRow, error: {%w}", err)
	}
	return t, nil
}
