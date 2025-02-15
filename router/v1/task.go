package v1Router

import (
	controller "intikom-test-go/controller/task"
	"intikom-test-go/router/middleware"

	"github.com/gin-gonic/gin"
)

func TaskRouter(r *gin.RouterGroup) *gin.RouterGroup {

	taskRouter := r.Group("/tasks")
	taskRouter.Use(middleware.AuthMiddleware())
	{
		taskRouter.GET("/", controller.GetAllTasks)

		taskRouter.POST("/", controller.CreateTask)

		taskRouter.GET("/:id", controller.GetTaskById)

		taskRouter.PUT("/:id", controller.UpdateTask)

		taskRouter.DELETE("/:id", controller.DeleteTask)
	}
	return taskRouter
}
