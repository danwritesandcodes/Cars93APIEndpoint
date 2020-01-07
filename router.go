package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

// newRouter uses the gorilla mux router to implement a request router and
// dispatcher for matching incoming requests to their respective handler.
// See https://github.com/gorilla/mux for more information.
func newRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
