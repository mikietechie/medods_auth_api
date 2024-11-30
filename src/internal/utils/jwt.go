/*
Date Created		28 Nov 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Status				Looking for a job!
Description			A Golang Auth API
*/

package utils

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mikietechie/gocurrenciesapi/internal/config"
)

func GerateToken(str string, lifetime time.Duration) (string, error) {
	// Create a new JWT token with claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": str,                             // Subject (user identifier)
		"iss": config.SYS_NAME,                 // Issuer
		"exp": time.Now().Add(lifetime).Unix(), // Expiration time
		"iat": time.Now().Unix(),               // Issued at
	})
	tokenString, err := claims.SignedString([]byte(config.SECRET_KEY))
	if err != nil {
		return "", err
	}
	// log.Println("Token claims added\t:\n", claims)
	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	// Parse the token with the secret key
	var token *jwt.Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SECRET_KEY), nil
	})

	// Check for verification errors
	if err != nil {
		return token, err
	}

	// Check if the token is valid
	if !token.Valid {
		return token, errors.New("invalid token")
	}

	// Return the verified token
	return token, nil
}

func GetHeadersAuthBearerToken(c *gin.Context) string {
	authorizationHeader := c.GetHeader("Authorization")
	bearerToken := strings.Replace(authorizationHeader, "Bearer ", "", 1)
	return bearerToken
}
