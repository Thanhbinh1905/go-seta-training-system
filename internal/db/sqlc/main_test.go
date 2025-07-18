package db_test

import (
	"context"
	"log"
	"os"
	"testing"

	sqlc "github.com/Thanhbinh1905/seta-training-system/internal/db/sqlc" // ⚠️ Replace `your-module` bằng tên module trong go.mod
	"github.com/jackc/pgx/v5/pgxpool"
)

const dbSource = "postgresql://root:secret@localhost:5432/training-system?sslmode=disable"

var testQueries *sqlc.Queries

func TestMain(m *testing.M) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dbSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}
	testQueries = sqlc.New(pool)
	os.Exit(m.Run())
}
