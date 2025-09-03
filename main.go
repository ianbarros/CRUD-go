package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ianbarros/CRUD-go/src/configuration/database/mongodb"
	"github.com/ianbarros/CRUD-go/src/configuration/logger"
	"github.com/ianbarros/CRUD-go/src/controller/routes"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting the application...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error connecting to database, error=%s \n", err.Error())
		return
	}

	userController := initDependencies(database)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
