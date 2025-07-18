package db_test

import (
	"context"
	"testing"

	sqlc "github.com/Thanhbinh1905/seta-training-system/internal/db/sqlc"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	arg := sqlc.CreateUserParams{
		Username:     "testuser",
		Email:        "testuser@gmail.com",
		Role:         "member",
		PasswordHash: "hashedpassword",
	}

	user, err := testQueries.CreateUser(ctx, arg)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	if user.Username != arg.Username || user.Email != arg.Email || user.Role != arg.Role {
		t.Errorf("created user does not match input: got %v, want %v", user, arg)
	}

	if user.PasswordHash != arg.PasswordHash {
		t.Errorf("password hash does not match: got %s, want %s", user.PasswordHash, arg.PasswordHash)
	}

	if user.CreatedAt == (sqlc.User{}).CreatedAt {
		t.Error("created_at should not be zero value")
	}

}
