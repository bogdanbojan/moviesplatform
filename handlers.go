package main

import "net/http"

func (app *application) Permissions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		app.clientError(w, http.StatusMethodNotAllowed)
	}
}

func (app *application) Users(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		app.clientError(w, http.StatusMethodNotAllowed)
	}
}
