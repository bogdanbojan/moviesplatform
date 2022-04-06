package main

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
)

// Permissions handles the collection of permissions for the user or service. It only supports GET.
func (app *application) Permissions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		app.clientError(w, http.StatusMethodNotAllowed)
	}

	if !checkPermissionsURL(r.URL.Path) {
		app.clientError(w, http.StatusBadRequest)
	}

}

func (app *application) Users(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		app.clientError(w, http.StatusMethodNotAllowed)
	}
	url2 := r.URL.Path
	urlClean := path.Clean(url2)
	_, end := path.Split(urlClean)

	if end == "movies" {
		return
	}
	parsed, _ := url.ParseRequestURI(url2)
	fmt.Println(parsed.User, parsed.Scheme, parsed.Host, parsed.Redacted(), parsed.EscapedFragment(), parsed.Path, parsed.EscapedPath())
}
