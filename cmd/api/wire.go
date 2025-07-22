//go:build wireinject
// +build wireinject

package main

import (
	"framework/apm"
	"framework/config"
	"framework/db"
	"framework/log"
	"framework/server/http"

	"github.com/google/wire"
)

func createApp() (*App, error) {
	wire.Build(
		wire.Struct(new(App), "*"),
		wire.NewSet(
			config.ProviderSet,
			db.ProviderSet,
			http.ProviderSet,
			log.ProviderSet,
			apm.ProviderSet,
		),
	)
	return &App{}, nil
}

type App struct {
	HTTPServer *http.HTTPServer
	Config     *config.Config
	Tracer     *apm.Tracer
}
