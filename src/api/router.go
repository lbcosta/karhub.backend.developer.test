package api

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"karhub.backend.developer.test/src/api/v1/handler"
)

const (
	RouterLogFormat = "${time} | ${latency} | ${ip}:${port} | ${method} ${path} [${queryParams}] | ${status}\n"
)

type Router struct {
	healthHandler handler.HealthHandler
}

func NewRouter() Router {
	return Router{}
}

func (r Router) Start(port string) error {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: RouterLogFormat,
	}))

	v1 := app.Group("/api/v1")
	v1.Get("/health", r.healthHandler.Handle)

	return app.Listen(port)
}
