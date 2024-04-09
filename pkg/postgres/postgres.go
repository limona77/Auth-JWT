package postgres

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxPool interface {
	Close()
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
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
		panic("can't connect to Postgres")
	}
	return db
}

func (db *DB) Close() {
	if db.Pool != nil {
		db.Pool.Close()
	}
}
