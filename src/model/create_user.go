package model

import (
	"github.com/ianbarros/CRUD-go/src/configuration/rest_err"
	"github.com/ianbarros/CRUD-go/src/configuration/rest_err/logger"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {

	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	return nil
}
