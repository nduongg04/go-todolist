package utils

import (
	"os"
	"time"
	"todolist-api/internal/config"
	"todolist-api/pkg/errors"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ACCESS_TOKEN_EXPIRATION_TIME, _  = time.ParseDuration(config.GetEnvOrDefault("ACCESS_TOKEN_EXPIRATION_TIME", "15m")) // "15m", "1h", etc.
	REFRESH_TOKEN_EXPIRATION_TIME, _ = time.ParseDuration(config.GetEnvOrDefault("REFRESH_TOKEN_EXPIRATION_TIME", "1h"))
)

func GenerateAccessToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(ACCESS_TOKEN_EXPIRATION_TIME).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))
}

func GenerateRefreshToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(REFRESH_TOKEN_EXPIRATION_TIME).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("REFRESH_TOKEN_SECRET")))
}

func ValidateAccessToken(tokenString string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.NewUnauthorizedError("invalid token signing method")
		}

		return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
	})
	if err != nil {
		return nil, errors.NewUnauthorizedError("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, errors.NewUnauthorizedError("invalid token claims")
	}

	return claims, nil
}

func ValidateRefreshToken(tokenString string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.NewUnauthorizedError("invalid token signing method")
		}

		return []byte(os.Getenv("REFRESH_TOKEN_SECRET")), nil
	})

	if err != nil {
		return nil, errors.NewUnauthorizedError("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, errors.NewUnauthorizedError("invalid token claims")
	}

	return claims, nil
}
