package controller

import (
	"intikom-test-go/controller"
	"intikom-test-go/model"
	"intikom-test-go/service"
	"intikom-test-go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var request model.CreateTaskRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		errorFields, err := utils.ValidateError(c, err)
		if err != nil {
			controller.ResponseError(c, http.StatusBadRequest, err.Error())
			return
		}
		controller.ValidationError(c, errorFields)
		return
	}

	task, err := service.NewTaskService().Create(c.GetUint("user_id"), request)
	if err != nil {
		controller.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	controller.ResponseSuccessWithMessage(c, "Task created successfully", task)
}
