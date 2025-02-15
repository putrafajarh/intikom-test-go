package controller

import (
	"intikom-test-go/controller"
	"intikom-test-go/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {
	tasks, err := service.NewTaskService().FindAll(c.GetUint("user_id"))
	if err != nil {
		controller.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	controller.ResponseSuccess(c, tasks)
}
