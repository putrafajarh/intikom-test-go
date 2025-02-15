package controller

import (
	"intikom-test-go/controller"
	"intikom-test-go/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	taskId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		controller.ResponseNotFound(c, "invalid task ID")
		return
	}

	task, err := service.NewTaskService().Delete(c.GetUint("user_id"), uint(taskId))
	if err != nil {
		controller.ResponseNotFound(c, err.Error())
		return
	}

	controller.ResponseSuccessWithMessage(c, "Task deleted successfully", task)
}
