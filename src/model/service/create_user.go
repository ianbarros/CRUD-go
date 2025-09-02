package service

import (
	"github.com/ianbarros/CRUD-go/src/configuration/rest_err"
	"github.com/ianbarros/CRUD-go/src/configuration/rest_err/logger"
	"github.com/ianbarros/CRUD-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {

	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	return nil
}
