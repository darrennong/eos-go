package system

import (
	pc "github.com/darrennong/pc-go"
)

func NewBuyRAM(payer, receiver pc.AccountName, pcQuantity uint64) *pc.Action {
	a := &pc.Action{
		Account: AN("potato"),
		Name:    ActN("buyram"),
		Authorization: []pc.PermissionLevel{
			{Actor: payer, Permission: PN("active")},
		},
		ActionData: pc.NewActionData(BuyRAM{
			Payer:    payer,
			Receiver: receiver,
			Quantity: pc.NewPCAsset(int64(pcQuantity)),
		}),
	}
	return a
}

// BuyRAM represents the `pcio.system::buyram` action.
type BuyRAM struct {
	Payer    pc.AccountName `json:"payer"`
	Receiver pc.AccountName `json:"receiver"`
	Quantity pc.Asset       `json:"quant"` // specified in pc
}
