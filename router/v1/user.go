package v1Router

import (
	controller "intikom-test-go/controller/user"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) *gin.RouterGroup {

	userRouter := r.Group("/users")
	{
		userRouter.GET("/", controller.GetAllUsers)

		userRouter.POST("/", controller.CreateUser)

		userRouter.GET("/:id", controller.GetUserById)

		userRouter.PUT("/:id", controller.UpdateUser)

		userRouter.DELETE("/:id", controller.DeleteUser)
	}
	return userRouter
}
