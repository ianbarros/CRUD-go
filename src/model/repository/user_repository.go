package repository

import (
	"github.com/ianbarros/CRUD-go/src/configuration/rest_err"
	"github.com/ianbarros/CRUD-go/src/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)
}
