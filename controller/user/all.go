package controller

import (
	"intikom-test-go/controller"
	"intikom-test-go/repository"
	"intikom-test-go/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	userService := service.UserService{
		UserRepository: repository.NewUserRepository(),
	}

	users, err := userService.FindAll()
	if err != nil {
		controller.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	controller.ResponseSuccess(c, users)
}
