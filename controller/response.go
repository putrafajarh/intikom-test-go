package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseSuccessWithMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Status: "success",
		Data:   data,
	})
}

func ResponseError(c *gin.Context, status int, message string) {
	c.JSON(status, Response{
		Status:  "error",
		Message: message,
		Data:    nil,
	})
}

func ResponseNotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, Response{
		Status:  "error",
		Message: message,
		Data:    nil,
	})
}

func ValidationError(c *gin.Context, data interface{}) {
	c.JSON(http.StatusUnprocessableEntity, Response{
		Status:  "error",
		Message: "validation error",
		Data:    data,
	})
}
