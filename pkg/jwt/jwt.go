package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID string, userRole string, secret string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    userRole,                              // Hoặc "manager" tùy theo vai trò của người dùng
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // Token hết hạn sau 1 ngày
		"iat":     time.Now().Unix(),                     // Issued at
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
