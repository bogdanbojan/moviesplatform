package web

import (
	"path"
)

// validatePermissionsURL is responsible for checking that the user/service exists and the path is correct.
func (app *Application) validatePermissionsURL(url string) bool {
	const permissionPath = "/v1/user/"
	urlClean := path.Clean(url)

	p, end := path.Split(urlClean)
	if p == permissionPath && app.checkUser(end) {
		return true
	}

	// get rid of trailing slash
	up, u := path.Split(p[:len(p)-1])

	if up == permissionPath && app.checkUser(u) && checkService(end) {
		return true
	}

	return false
}

func (app *Application) checkUser(u string) bool {
	_, ok := app.GetUser(u)
	if !ok {
		return false
	}
	return true
}

func checkService(s string) bool {
	ss := map[string]struct{}{
		"blockbusters": {},
		"commercials":  {},
		"shorts":       {},
	}

	_, ok := ss[s]
	if !ok {
		return false
	}
	return true
}

func (app *Application) validateUsersURL(url string) bool {
	const userPath = "/v1/service/"

	urlClean := path.Clean(url)

	p, permission := path.Split(urlClean)
	p, feature := path.Split(p[:len(p)-1])
	p, service := path.Split(p[:len(p)-1])

	if p != userPath || !app.checkServiceRoute([]string{service, feature, permission}) {
		return false
	}
	return true
}

func (app *Application) checkServiceRoute(sfp []string) bool {
	return app.ValidateServicePermission(sfp)

}
