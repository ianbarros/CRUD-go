package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ianbarros/CRUD-go/src/configuration/rest_err/logger"
	"github.com/ianbarros/CRUD-go/src/configuration/rest_err/validation"
	"github.com/ianbarros/CRUD-go/src/controller/model/request"
	"github.com/ianbarros/CRUD-go/src/model"
	"github.com/ianbarros/CRUD-go/src/model/service"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Starting CreateUser controller...",
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
	service := service.NewUserDomainService()
	if err := service.CreateUser(domain); err != nil {
		c.JSON(int(err.Code), err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"),
	)

	c.String(http.StatusOK, "")
}
