package routes

import (
	"github.com/Lexus123/lnurl/models"
	"github.com/Lexus123/lnurl/server/handlers"
)

/*
GetRoutes ...
*/
var GetRoutes = models.Routes{
	models.Route{
		Name:        "GetHomePage",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: handlers.GetPaymentRequest,
	},
}
