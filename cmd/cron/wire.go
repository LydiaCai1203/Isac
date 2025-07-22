//go:build wireinject
// +build wireinject

package main

import (
	"framework/config"
	"framework/cronjob"
	"framework/db"
	"framework/log"
	"framework/repo"

	"github.com/google/wire"
)

func createApp() (*App, error) {
	wire.Build(
		wire.Struct(new(App), "*"),
		wire.NewSet(
			config.ProviderSet,
			db.ProviderSet,
			log.ProviderSet,
			cronjob.ProviderSet,
			repo.ProviderSet,
		),
	)
	return &App{}, nil
}

type App struct {
	Config        *config.Config
	HelloWorldJob *cronjob.HelloWorldJob
}
