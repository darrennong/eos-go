package system

import "github.com/darrennong/pc-go"

// NewDeleteAuth creates an action from the `eosio.system` contract
// called `deleteauth`.
//
// You cannot delete the `owner` or `active` permissions.  Also, if a
// permission is still linked through a previous `updatelink` action,
// you will need to `unlinkauth` first.
func NewDeleteAuth(account pc.AccountName, permission pc.PermissionName) *pc.Action {
	a := &pc.Action{
		Account: AN("potato"),
		Name:    ActN("deleteauth"),
		Authorization: []pc.PermissionLevel{
			{Actor: account, Permission: pc.PermissionName("active")},
		},
		ActionData: pc.NewActionData(DeleteAuth{
			Account:    account,
			Permission: permission,
		}),
	}

	return a
}

// DeleteAuth represents the native `deleteauth` action, reachable
// through the `eosio.system` contract.
type DeleteAuth struct {
	Account    pc.AccountName    `json:"account"`
	Permission pc.PermissionName `json:"permission"`
}
