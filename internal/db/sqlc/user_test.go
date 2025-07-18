package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	arg := CreateUserParams{
		Username:     "testuser",
		Email:        "testuser@gmail.com",
		Role:         "member",
		PasswordHash: "hashedpassword",
	}

	user, err := testQueries.CreateUser(ctx, arg)
	require.NoError(t, err)

	require.NoError(t, err)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Role, user.Role)
	require.NotEmpty(t, user.PasswordHash)
	require.NotZero(t, user.CreatedAt)
}
