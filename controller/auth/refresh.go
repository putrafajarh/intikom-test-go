package controller

import (
	"intikom-test-go/controller"
	"intikom-test-go/model"
	"intikom-test-go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRefresh(c *gin.Context) {
	var request model.RefreshRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		errorFields, err := utils.ValidateError(c, err)
		if err != nil {
			controller.ResponseError(c, http.StatusBadRequest, err.Error())
			return
		}
		controller.ValidationError(c, errorFields)
		return
	}

	accessToken, err := utils.GenerateAccessTokenFromRefreshToken(request.RefreshToken)
	if err != nil {
		controller.ResponseError(c, http.StatusUnauthorized, err.Error())
		return
	}

	controller.ResponseSuccess(c, gin.H{"access_token": accessToken})
}
