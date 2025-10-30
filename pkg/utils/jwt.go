package utils

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/RajaSunrise/pajakku/config"
	"github.com/RajaSunrise/pajakku/internal/databases"
	"github.com/golang-jwt/jwt/v5"
)

var cfg = config.AppConfig

var jwtSecret = []byte(cfg.JWT.Secret) // Change this to a secure key

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	RoleID uint   `json:"role_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userID string, email string, roleID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		Email:  email,
		RoleID: roleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	// Store token in Redis with TTL
	err = databases.RDB.Set(context.Background(), tokenString, "valid", time.Until(expirationTime)).Err()
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Check if token exists in Redis
	val, err := databases.RDB.Get(context.Background(), tokenString).Result()
	if err != nil || val != "valid" {
		return nil, errors.New("token not found or invalid")
	}

	return claims, nil
}

func InvalidateToken(tokenString string) error {
	return databases.RDB.Del(context.Background(), tokenString).Err()
}

// GenerateRandomID generates a random numeric string with length between 8 and 10 digits
func GenerateRandomID() uint {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	length := rand.Intn(3) + 8 // 8, 9, or 10
	min := 1
	for i := 1; i < length; i++ {
		min *= 10
	}
	max := min*10 - 1
	id := rand.Intn(max-min+1) + min
	return uint(id)
}
