package web

import (
	"path"
)

// validatePermissionsURL is responsible for checking that the user/service exists and the path is correct.
func (app *Application) validatePermissionsURL(url string) bool {
	const permissionPath = "/v1/user/"

	urlClean := path.Clean(url)
	if len(urlClean) < len(permissionPath) {
		return false
	}

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

// checkUser verifies if the user exists in the database.
func (app *Application) checkUser(userId string) bool {
	_, ok := app.GetUser(userId)
	if !ok {
		return false
	}
	return true
}

// checkService verifies that the service is a valid one.
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

// validateUsersURL verifies that the URL from the Users Handler is a valid one.
func (app *Application) validateUsersURL(url string) bool {
	const userPath = "/v1/service/"

	urlClean := path.Clean(url)
	if len(urlClean) < len(userPath) {
		return false
	}

	p, permission := path.Split(urlClean)
	p, feature := path.Split(p[:len(p)-1])
	p, service := path.Split(p[:len(p)-1])

	if p == userPath && app.checkServiceRoute([]string{service, feature, permission}) {
		return true
	}
	return false
}

// checkServiceRoute verifies that the service.feature.permission model is correct.
func (app *Application) checkServiceRoute(sfp []string) bool {
	return app.ValidateServicePermission(sfp)

}
