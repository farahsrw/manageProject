package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/farahsrw/manageProject/config"
	"github.com/farahsrw/manageProject/models"
	"github.com/farahsrw/manageProject/utils"
)

func generateToken(app *config.App, username, xataID string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &models.Claims{
		Username: username,
		XataID:   xataID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(app.JWTKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Register
func Register(app *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds models.Credentials

		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		hashPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error in password")
			return
		}

		var xataID string
		err = app.DB.QueryRow("INSERT INTO \"users\" (username, password) VALUES ($1, $2) RETURNING xata_id", creds.Username, string(hashPassword)).Scan(&xataID)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error creating user")
			return
		}

		tokenString, err := generateToken(app, creds.Username, xataID)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error generating token")
			return
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(models.UserResponse{XataID: xataID, Username: creds.Username, Token: tokenString})
	}
}

// Login
func Login(app *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds models.Credentials

		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		var savedCreds models.Credentials
		var xataID string

		err = app.DB.QueryRow("SELECT xata_id, username, password FROM \"users\" WHERE username=$1", creds.Username).Scan(&xataID, &savedCreds.Username, &savedCreds.Password)
		if err != nil {
			if err == sql.ErrNoRows {
				utils.RespondError(w, http.StatusUnauthorized, "Invalid username or password")
				return
			}
			utils.RespondError(w, http.StatusInternalServerError, "Invalid request payload")
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(savedCreds.Password), []byte(creds.Password))
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Invalid username or password")
			return
		}

		tokenString, err := generateToken(app, creds.Username, xataID)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error generating token")
			return
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(models.UserResponse{XataID: xataID, Username: creds.Username, Token: tokenString})
	}
}
