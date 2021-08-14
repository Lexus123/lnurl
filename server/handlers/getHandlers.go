package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/btcsuite/btcutil"
	"github.com/lightninglabs/lndclient"
	"github.com/lightningnetwork/lnd/lnrpc/invoicesrpc"
	"github.com/lightningnetwork/lnd/lnwire"
)

func str2f(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

func f2a(f float64) (btcutil.Amount, error) {
	return btcutil.NewAmount(f)
}

func a2msat(a btcutil.Amount) lnwire.MilliSatoshi {
	return lnwire.NewMSatFromSatoshis(a)
}

/*
GetPaymentRequest ...
*/
func GetPaymentRequest(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()

	// Getting the amount
	// TODO: Create seperate func for it
	amount := r.FormValue("amount")

	f, err := str2f(amount)
	if err != nil {
		fmt.Printf("GetPaymentRequest (str2f) error: %v\n", err)
	}

	a, err := f2a(f)
	if err != nil {
		fmt.Printf("GetPaymentRequest (f2a) error: %v\n", err)
	}

	msat := a2msat(a)

	// Setup LND Services client
	conf := &lndclient.LndServicesConfig{
		LndAddress:  "http://localhost:10009",
		Network:     lndclient.NetworkMainnet,
		MacaroonDir: "~/.lnd/data/chain/bitcoin/mainnet",
		TLSPath:     "~/.lnd/tls.cert",
	}

	lndServices, err := lndclient.NewLndServices(conf)
	if err != nil {
		fmt.Printf("GetPaymentRequest (NewLndServices) error: %v\n", err)
	}

	// Create invoice
	invoice := &invoicesrpc.AddInvoiceData{
		Value:       msat,
		Expiry:      60,
		HodlInvoice: false,
	}

	_, pr, err := lndServices.Client.AddInvoice(ctx, invoice)
	if err != nil {
		fmt.Printf("GetPaymentRequest (AddInvoice) error: %v\n", err)
	}

	output := []byte(pr)

	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
