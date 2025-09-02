package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ianbarros/CRUD-go/src/configuration/rest_err/logger"
	"github.com/ianbarros/CRUD-go/src/configuration/rest_err/validation"
	"github.com/ianbarros/CRUD-go/src/controller/model/request"
	"go.uber.org/zap"
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

	logger.Info("User created successfully",
		zap.String("journey", "createUser"),
	)
}
