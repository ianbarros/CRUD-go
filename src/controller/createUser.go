package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ianbarros/CRUD-go/src/configuration/rest_err/validation"
	"github.com/ianbarros/CRUD-go/src/controller/model/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s\n", err.Error())
		restErr := validation.ValidateUserError(err)

		c.JSON(int(restErr.Code), restErr)
		return
	}
}
