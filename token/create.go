package token

import  pc "github.com/darrennong/pc-go"

func NewCreate(issuer pc.AccountName, maxSupply pc.Asset) *pc.Action {
	return &pc.Action{
		Account: AN("eosio.token"),
		Name:    ActN("create"),
		Authorization: []pc.PermissionLevel{
			{Actor: AN("eosio.token"), Permission: PN("active")},
		},
		ActionData: pc.NewActionData(Create{
			Issuer:        issuer,
			MaximumSupply: maxSupply,
		}),
	}
}

// Create represents the `create` struct on the `eosio.token` contract.
type Create struct {
	Issuer        pc.AccountName `json:"issuer"`
	MaximumSupply pc.Asset       `json:"maximum_supply"`
}
