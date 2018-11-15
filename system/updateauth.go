package system

import "github.com/darrennong/pc-go"

// NewUpdateAuth creates an action from the `eosio.system` contract
// called `updateauth`.
//
// usingPermission needs to be `owner` if you want to modify the
// `owner` authorization, otherwise `active` will do for the rest.
func NewUpdateAuth(account pc.AccountName, permission, parent pc.PermissionName, authority pc.Authority, usingPermission pc.PermissionName) *pc.Action {
	a := &pc.Action{
		Account: AN("eosio"),
		Name:    ActN("updateauth"),
		Authorization: []pc.PermissionLevel{
			{account, usingPermission},
		},
		ActionData: pc.NewActionData(UpdateAuth{
			Account:    account,
			Permission: permission,
			Parent:     parent,
			Auth:       authority,
		}),
	}

	return a
}

// UpdateAuth represents the hard-coded `updateauth` action.
//
// If you change the `active` permission, `owner` is the required parent.
//
// If you change the `owner` permission, there should be no parent.
type UpdateAuth struct {
	Account    pc.AccountName    `json:"account"`
	Permission pc.PermissionName `json:"permission"`
	Parent     pc.PermissionName `json:"parent"`
	Auth       pc.Authority      `json:"auth"`
}
