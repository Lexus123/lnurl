package models

import (
	lnurl "github.com/fiatjaf/go-lnurl"
)

func NewPaymentResponse(pr string) lnurl.LNURLPayResponse2 {
	return lnurl.LNURLPayResponse2{
		PR:     pr,
		Routes: [][]lnurl.RouteInfo{},
	}
}
