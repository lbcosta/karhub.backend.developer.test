package handler

import (
	"github.com/gofiber/fiber/v2"
	"karhub.backend.developer.test/src/config/database"
	"time"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

type HealthCheckResponse struct {
	Status         string `json:"status"`
	Version        string `json:"version"`
	Uptime         string `json:"uptime"`
	DatabaseStatus string `json:"database_status"`
}

type HealthHandler struct{}

func NewHealthHandler() HealthHandler {
	return HealthHandler{}
}

func (h HealthHandler) Handle(c *fiber.Ctx) error {
	databaseStatus := checkDatabaseConnection()

	return c.JSON(HealthCheckResponse{
		Status:         "OK",
		Version:        "1.0.0",
		Uptime:         time.Since(startTime).String(),
		DatabaseStatus: databaseStatus,
	})
}

func checkDatabaseConnection() string {
	err := database.PingPostgres()
	if err != nil {
		return "ERROR"
	}

	return "OK"
}
