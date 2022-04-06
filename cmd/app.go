package main

import (
	"github.com/bogdanbojan/moviesplatform/api"
	"github.com/bogdanbojan/moviesplatform/api/web"
	"github.com/bogdanbojan/moviesplatform/db"
)

type Application struct {
	*api.ServiceLogger
}

func start() {
	db.InitStorage()
	slog := api.NewServiceLogger()
	App := &web.Application{
		slog,
	}
	App.StartServer()
}
