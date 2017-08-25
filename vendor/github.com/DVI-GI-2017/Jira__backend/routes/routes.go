package routes

import (
	"fmt"
	"log"
	"net/http"
)

// Helper struct for defining routes
type Route struct {
	Name    string
	Pattern string
	Method  string
	Handler http.HandlerFunc
}

type Mux interface {
	Get(pattern string, handler http.HandlerFunc) error
	Post(pattern string, handler http.HandlerFunc) error
}

// Setup routes defined in this package
func SetupRoutes(m Mux) {
	addRoutesToMux(m, defaultRoutes)
}

// Slice of default routes will be resolved automatically
var defaultRoutes []Route

// Adds slice of routes to mux
func addRoutesToMux(m Mux, routes []Route) {
	for _, route := range routes {
		err := addRouteToMux(m, route)
		if err != nil {
			log.Panicf("can not add route %v: %v", route, err)
		}
	}
}

// Adds one route to mux
func addRouteToMux(m Mux, r Route) error {
	switch r.Method {
	case http.MethodGet:
		return m.Get(r.Pattern, r.Handler)
	case http.MethodPost:
		return m.Post(r.Pattern, r.Handler)
	}
	return fmt.Errorf("method '%s' not supported", r.Method)
}
