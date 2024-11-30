/*
Date Created		28 Nov 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Status				Looking for a job!
Description			A Golang Auth API
*/

package structs

type TokenPayload struct {
	GUID string `json:"guid"`
	IP   string `json:"ip"`
}

type LoginParameters struct {
	GUID string `json:"guid"`
	Ip   string `json:"ip"`
}

type RefreshParameters struct {
	Token   string `json:"token"`
	Refresh string `json:"refresh"`
	Ip      string `json:"ip"`
}

type Error400Response struct {
	Error interface{} `json:"error"`
}
