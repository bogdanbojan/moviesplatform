package main

import (
	"github.com/bogdanbojan/moviesplatform/api"
	"github.com/bogdanbojan/moviesplatform/api/web"
	"github.com/bogdanbojan/moviesplatform/db"
)

func Init() {
	db.InitStorage()
	slog := api.NewServiceLogger()
	App := &web.Application{
		slog,
	}
	App.StartServer()
}
