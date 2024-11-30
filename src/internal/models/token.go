/*
Date Created		28 Nov 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Status				Looking for a job!
Description			A Golang Auth API
*/

package models

import (
	"fmt"

	"github.com/mikietechie/gocurrenciesapi/internal/config"
	"github.com/mikietechie/gocurrenciesapi/internal/utils"
	"gorm.io/gorm"
)

type RefreshToken struct {
	gorm.Model
	Ip   string `gorm:"size:255;not null;" json:"ip"`
	Hash string `gorm:"size:255;not null;unique" json:"hash"`
	GUID string `gorm:"size:255;not null;" json:"guid"`
	// Breached bool   `gorm:"default:false" json:"breached"`
}

func (token *RefreshToken) Blacklist() {
	fmt.Println("Sending warning email to user.")
	// token.Breached = false // and save!
}

func (token RefreshToken) GetHash() string {
	return utils.Decrypt(token.Hash)
}

func (token RefreshToken) GetAccessToken() (string, error) {
	return utils.GerateToken(token.GUID, config.JWT_ACCESS_TOKEN_LIFETIME)
}

func (token *RefreshToken) SetHash() {
	token.Hash = utils.Encrypt(token.Hash)
}

func (token *RefreshToken) BeforeCreate(tx *gorm.DB) (err error) {
	token.SetHash()
	return
}
