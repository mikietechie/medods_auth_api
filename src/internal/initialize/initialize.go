/*
Date Created		28 Nov 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Status				Looking for a job!
Description			A Golang Auth API
*/

package initialize

import (
	"log"
	"time"

	"github.com/mikietechie/gocurrenciesapi/internal/models"
)

func Init() {
	models.ConnectDb()
}

func Tear() {
	log.Println("Tearing down, will sleep for 10 seconds to allow go routines to finish")
	time.Sleep(time.Second * 10)
	models.DisonnectDb()
}
