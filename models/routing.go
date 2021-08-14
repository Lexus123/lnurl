package models

import (
	"net/http"
)

/*
Route defines a single route
*/
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Queries     []string
}

/*
Routes defines a slice of routes
*/
type Routes []Route
