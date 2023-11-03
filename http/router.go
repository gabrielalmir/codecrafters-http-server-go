package httphandler

import (
	"fmt"
	"regexp"
)

type Route struct {
	Path    string
	Handler func([]byte) string
}

var Routes = []Route{}

type Router struct {
	Routes []Route
}

func (r *Router) AddRoute(route Route) {
	r.Routes = append(r.Routes, route)
}

func (r *Router) Route(path string) (Route, bool) {
	fmt.Println("path: ", path)
	for _, route := range r.Routes {
		if ok, _ := regexp.MatchString(route.Path, path); ok {
			return route, true
		}
	}
	return Route{}, false
}
