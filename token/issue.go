package token

import  pc "github.com/darrennong/pc-go"

func NewIssue(to pc.AccountName, quantity pc.Asset, memo string) *pc.Action {
	return &pc.Action{
		Account: AN("pc.token"),
		Name:    ActN("issue"),
		Authorization: []pc.PermissionLevel{
			{Actor: AN("potato"), Permission: PN("active")},
		},
		ActionData: pc.NewActionData(Issue{
			To:       to,
			Quantity: quantity,
			Memo:     memo,
		}),
	}
}

// Issue represents the `issue` struct on the `pc.token` contract.
type Issue struct {
	To       pc.AccountName `json:"to"`
	Quantity pc.Asset       `json:"quantity"`
	Memo     string          `json:"memo"`
}
