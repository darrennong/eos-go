package system

import (
	 pc "github.com/darrennong/pc-go"
)

// NewSellRAM will sell at current market price a given number of
// bytes of RAM.
func NewSellRAM(account pc.AccountName, bytes uint64) *pc.Action {
	a := &pc.Action{
		Account: AN("eosio"),
		Name:    ActN("sellram"),
		Authorization: []pc.PermissionLevel{
			{Actor: account, Permission: pc.PermissionName("active")},
		},
		ActionData: pc.NewActionData(SellRAM{
			Account: account,
			Bytes:   bytes,
		}),
	}
	return a
}

// SellRAM represents the `eosio.system::sellram` action.
type SellRAM struct {
	Account pc.AccountName `json:"account"`
	Bytes   uint64          `json:"bytes"`
}
