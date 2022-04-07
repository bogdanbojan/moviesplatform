package web

import (
	"encoding/json"
	"net/http"
	"path"
	"strings"
)

func (app *Application) writePermissionsResponse(w http.ResponseWriter, status int, url string) {
	u, s := extractData(url)
	if s == "" {
		user, _ := app.GetUser(u)
		err := app.writeJSON(w, status, user)
		if err != nil {
			return
		}
		return
	}

	data := app.constructServicePermissionData(u, s)
	err := app.writeJSON(w, status, data)
	if err != nil {
		return
	}

}

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

func (app *Application) constructServicePermissionData(user string, service string) map[string]interface{} {
	data := make(map[string]interface{})
	u, _ := app.GetUser(user)
	for p, k := range u.Permissions {
		perm := strings.Split(p, ".")
		if perm[0] == service {
			data[p] = k
		}
	}
	return data
}

func (app *Application) writeUsersResponse(w http.ResponseWriter, status int, url string) {
	sfp := constructUrlPermission(url)
	uu := app.constructUsersCollection(sfp)
	err := app.writeJSON(w, status, uu)
	if err != nil {
		return
	}
}

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

func constructUrlPermission(url string) string {
	urlClean := path.Clean(url)
	p, permission := path.Split(urlClean)
	p, feature := path.Split(p[:len(p)-1])
	p, service := path.Split(p[:len(p)-1])
	return service + "." + feature + "." + permission
}

func (app *Application) writeJSON(w http.ResponseWriter, status int, data interface{}) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(status)
	_, err = w.Write(js)
	if err != nil {
		return err
	}
	return nil
}
