package system

import (
	pc "github.com/darrennong/pc-go"
)

// NewBuyRAMBytes will buy at current market price a given number of
// bytes of RAM, and grant them to the `receiver` account.
func NewBuyRAMBytes(payer, receiver pc.AccountName, bytes uint32) *pc.Action {
	a := &pc.Action{
		Account: AN("potato"),
		Name:    ActN("buyrambytes"),
		Authorization: []pc.PermissionLevel{
			{Actor: payer, Permission: pc.PermissionName("active")},
		},
		ActionData: pc.NewActionData(BuyRAMBytes{
			Payer:    payer,
			Receiver: receiver,
			Bytes:    bytes,
		}),
	}
	return a
}

// BuyRAMBytes represents the `pcio.system::buyrambytes` action.
type BuyRAMBytes struct {
	Payer    pc.AccountName `json:"payer"`
	Receiver pc.AccountName `json:"receiver"`
	Bytes    uint32          `json:"bytes"`
}
