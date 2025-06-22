package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/farahsrw/manageProject/config"
	"github.com/farahsrw/manageProject/models"
	"github.com/farahsrw/manageProject/utils"

	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(app *config.App) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				utils.RespondError(w, http.StatusUnauthorized, "No token provided")
				return
			}
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			claims := &models.Claims{}

			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return app.JWTKey, nil
			})

			if err != nil || !token.Valid {
				utils.RespondError(w, http.StatusUnauthorized, "Invalid token")
				return
			}

			ctx := context.WithValue(r.Context(), "claims", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
