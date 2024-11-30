/*
Date Created		28 Nov 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Status				Looking for a job!
Description			A Golang Auth API
*/

package main

import (
	"github.com/mikietechie/gocurrenciesapi/internal/initialize"
	"github.com/mikietechie/gocurrenciesapi/internal/server"
)

// @title                       Go Auth API
// @version                     0.1
// @description                 Go Auth API
// @contact.name                Mike Zinyoni
// @contact.email               mzinyoni7
// @securityDefinitions.apikey	Bearer
// @in                         	header
// @name                       	Authorization
// @description              	Type "Bearer " followed by a space and JWT token.
func main() {
	initialize.Init()
	defer initialize.Tear()
	server.RunServer()
}
