package main

import (
	"github.com/bogdanbojan/moviesplatform/api"
	"github.com/bogdanbojan/moviesplatform/api/web"
	"github.com/bogdanbojan/moviesplatform/db"
)

// Init initializes the application with the logging, error handling and the storage that it needs.
func Init() {
	servLog := api.NewServiceLogger()
	storage := db.NewStorage()
	storage.InitStorage(db.Embed)
	App := &web.Application{
		ServiceLogger: servLog,
		DataPuller:    storage,
	}
	App.StartServer()
}
