package controller

import (
	"intikom-test-go/controller"
	"intikom-test-go/repository"
	"intikom-test-go/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		controller.ResponseNotFound(c, "Invalid user ID")
		return
	}

	userService := service.UserService{
		UserRepository: repository.NewUserRepository(),
	}

	user, err := userService.Delete(uint(userId))
	if err != nil {
		controller.ResponseNotFound(c, err.Error())
		return
	}

	controller.ResponseSuccessWithMessage(c, "User deleted successfully", user)
}
