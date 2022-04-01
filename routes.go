package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/user/", app.Permissions)
	mux.HandleFunc("/v1/service/", app.Users)
	return mux
}
