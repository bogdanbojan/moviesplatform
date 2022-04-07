package web

import (
	"net/http"
)

// Permissions handles the validation and response regarding user permissions.
func (app *Application) Permissions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		app.ClientError(w, http.StatusMethodNotAllowed)
	}

	if !app.validatePermissionsURL(r.URL.Path) {
		app.ClientError(w, http.StatusBadRequest)
	} else {
		app.writePermissionsResponse(w, r.URL.Path)
	}

}

// Users handles the validation and response for the service/feature/permission route.
func (app *Application) Users(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		app.ClientError(w, http.StatusMethodNotAllowed)
	}

	if !app.validateUsersURL(r.URL.Path) {
		app.ClientError(w, http.StatusBadRequest)
	} else {
		app.writeUsersResponse(w, r.URL.Path)
	}

}
