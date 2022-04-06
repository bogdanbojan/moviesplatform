package web

import (
	"github.com/bogdanbojan/moviesplatform/db"
	"path"
)

// validatePermissionsURL is responsible for checking that the user/service exists and the path is correct.
func validatePermissionsURL(url string) bool {
	const permissionPath = "/v1/user/"

	urlClean := path.Clean(url)

	p, end := path.Split(urlClean)
	if p == permissionPath && checkUser(end) {
		return true
	}

	// get rid of trailing slash
	up, u := path.Split(p[:len(p)-1])

	if up == permissionPath && checkUser(u) && checkService(end) {
		return true
	}

	return false
}

func checkUser(u string) bool {
	_, ok := db.UsersCollection[u]
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

func validateUsersURL(url string) bool {
	const userPath = "/v1/service/"

	urlClean := path.Clean(url)

	p, permission := path.Split(urlClean)
	p, feature := path.Split(p[:len(p)-1])
	p, service := path.Split(p[:len(p)-1])

	if p != userPath && !checkServiceRoute([]string{service, feature, permission}) {
		return false
	}
	return true
}

type services map[string]features

type features map[string]permissions

type permissions map[string]struct{}

// TODO: Move this structure in db somewhere.
func checkServiceRoute(sfp []string) bool {
	blockbustersFeatures := features{
		"director":        permissions{"direct": {}, "instructActors": {}, "shoot": {}},
		"cinematographer": permissions{"shoot": {}, "changeLens": {}, "changeCamera": {}},
		"producer":        permissions{"changeBudget": {}, "changeSalary": {}, "addActor": {}},
	}

	commercialsFeatures := features{
		"artist":   permissions{"createConcept": {}, "creativitySwitch": {}},
		"producer": permissions{"getDeals": {}, "onboardPeople": {}},
		"manager":  permissions{"adviseBoard": {}, "cancelMeetings": {}},
	}

	shortsFeatures := features{
		"actor":    permissions{"act": {}, "readScript": {}, "cryOnCommand": {}},
		"investor": permissions{"invest": {}, "scandal": {}, "increaseBudget": {}},
		"director": permissions{"act": {}, "invest": {}, "direct": {}},
	}

	s := services{
		"blockbusters": blockbustersFeatures,
		"commercials":  commercialsFeatures,
		"shorts":       shortsFeatures,
	}

	_, ok := s[sfp[0]][sfp[1]][sfp[2]]

	if !ok {
		return false
	}
	return true
}
