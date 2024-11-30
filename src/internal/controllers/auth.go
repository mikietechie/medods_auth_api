/*
Date Created		28 Nov 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Status				Looking for a job!
Description			A Golang Auth API
*/

package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mikietechie/gocurrenciesapi/internal/middleware"
	"github.com/mikietechie/gocurrenciesapi/internal/services"
	"github.com/mikietechie/gocurrenciesapi/internal/structs"
)

// Login         godoc
// @Summary      Login
// @Description  Gets Credentials and Returns Auth Token
// @Tags         Auth
// @Produce      json
// @Param        guid    path      string  true  "structs.LoginParameters GUID"
// @Success      200     {object}  structs.LoginResponse
// @Router       /api/v1/auth/login/{guid} [post]
func Login(c *gin.Context) {
	body := structs.LoginParameters{GUID: c.Param("guid"), Ip: c.ClientIP()}
	data, err := services.Login(body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, data)
}

// Refresh         godoc
// @Summary      Refresh
// @Description  Gets Credentials and Returns Auth Token
// @Tags         Auth
// @Produce      json
// @Param        Authorization  header      string  true  "Starts with Bearer "
// @Success      200      {object}  structs.RefreshResponse
// @Router       /api/v1/auth/refresh [post]
func Refresh(c *gin.Context) {
	bearerToken := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", 1)
	body := structs.RefreshParameters{Token: bearerToken}
	data, err := services.Refresh(body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, data)
}

// Refresh       godoc
// @Summary      Refresh
// @Description  Gets Credentials and Returns Auth Token
// @Tags         Auth
// @Produce      json
// @Param        Authorization  header      string  true  "Starts with Bearer "
// @Success      200      {object}  structs.SecureResponse
// @Router       /api/v1/auth/secure [get]
func Secure(c *gin.Context) {
	token := c.Value("token").(jwt.Token)
	fmt.Println(token.Claims)
	subject, err := token.Claims.GetSubject()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	expiration, err := token.Claims.GetExpirationTime()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	data := structs.SecureResponse{Subject: subject, Expiration: expiration.Time}
	c.IndentedJSON(http.StatusOK, data)
}

func AuthRouter(r gin.RouterGroup) {
	r.POST("/login/:guid", Login)
	r.POST("/refresh", Refresh)
	r.GET("/secure", middleware.WithAuth(), Secure)
}
