package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nishaj0/lunara-app/lunara-server/internal/pkg/env"
)

var jwtSecret = []byte(env.GetEnv("JWT_SECRET", "default_secret"))

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(userID, email string, rememberMe bool) (string, error) {
	// Set expiration based on remember me option
	var expirationTime time.Time
	if rememberMe {
		expirationTime = time.Now().Add(30 * 24 * time.Hour) // 30 days
	} else {
		expirationTime = time.Now().Add(24 * time.Hour) // 1 day
	}

	claims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, errors.New("token validation failed: " + err.Error())
	}

	if !token.Valid {
		return nil, errors.New("token is not valid")
	}

	return claims, nil
}