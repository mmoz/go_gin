package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")

		// Check if the Authorization header is present and starts with "Bearer"
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Split the Authorization header to get the token
		tokenString := strings.Split(authHeader, " ")[1]

		secretKey := []byte(os.Getenv("JWT_ACCESS_SECRET_KEY"))
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verify the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// Provide the secret key used for signing
			return []byte(secretKey), nil
		})

		// Check if there's an error or if the token is invalid
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Set the user information in the Gin context for further processing
		c.Set("user", token.Claims.(jwt.MapClaims)["sub"])
		c.Set("role", token.Claims.(jwt.MapClaims)["role"])
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")

		// Check if the Authorization header is present and starts with "Bearer"
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Split the Authorization header to get the token
		tokenString := strings.Split(authHeader, " ")[1]

		secretKey := []byte(os.Getenv("JWT_ACCESS_SECRET_KEY"))
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verify the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// Provide the secret key used for signing
			return []byte(secretKey), nil
		})

		// Check if there's an error or if the token is invalid
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Check if the role is admin
		claims := token.Claims.(jwt.MapClaims)
		role := claims["role"].(string)
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Set the user information in the Gin context for further processing
		c.Set("user", claims["sub"])
		c.Set("role", claims["role"])
		c.Next()
	}
}
