/*
Date Created		28 Nov 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Status				Looking for a job!
Description			A Golang Auth API
*/

package structs

import "time"

type LoginResponse struct {
	Token   string `json:"token"`
	Refresh string `json:"refresh"`
}

type RefreshResponse struct {
	Token string `json:"token"`
}

type SecureResponse struct {
	Subject    string    `json:"subject"`
	Expiration time.Time `json:"expiration"`
}
