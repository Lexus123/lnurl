package server

import (
	"net/http"

	"github.com/Lexus123/lnurl/models"
	"github.com/Lexus123/lnurl/server/handlers"
	"github.com/gorilla/mux"
	"github.com/lightninglabs/lndclient"
)

/*
NewRouter ...
*/
func NewRouter(lndServices *lndclient.GrpcLndServices) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	getRoutes := models.Routes{
		models.Route{
			Name:        "GetHomePage",
			Method:      "GET",
			Pattern:     "/lnurl-pay",
			Queries:     []string{"amount", "{amount}"},
			HandlerFunc: handlers.GetPaymentRequest(lndServices),
		},
	}

	// Add GET requests
	for _, route := range getRoutes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Queries(route.Queries...).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
