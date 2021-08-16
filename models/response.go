package models

import (
	lnurl "github.com/fiatjaf/go-lnurl"
)

func NewPaymentResponse(pr string) lnurl.LNURLPayResponse2 {
	return lnurl.LNURLPayResponse2{
		PR:     pr,
		Routes: [][]lnurl.RouteInfo{},
		SuccessAction: &lnurl.SuccessAction{
			Tag:         "url",
			URL:         "https://node.hfs.pw",
			Description: "Thanks for donating",
		},
	}
}
