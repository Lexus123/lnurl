package server

import (
	"log"
	"net/http"
)

/*
Host will init the routing and fire up a server
*/
func Host() {
	// Make a new Mux router
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8003", router))
}
