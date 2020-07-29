package http_utils

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const (
	// PathVarFormat format string to add vars to path
	PathVarFormat = "{%s}"
)

// Route defines a Route type simpler than the one
// defined in gorilla mux
type Route struct {
	Name        string
	Method      string
	Pattern     string
	QueryParams []string
	HandlerFunc http.HandlerFunc
}

// NewRouter Creates new router with prefix and handlers for routes specified
func NewRouter(prefix string, routes []Route) (r *mux.Router) {
	r = mux.NewRouter().StrictSlash(true)
	s := r.PathPrefix(prefix).Subrouter()
	for _, route := range routes {
		if len(route.QueryParams) > 0 {
			log.Debugf("registering route for %s with query params", route.Pattern)
			s.HandleFunc(route.Pattern, route.HandlerFunc).
				Methods(route.Method).
				Queries(route.QueryParams...).
				Name(route.Name)
		} else {
			log.Debugf("registering route for %s", route.Pattern)
			s.HandleFunc(route.Pattern, route.HandlerFunc).
				Methods(route.Method).
				Name(route.Name)
		}
	}

	return
}
