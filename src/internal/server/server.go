/*
Date Created		28 Nov 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Status				Looking for a job!
Description			A Golang Auth API
*/

package server

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mikietechie/gocurrenciesapi/docs"
	"github.com/mikietechie/gocurrenciesapi/internal/config"
	"github.com/mikietechie/gocurrenciesapi/internal/controllers"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func RunServer() {
	docs.SwaggerInfo.Title = "Gin Swagger"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.ReadDoc()

	server := gin.Default()
	// server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.Use(gin.Logger())
	server.Use(cors.Default())
	server.Use(gin.Recovery())

	// Routing
	api_router := server.Group("/api/v1")
	controllers.AuthRouter(*api_router.Group("/auth"))
	controllers.IndexRouter(*api_router.Group("/"))
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Run Server
	ADDRESS := "http://" + config.SERVER_ADDRESS
	log.Println(ADDRESS)
	log.Println(ADDRESS + "/docs/index.html")
	http.ListenAndServe(config.SERVER_ADDRESS, server)
}
