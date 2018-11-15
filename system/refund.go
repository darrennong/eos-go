package system

import (
	 pc "github.com/darrennong/pc-go"
)

// NewRefund returns a `refund` action that lives on the
// `eosio.system` contract.
func NewRefund(owner pc.AccountName) *pc.Action {
	return &pc.Action{
		Account: AN("eosio"),
		Name:    ActN("refund"),
		Authorization: []pc.PermissionLevel{
			{Actor: owner, Permission: PN("active")},
		},
		ActionData: pc.NewActionData(Refund{
			Owner: owner,
		}),
	}
}

// Refund represents the `eosio.system::refund` action
type Refund struct {
	Owner pc.AccountName `json:"owner"`
}
