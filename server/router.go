package server

import (
	"context"
	"net/http"

	"github.com/Lexus123/lndclient"
	"github.com/Lexus123/lnurl/models"
	"github.com/Lexus123/lnurl/server/handlers"
	"github.com/gorilla/mux"
)

/*
NewRouter creates a new router and needs LND Services to do so
*/
func NewRouter(lndServices *lndclient.GrpcLndServices) *mux.Router {
	ctx := context.TODO()
	router := mux.NewRouter().StrictSlash(true)

	// Define the GET requests
	getRoutes := []models.Route{
		{
			Name:        "GetLnurlPayPage",
			Method:      "GET",
			Pattern:     "/lnurl-pay",
			Queries:     []string{"amount", "{amount}"},
			HandlerFunc: handlers.Payment(ctx, lndServices),
		},
	}

	// Add all GET requests
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
