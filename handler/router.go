package handler

import (
	"regexp"
)

type Route struct {
	Path    string
	Method  string
	Handler func([]byte) string
}

var Routes = []Route{}

type Router struct {
	Routes []Route
}

func (r *Router) AddRoute(route Route) {
	r.Routes = append(r.Routes, route)
}

func (r *Router) Route(path string, method string) (Route, bool) {
	for _, route := range r.Routes {
		if ok, _ := regexp.MatchString(route.Path, path); ok && route.Method == method {
			return route, true
		}
	}
	return Route{}, false
}

func NotFound(r []byte) string {
	return "HTTP/1.1 404 Not Found\r\n\r\n"
}
