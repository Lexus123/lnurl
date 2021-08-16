package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Lexus123/lndclient"
)

type Flags struct {
	Host, Mac, TLS *string
}

/*
Host will init the routing and fire up a server
*/
func Host(flags Flags) {
	fmt.Printf("\nHost: %v\n", *flags.Host)
	fmt.Printf("Mac: %v\n", *flags.Mac)
	fmt.Printf("TLS: %v\n", *flags.TLS)

	// Setup LND Services client
	conf := &lndclient.LndServicesConfig{
		LndAddress:  *flags.Host,
		Network:     lndclient.NetworkMainnet,
		MacaroonDir: *flags.Mac,
		TLSPath:     *flags.TLS,
	}

	lndServices, err := lndclient.NewLndServices(conf)
	if err != nil {
		fmt.Printf("GetPaymentRequest (NewLndServices) error: %v\n", err)
	} else {
		fmt.Println("Connected with LND")
	}

	router := NewRouter(lndServices)

	fmt.Print("Server running on port 8003\n")

	log.Fatal(http.ListenAndServe(":8003", router))
}
