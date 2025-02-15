package controller

import (
	"intikom-test-go/controller"
	"intikom-test-go/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTaskById(c *gin.Context) {
	taskId := c.Param("id")
	taskIdUint, err := strconv.ParseUint(taskId, 10, 64)
	if err != nil {
		controller.ResponseNotFound(c, "Invalid task ID")
		return
	}

	task, err := service.NewTaskService().FindByUserTaskId(c.GetUint("user_id"), uint(taskIdUint))
	if err != nil {
		controller.ResponseNotFound(c, err.Error())
		return
	}

	controller.ResponseSuccess(c, task)
}
