/*
Date Created		28 Nov 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Status				Looking for a job!
Description			A Golang Auth API
*/

package services

import (
	"errors"

	"github.com/mikietechie/gocurrenciesapi/internal/config"
	"github.com/mikietechie/gocurrenciesapi/internal/models"
	"github.com/mikietechie/gocurrenciesapi/internal/structs"
	"github.com/mikietechie/gocurrenciesapi/internal/utils"
)

func Login(parameters structs.LoginParameters) (structs.LoginResponse, error) {
	var data structs.LoginResponse

	refreshToken, err := utils.GerateToken(parameters.GUID, config.JWT_REFRESH_TOKEN_LIFETIME)
	if err != nil {
		return data, err
	}
	token := models.RefreshToken{
		Ip:   parameters.Ip,
		GUID: parameters.GUID,
		Hash: refreshToken,
	}
	accessToken, err := token.GetAccessToken()
	if err != nil {
		return data, err
	}
	err = models.Db.Save(&token).Error
	if err != nil {
		return data, err
	}
	data = structs.LoginResponse{Token: accessToken, Refresh: refreshToken}
	return data, nil
}

func Refresh(parameters structs.RefreshParameters) (structs.RefreshResponse, error) {
	var data structs.RefreshResponse
	var token *models.RefreshToken
	models.Db.Model(&token).First(&token, "Hash = ?", utils.Encrypt(parameters.Refresh))
	if token == nil {
		return data, errors.New("Token not found.")
	}
	if token.Ip != parameters.Ip {
		go token.Blacklist()
		return data, errors.New("Token breached.")
	}
	accessToken, err := token.GetAccessToken()
	if err != nil {
		return data, err
	}
	data = structs.RefreshResponse{Token: accessToken}
	return data, nil
}
