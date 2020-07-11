package http_utils

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	// Get http GET method string
	Get = "GET"

	// Post http POST method string
	Post = "POST"

	// Delete http DELETE method string
	Delete = "DELETE"

	// PathVarFormat format string to add vars to path
	PathVarFormat = "{%s}"
)

// Route defines a Route type simpler than the one
// defined in gorilla mux
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func NewRouter(routes []Route) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.HandleFunc(route.Pattern, route.HandlerFunc).
			Methods(route.Method).
			Name(route.Name)
	}

	return router
}
