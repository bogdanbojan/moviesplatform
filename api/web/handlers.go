package web

import (
	"net/http"
)

// Permissions handles the collection of permissions for the user or service. It only supports GET.
func (app *Application) Permissions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		app.ClientError(w, http.StatusMethodNotAllowed)

	}

	if !validatePermissionsURL(r.URL.Path) {
		app.clientError(w, http.StatusBadRequest)
	}

	app.writePermissionsResponse(w, 200, r.URL.Path)

}

func (app *Application) Users(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		app.clientError(w, http.StatusMethodNotAllowed)
	}

	if !validateUsersURL(r.URL.Path) {
		app.clientError(w, http.StatusBadRequest)
	}

	app.writeUsersResponse(w, 200, r.URL.Path)

}
