package handlers

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Lexus123/lnurl/models"
	"github.com/btcsuite/btcutil"
	"github.com/lightninglabs/lndclient"
	"github.com/lightningnetwork/lnd/lnrpc/invoicesrpc"
	"github.com/lightningnetwork/lnd/lnwire"
)

func str2f(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

func f2a(f float64) (btcutil.Amount, error) {
	return btcutil.NewAmount(f / 100000000000)
}

func a2msat(a btcutil.Amount) lnwire.MilliSatoshi {
	return lnwire.NewMSatFromSatoshis(a)
}

/*
retrieveAmount ...
*/
func retrieveAmount(r *http.Request) lnwire.MilliSatoshi {
	queryParamAmount := r.FormValue("amount")

	f, err := str2f(queryParamAmount)
	if err != nil {
		fmt.Printf("retrieveAmount (str2f) error: %v\n", err)
	}

	a, err := f2a(f)
	if err != nil {
		fmt.Printf("retrieveAmount (f2a) error: %v\n", err)
	}

	return a2msat(a)
}

// NewSHA256 ...
func NewSHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

/*
Payment ...
*/
func Payment(ctx context.Context, lndServices *lndclient.GrpcLndServices) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Getting the amount
		value := retrieveAmount(r)

		//
		ba := []byte("[[\\\"text/plain\\\", \\\"donate@theroadtonode.com\\\"],[\\\"text/identifier\\\", \\\"donate@theroadtonode.com\\\"]]")
		ba2 := NewSHA256(ba)

		// Create invoice configuration
		invoice := &invoicesrpc.AddInvoiceData{
			Value:           value,
			Expiry:          60,
			HodlInvoice:     false,
			DescriptionHash: ba2,
		}

		// Create the invoice
		// "[[\\"text/plain\\", \\"donate@theroadtonode.com\\"],[\\"text/identifier\\", \\"donate@theroadtonode.com\\"]]"
		_, pr, err := lndServices.Client.AddInvoice(ctx, invoice)
		if err != nil {
			fmt.Printf("GetPaymentRequest (AddInvoice) error: %v\n", err)
			http.Error(w, err.Error(), 500)
			return
		}

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
