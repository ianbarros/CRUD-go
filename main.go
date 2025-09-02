package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ianbarros/CRUD-go/src/configuration/database/mongodb"
	"github.com/ianbarros/CRUD-go/src/configuration/logger"
	"github.com/ianbarros/CRUD-go/src/controller"
	"github.com/ianbarros/CRUD-go/src/controller/routes"
	"github.com/ianbarros/CRUD-go/src/model/service"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting the application...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongodb.InitConnection()

	//init dependencies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
