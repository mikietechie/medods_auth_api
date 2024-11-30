/*
Date Created		28 Nov 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Status				Looking for a job!
Description			A Golang Auth API
*/

package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikietechie/gocurrenciesapi/internal/utils"
)

func WithAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := utils.GetHeadersAuthBearerToken(c)
		verifiedToken, err := utils.VerifyToken(bearerToken)
		if err != nil {
			log.Println("WithAuth ", err)
			c.AbortWithError(http.StatusForbidden, err)
			return
		}
		c.Set("token", *verifiedToken)
		c.Next()
	}
}
