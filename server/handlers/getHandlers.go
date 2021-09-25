package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Lexus123/lndclient"
	"github.com/Lexus123/lnurl/helpers"
	"github.com/Lexus123/lnurl/models"
	"github.com/lightningnetwork/lnd/lnrpc/invoicesrpc"
	"github.com/lightningnetwork/lnd/lnwire"
)

/*
retrieveAmount returns the amount of millisats which was requested
*/
func retrieveAmount(r *http.Request) lnwire.MilliSatoshi {
	queryParamAmount := r.FormValue("amount")

	f, err := helpers.Str2f(queryParamAmount)
	if err != nil {
		fmt.Printf("retrieveAmount (str2f) error: %v\n", err)
	}

	a, err := helpers.F2a(f)
	if err != nil {
		fmt.Printf("retrieveAmount (f2a) error: %v\n", err)
	}

	return helpers.A2msat(a)
}

/*
Payment handles requests made to /lnurl-pay
*/
func Payment(ctx context.Context, lndServices *lndclient.GrpcLndServices) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Get the amount from the URL query params
		value := retrieveAmount(r)

		// Create the hash that needs to be present in the invoice
		s := "[[\"text/plain\", \"donate@theroadtonode.com\"],[\"text/identifier\", \"donate@theroadtonode.com\"]]"
		b := []byte(s)
		h := helpers.NewSHA256(b)

		// Create invoice configuration
		invoice := &invoicesrpc.AddInvoiceData{
			Value:           value,
			Expiry:          60,
			HodlInvoice:     false,
			DescriptionHash: h,
		}

		// Create the invoice
		_, pr, err := lndServices.Client.AddInvoice(ctx, invoice)
		if err != nil {
			fmt.Printf("GetPaymentRequest (AddInvoice) error: %v\n", err)
			http.Error(w, err.Error(), 500)
			return
		}

		// Create the response
		response := models.NewPaymentResponse(pr)

		output, err := json.Marshal(response)
		if err != nil {
			fmt.Printf("GetPaymentRequest (Marshal) error: %v\n", err)
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)
	}

	return http.HandlerFunc(fn)
}
