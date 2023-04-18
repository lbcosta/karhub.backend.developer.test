package main

import (
	"fmt"
	"karhub.backend.developer.test/src/api"
	"karhub.backend.developer.test/src/config"
	"karhub.backend.developer.test/src/config/provider"
)

func main() {
	app := provider.NewApp(provider.AppOptions{
		Port:   fmt.Sprintf(":%s", config.GetPort()),
		Router: api.NewRouter,
	})

	app.Run()
}
