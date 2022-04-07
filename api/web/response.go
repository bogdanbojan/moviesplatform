package web

import (
	"encoding/json"
	"net/http"
	"path"
	"strings"
)

// writePermissionsResponse is a controller for writing the response of the Permissions Handler.
func (app *Application) writePermissionsResponse(w http.ResponseWriter, url string) {
	u, s := extractData(url)
	if s == "" {
		user, _ := app.GetUser(u)
		err := app.writeJSON(w, user)
		if err != nil {
			return
		}
		return
	}

	data := app.constructServicePermissionData(u, s)
	err := app.writeJSON(w, data)
	if err != nil {
		return
	}

}

// extractData gets the user and service from the verified url string.
func extractData(url string) (user string, service string) {
	const userPath = "/v1/user/"
	urlClean := path.Clean(url)
	p, end := path.Split(urlClean)
	if p == userPath {
		return end, ""
	}

	// get rid of trailing slash
	_, u := path.Split(p[:len(p)-1])
	return u, end

}

// constructServicePermissionData constructs the response for the service permissions of a particular user.
func (app *Application) constructServicePermissionData(userId string, service string) map[string]interface{} {
	data := make(map[string]interface{})
	u, _ := app.GetUser(userId)
	for p, k := range u.Permissions {
		perm := strings.Split(p, ".")
		if perm[0] == service {
			data[p] = k
		}
	}
	return data
}

// writeUsersResponse is a controller which handles the response for the v1/service route.
func (app *Application) writeUsersResponse(w http.ResponseWriter, url string) {
	sfp := constructUrlPermission(url)
	uu := app.constructUsersCollection(sfp)
	err := app.writeJSON(w, uu)
	if err != nil {
		return
	}
}

// constructUsersCollection constructs the response with all the users that have the same permission.
func (app *Application) constructUsersCollection(sfp string) map[string]string {
	uu := make(map[string]string)
	for n, u := range app.GetUsers() {
		for p := range u.Permissions {
			if p == sfp {
				uu[n] = u.Name
				break
			}
		}
	}
	return uu
}

// constructUrlPermission constructs a string representing the service.feature.permission from the verified url.
func constructUrlPermission(url string) string {
	urlClean := path.Clean(url)
	p, permission := path.Split(urlClean)
	p, feature := path.Split(p[:len(p)-1])
	p, service := path.Split(p[:len(p)-1])
	return service + "." + feature + "." + permission
}

// writeJSON writes the JSON response with the various data that it is given.
func (app *Application) writeJSON(w http.ResponseWriter, data interface{}) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		return err
	}
	return nil
}
