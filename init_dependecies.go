package main

import (
	"github.com/ianbarros/CRUD-go/src/controller"
	"github.com/ianbarros/CRUD-go/src/model/repository"
	"github.com/ianbarros/CRUD-go/src/model/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
