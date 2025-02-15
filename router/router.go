package router

import (
	"intikom-test-go/controller"
	"intikom-test-go/database"
	"intikom-test-go/utils"
	"time"

	v1Router "intikom-test-go/router/v1"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitRouter() *gin.Engine {
	debug := viper.GetBool("DEBUG")
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	database.DatabaseSeeder(database.GetDB(), 5)

	logger := utils.GetLogger()

	if err := utils.InitTranslate("en"); err != nil {
		logger.Fatal("Failed to initialize translator", zap.Error(err))
	}

	r := gin.New()
	r.Use(requestid.New())
	r.Use(cors.Default())

	r.Use(ginzap.GinzapWithConfig(logger, &ginzap.Config{
		TimeFormat: time.RFC3339,
		Context: ginzap.Fn(func(c *gin.Context) []zap.Field {
			fields := []zapcore.Field{}

			if requestID := c.Writer.Header().Get("X-Request-Id"); requestID != "" {
				fields = append(fields, zap.String("request_id", requestID))
			}

			return fields
		}),
	}))

	r.NoRoute(func(c *gin.Context) {
		controller.ResponseNotFound(c, "Route Not Found")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/health", controller.HealthCheck)
		v1Router.AuthRouter(v1)
		v1Router.UserRouter(v1)
		v1Router.TaskRouter(v1)
	}

	return r
}
