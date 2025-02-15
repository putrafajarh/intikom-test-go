package controller

import (
	"intikom-test-go/controller"
	"intikom-test-go/model"
	"intikom-test-go/service"
	"intikom-test-go/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateTask(c *gin.Context) {
	taskId := c.Param("id")
	taskIdUint, err := strconv.ParseUint(taskId, 10, 64)
	if err != nil {
		controller.ResponseNotFound(c, "Invalid task ID")
		return
	}

	var request model.UpdateTaskRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		errorFields, err := utils.ValidateError(c, err)
		if err != nil {
			controller.ResponseError(c, http.StatusBadRequest, err.Error())
			return
		}
		controller.ValidationError(c, errorFields)
		return
	}

	task, err := service.NewTaskService().Update(c.GetUint("user_id"), uint(taskIdUint), request)
	if err != nil {
		controller.ResponseError(c, http.StatusNotFound, err.Error())
		return
	}

	controller.ResponseSuccess(c, task)
}
