package middlewares

import (
	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"

	"strings"

	"net/http"

	"os"

	"path/filepath"
)

func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		Claims, _ := extractJWTPayloads(c.GetHeader("Authorization")).(jwt.MapClaims)
		if Claims == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		if Claims["role"] == "users" || Claims["role"] == "admin" {
			c.Header("X-UserId", Claims["id"].(string))
			c.Next()
			return
		}
	}
}

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		Claims, _ := extractJWTPayloads(c.GetHeader("Authorization")).(jwt.MapClaims)
		if Claims == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		if Claims["role"] == "admin" {
			c.Header("X-UserId", Claims["id"].(string))
			c.Next()
			return
		}
	}
}

func extractJWTPayloads(jwt_token string) interface{} {
	secret_key, err := os.ReadFile(filepath.Join(".", "jwtRS256.key.pub"))
	if err != nil {
		panic(err)
	}

	tokenString := strings.Split(jwt_token, "Bearer ")[1]
	if tokenString == "" {
		return nil
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secret_key, nil
	})

	if err != nil {
		return nil
	}

	if Claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return Claims
	}

	return nil
}
