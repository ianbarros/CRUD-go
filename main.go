package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ianbarros/CRUD-go/src/configuration/rest_err/logger"
	"github.com/ianbarros/CRUD-go/src/controller/routes"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting the application...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
