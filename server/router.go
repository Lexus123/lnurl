package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Lexus123/lnurl/server/routes"
)

/*
NewRouter ...
*/
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Add GET requests
	for _, route := range routes.GetRoutes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
