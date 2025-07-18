package db

import (
	"context"
	"os"
	"testing"

	"github.com/Thanhbinh1905/seta-training-system/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

const dbSource = "postgresql://root:secret@localhost:5432/training-system?sslmode=disable"

var testQueries *Queries

func TestMain(m *testing.M) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dbSource)
	if err != nil {
		logger.Log.Fatal("failed to connect to database", zap.Error(err))
	}
	testQueries = New(pool)
	os.Exit(m.Run())
}
