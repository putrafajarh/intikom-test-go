package v1Router

import (
	controller "intikom-test-go/controller/auth"

	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.RouterGroup) *gin.RouterGroup {

	auth := r.Group("/auth")
	{
		auth.POST("/login", controller.HandleLogin)
		auth.POST("/register", controller.HandleRegister)
		auth.POST("/refresh", controller.HandleRefresh)
	}

	return auth
}
