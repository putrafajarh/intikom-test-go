package controller

import (
	"errors"
	"intikom-test-go/controller"
	"intikom-test-go/model"
	"intikom-test-go/repository"
	"intikom-test-go/router/middleware"
	"intikom-test-go/service"
	"intikom-test-go/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

func HandleLogin(c *gin.Context) {
	var request model.LoginRequest

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
	user, err := userService.FindByEmail(request.Email)
	if err != nil {
		controller.ResponseError(c, http.StatusUnauthorized, ErrInvalidCredentials.Error())
		return
	}

	if !utils.ComparePassword(user.Password, request.Password) {
		controller.ResponseError(c, http.StatusUnauthorized, ErrInvalidCredentials.Error())
		return
	}

	userId := strconv.Itoa(int(user.ID))
	accessToken, err := utils.GenerateAccessToken(userId)
	if err != nil {
		controller.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(userId)
	if err != nil {
		controller.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	controller.ResponseSuccess(c, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"token_type":    middleware.TokenHeaderPrefix,
	})
}
