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
	App := &web.Application{
		ServiceLogger: servLog,
		DataPuller:    storage,
	}
	err := storage.InitStorage(db.Embed)
	if err != nil {
		App.ErrorLog.Fatal(err)
	}
	App.StartServer()
}
