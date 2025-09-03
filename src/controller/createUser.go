package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ianbarros/CRUD-go/src/configuration/logger"
	"github.com/ianbarros/CRUD-go/src/configuration/validation"
	"github.com/ianbarros/CRUD-go/src/controller/model/request"
	"github.com/ianbarros/CRUD-go/src/model"
	"github.com/ianbarros/CRUD-go/src/view"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller...",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"),
		)
		restErr := validation.ValidateUserError(err)

		c.JSON(int(restErr.Code), restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)
	domainResult, err := uc.service.CreateUser(domain)
	if err != nil {
		logger.Error(
			"Error trying to call CreateUser service",
			err,
			zap.String("journey", "createUser"))
		c.JSON(int(err.Code), err)
		return
	}

	logger.Info(
		"CreateUser controller executed successfully",
		zap.String("userID", domainResult.GetID()),
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult,
	))
}
