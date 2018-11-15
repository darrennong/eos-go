package token

import  pc "github.com/darrennong/pc-go"

func NewTransfer(from, to pc.AccountName, quantity pc.Asset, memo string) *pc.Action {
	return &pc.Action{
		Account: AN("eosio.token"),
		Name:    ActN("transfer"),
		Authorization: []pc.PermissionLevel{
			{Actor: from, Permission: PN("active")},
		},
		ActionData: pc.NewActionData(Transfer{
			From:     from,
			To:       to,
			Quantity: quantity,
			Memo:     memo,
		}),
	}
}

// Transfer represents the `transfer` struct on `eosio.token` contract.
type Transfer struct {
	From     pc.AccountName `json:"from"`
	To       pc.AccountName `json:"to"`
	Quantity pc.Asset       `json:"quantity"`
	Memo     string          `json:"memo"`
}
