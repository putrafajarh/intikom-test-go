package controller

import (
	"intikom-test-go/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckResponse struct {
	Status  string  `json:"status"`
	Service Service `json:"service"`
}

type Service struct {
	Database string `json:"database"`
}

func HealthCheck(c *gin.Context) {
	response := HealthCheckResponse{
		Status: "Fail",
		Service: Service{
			Database: "Failed",
		},
	}

	db := database.GetDB()
	sqlDB, err := db.DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	err = sqlDB.Ping()
	if err == nil {
		response.Status = "OK"
		response.Service.Database = "Connected"
	}

	c.JSON(http.StatusOK, response)
}
