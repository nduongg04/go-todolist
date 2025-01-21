package middleware

import (
	"context"
	"net/http"
	"strings"
	"todolist-api/internal/config"
	"todolist-api/pkg/utils"
)

func AuthMiddleWare(cfg *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header is required", http.StatusUnauthorized)
				return
			}
			bearerToken := strings.Split(authHeader, " ")
			if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
				http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
				return
			}

			claims, err := utils.ValidateAccessToken(bearerToken[1])
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "userID", claims["sub"])
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
