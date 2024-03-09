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

type UserRepository struct {
	*postgres.DB
}

func NewUserRepository(db *postgres.DB) *UserRepository {
	return &UserRepository{db}
}

func (uR *UserRepository) CreateUser(ctx context.Context, user model.User) (int, error) {
	path := "internal.repository.user.CreateUser"
	sql, args, err := uR.Builder.Insert("public.user").
		Into("public.user").
		Columns("email", "password").
		Values(user.Email, user.Password).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf(path+".ToSql, error: {%w}", err)
	}

	var id int
	err = uR.Pool.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == "23505" {
				return 0, custom_errros.ErrAlreadyExists
			}
		}
		return 0, fmt.Errorf(path+".QueryRow, error: {%w}", err)
	}

	return id, nil
}

func (uR *UserRepository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	path := "internal.repository.user.GetUserByEmail"
	sql, args, err := uR.Builder.Select("id,email,password").
		From("public.user").
		Where("email = ?", email).
		ToSql()
	if err != nil {
		return model.User{}, fmt.Errorf(path+".ToSql, error: {%w}", err)
	}

	var user model.User
	err = uR.Pool.QueryRow(ctx, sql, args...).
		Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			return model.User{}, err
		}
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, custom_errros.ErrUserNotFound
		}
		return model.User{}, fmt.Errorf(path+".QueryRow, error: {%w}", err)
	}
	return user, nil
}
