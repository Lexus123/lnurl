package main

import (
	"flag"

	"github.com/Lexus123/lnurl/server"
)

func main() {
	// Get the flags set by the user or use some default values
	host := flag.String("host", "localhost:10009", "URL of LND gRPC")
	macDir := flag.String("mac", "~/.lnd/data/chain/bitcoin/mainnet", "Path to the macaroon directory")
	tlsPath := flag.String("tls", "~/.lnd/tls.cert", "Path to the tls.cert file")

	flag.Parse()

	flags := server.Flags{
		Host: host,
		Mac:  macDir,
		TLS:  tlsPath,
	}

	server.Host(flags)
}
