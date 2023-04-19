package provider

import (
	"context"
	"go.uber.org/fx"
	"karhub.backend.developer.test/src/api/v1/handler"
	repository "karhub.backend.developer.test/src/api/v1/repository/gorm"
	"karhub.backend.developer.test/src/api/v1/service"
	"karhub.backend.developer.test/src/config/database"
)

type AppOptions struct {
	Port   string
	Router interface{}
}

type Api interface {
	Start(port string) error
}

func NewApp(options AppOptions) *fx.App {
	appPort := options.Port
	return fx.New(
		fx.Provide(fx.Annotate(options.Router, fx.As(new(Api)))),
		fx.Provide(providers()...),
		fx.Invoke(func(lifecycle fx.Lifecycle, api Api) {
			lifecycle.Append(
				fx.Hook{
					OnStart: func(context.Context) error {
						go api.Start(appPort)
						return nil
					},
				},
			)
		}),
		fx.NopLogger,
	)
}

func providers() []interface{} {
	return []interface{}{
		handler.NewBeerHandler,
		handler.NewHealthHandler,
		service.NewBeerService,
		repository.NewGormBeerRepository,
		database.NewPostgresDatabase,
	}
}
