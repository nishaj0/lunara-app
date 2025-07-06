package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nishaj0/lunara-app/lunara-server/internal/pkg/env"
)

var jwtSecret = []byte(env.GetEnv("JWT_SECRET", "default_secret"))

func GenerateToken(userID, email string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "email":   email,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}