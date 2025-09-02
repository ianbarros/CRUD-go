package service

import (
	"github.com/ianbarros/CRUD-go/src/configuration/rest_err"
	"github.com/ianbarros/CRUD-go/src/model"
)

func (*userDomainService) FindUser(string) (
	*model.UserDomainInterface,
	*rest_err.RestErr,
) {
	return nil, nil
}
