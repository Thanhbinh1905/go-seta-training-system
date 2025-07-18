package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

func Connect(databaseURL string) (*Queries, error) {
	pool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		return nil, err
	}

	dbPool = pool

	return New(dbPool), nil
}

// Close cleans up the DB connection pool
func Close() {
	if dbPool != nil {
		dbPool.Close()
	}
}
