package controller

import (
	"intikom-test-go/controller"
	"intikom-test-go/model"
	"intikom-test-go/repository"
	"intikom-test-go/service"
	"intikom-test-go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRegister(c *gin.Context) {
	var request model.RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		errorFields, err := utils.ValidateError(c, err)
		if err != nil {
			controller.ResponseError(c, http.StatusBadRequest, err.Error())
			return
		}
		controller.ValidationError(c, errorFields)
		return
	}

	userService := service.UserService{
		UserRepository: repository.NewUserRepository(),
	}
	user, err := userService.Create(request)
	if err != nil {
		if err == repository.ErrEmailExists {
			controller.ValidationError(c, []utils.ErrorField{
				{
					Field:   "email",
					Message: err.Error(),
				},
			})
			return
		}
		controller.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	controller.ResponseSuccessWithMessage(c, "User registered successfully", user)
}
