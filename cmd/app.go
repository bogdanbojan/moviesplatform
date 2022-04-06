package main

import (
	"github.com/bogdanbojan/moviesplatform/api"
	"github.com/bogdanbojan/moviesplatform/api/web"
	"github.com/bogdanbojan/moviesplatform/db"
)

func Init() {
	slog := api.NewServiceLogger()
	store := db.NewStorage()
	App := &web.Application{
		slog,
		store,
	}
	App.StartServer()
}
