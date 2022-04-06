package main

import (
	"path"
)

func checkPermissionsURL(url string) bool {
	const userPath = "/v1/user/"

	urlClean := path.Clean(url)

	p, end := path.Split(urlClean)
	if p == userPath && checkUser(end) {
		return true
	}

	up, u := path.Split(p)
	if up == userPath && checkUser(u) && checkService(end) {
		return true
	}

	return false
}

// TODO: implement it so that it checks the values from the data store.
func checkUser(s string) bool {
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
