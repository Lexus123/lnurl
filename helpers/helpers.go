package helpers

import (
	"crypto/sha256"
	"strconv"

	"github.com/btcsuite/btcutil"
	"github.com/lightningnetwork/lnd/lnwire"
)

func Str2f(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

func F2a(f float64) (btcutil.Amount, error) {
	return btcutil.NewAmount(f / 100000000000)
}

func A2msat(a btcutil.Amount) lnwire.MilliSatoshi {
	return lnwire.NewMSatFromSatoshis(a)
}

func NewSHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}
