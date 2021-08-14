package server

import (
	"fmt"
	"log"
	"net/http"
)

/*
Host will init the routing and fire up a server
*/
func Host() {
	router := NewRouter()

	fmt.Print("Server running on port 8003")

	log.Fatal(http.ListenAndServe(":8003", router))
}
