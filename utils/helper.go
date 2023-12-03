package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func GenerateRefreshToken(id string, userName string, Role string) (string, error) {

	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	// Create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["username"] = userName
	claims["role"] = Role
	claims["exp"] = jwt.StandardClaims{ExpiresAt: 0}.ExpiresAt

	// Sign the token with the secret key
	secretKey := []byte(os.Getenv("JWT_REFRESH_SECRET_KEY"))

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateAccessToken(id, username, role string) (string, error) {
	// Set up the claims
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{
		"id":       id,
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	secretKey := []byte(os.Getenv("JWT_ACCESS_SECRET_KEY"))
	accessToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
