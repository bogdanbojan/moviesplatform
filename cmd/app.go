package main

import (
	"github.com/bogdanbojan/moviesplatform/api"
	"github.com/bogdanbojan/moviesplatform/api/web"
	"github.com/bogdanbojan/moviesplatform/db"
)

// Init initializes the application with the logging, error handling and the storage that it needs.
func Init() {
	slog := api.NewServiceLogger()
	storage := db.NewStorage()
	App := &web.Application{
		ServiceLogger: slog,
		DataPuller:    storage,
	}
	App.StartServer()
}
