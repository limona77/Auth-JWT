package postgres

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/gookit/slog"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxPool interface {
	Close()
}
type DB struct {
	Builder squirrel.StatementBuilderType
	Pool    PgxPool
}

func New(url string) *DB {
	db := &DB{
		Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
	var err error
	db.Pool, err = pgxpool.New(context.Background(), url)
	if err != nil {
		slog.Fatal("can't connect to Postgres", err)
	}
	return db
}
