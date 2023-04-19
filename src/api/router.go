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
	beerHandler   handler.BeerHandler
	healthHandler handler.HealthHandler
}

func NewRouter(beerHandler handler.BeerHandler, healthHandler handler.HealthHandler) Router {
	return Router{beerHandler: beerHandler, healthHandler: healthHandler}
}

func (r Router) Start(port string) error {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: RouterLogFormat,
	}))

	v1 := app.Group("/api/v1")
	v1.Get("/health", r.healthHandler.Handle)

	beer := v1.Group("/beer")
	beer.Get("/", r.beerHandler.HandleGetAll)
	beer.Post("/", r.beerHandler.HandleCreate)
	beer.Put("/:id", r.beerHandler.HandleUpdate)
	beer.Delete("/:id", r.beerHandler.HandleDelete)
	beer.Get("/style", r.beerHandler.HandleGetClosestBeerStyles)

	return app.Listen(port)
}
