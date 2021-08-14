package main

import (
	"flag"

	"github.com/Lexus123/lnurl/server"
)

func main() {
	host := flag.String("host", "http://localhost:10009", "URL of LND gRPC")
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
