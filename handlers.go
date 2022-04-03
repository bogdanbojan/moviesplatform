package main

import (
	"net/http"
	"strings"
)

func formatForCheck(s string) string {
	lowercase := strings.ToLower(s)
	trimmed := strings.TrimSpace(lowercase)
	wospaces := strings.ReplaceAll(trimmed, " ", "")
	return wospaces
}

func checkService(s string) bool {
	if s != "movies" {
		return false
	}
	return true
}

func checkPermission(p string) bool {
	pp := map[string]struct{}{
		"create":  {},
		"modify":  {},
		"comment": {},
		"rate":    {},
		"none":    {},
	}

	_, ok := pp[p]
	if !ok {
		return false
	}
	return true
}

func checkFeature(f string) bool {
	ff := map[string]struct{}{
		"directors": {},
		"producers": {},
		"users":     {},
		"guests":    {},
	}
	_, ok := ff[f]
	if !ok {
		return false
	}
	return true

}

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
