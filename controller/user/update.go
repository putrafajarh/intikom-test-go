package controller

import (
	"intikom-test-go/controller"
	"intikom-test-go/model"
	"intikom-test-go/repository"
	"intikom-test-go/service"
	"intikom-test-go/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		controller.ResponseNotFound(c, "User not found")
		return
	}

	var request model.UpdateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		errorFields, err := utils.ValidateError(c, err)
		if err != nil {
			controller.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		controller.ValidationError(c, errorFields)
		return
	}

	userService := service.UserService{
		UserRepository: repository.NewUserRepository(),
	}

	user, err := userService.FindById(uint(userId))
	if err != nil {
		controller.ResponseNotFound(c, "User not found")
		return
	}

	updatedUser, err := userService.Update(user, request)
	if err != nil {
		controller.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	controller.ResponseSuccessWithMessage(c, "User updated successfully", updatedUser)

}
