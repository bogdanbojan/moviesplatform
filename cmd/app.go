package main

import (
	"github.com/bogdanbojan/moviesplatform/api"
	"github.com/bogdanbojan/moviesplatform/api/web"
	"github.com/bogdanbojan/moviesplatform/db"
)

func Init() {
	slog := api.NewServiceLogger()
	storage := db.NewStorage()
	App := &web.Application{
		ServiceLogger: slog,
		DataPuller:    storage,
	}
	App.StartServer()
}
