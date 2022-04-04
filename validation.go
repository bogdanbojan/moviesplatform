package main

import (
	"path"
	"strings"
)

func formatForCheck(s string) string {
	lowercase := strings.ToLower(s)
	trimmed := strings.TrimSpace(lowercase)
	wospaces := strings.ReplaceAll(trimmed, " ", "")
	return wospaces
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

func checkPermissionsURL(url string) bool {
	const userPath = "/v1/user/"
	const service = "movies"

	urlClean := path.Clean(url)
	p, end := path.Split(urlClean)

	if p == userPath {
		return true
	}
	if end == service &&
		p[:9] == userPath &&
		strings.Count(p, "/") == 4 {
		return true
	}
	return false
}
