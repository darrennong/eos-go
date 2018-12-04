package system

import (
	pc "github.com/darrennong/pc-go"
)

func NewBidname(bidder, newname pc.AccountName, bid pc.Asset) *pc.Action {
	a := &pc.Action{
		Account: AN("potato"),
		Name:    ActN("bidname"),
		Authorization: []pc.PermissionLevel{
			{Actor: bidder, Permission: PN("active")},
		},
		ActionData: pc.NewActionData(Bidname{
			Bidder:  bidder,
			Newname: newname,
			Bid:     bid,
		}),
	}
	return a
}

// Bidname represents the `eosio.system_contract::bidname` action.
type Bidname struct {
	Bidder  pc.AccountName `json:"bidder"`
	Newname pc.AccountName `json:"newname"`
	Bid     pc.Asset       `json:"bid"` // specified in EOS
}
