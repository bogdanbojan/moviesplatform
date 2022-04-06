package web

import (
	"flag"
	"github.com/bogdanbojan/moviesplatform/api"

	"net/http"
)

type Application struct {
	*api.ServiceLogger
}

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
