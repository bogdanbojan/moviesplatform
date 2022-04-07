package web

import (
	"flag"
	"github.com/bogdanbojan/moviesplatform/api"
	"github.com/bogdanbojan/moviesplatform/db"
	"net/http"
)

// Application is the main struct used throughout the api. It holds the logging, error handling and storage.
type Application struct {
	*api.ServiceLogger
	db.DataPuller
}

// StartServer starts the application server.
func (app *Application) StartServer() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: app.ErrorLog,
		Handler:  app.routes(),
	}

	app.InfoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	app.ErrorLog.Fatal(err)
}
