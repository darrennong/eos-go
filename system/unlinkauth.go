package system

import "github.com/darrennong/pc-go"

// NewUnlinkAuth creates an action from the `eosio.system` contract
// called `unlinkauth`.
//
// `unlinkauth` detaches a previously set permission from a
// `code::actionName`. See `linkauth`.
func NewUnlinkAuth(account, code pc.AccountName, actionName pc.ActionName) *pc.Action {
	a := &pc.Action{
		Account: AN("potato"),
		Name:    ActN("unlinkauth"),
		Authorization: []pc.PermissionLevel{
			{account, pc.PermissionName("active")},
		},
		ActionData: pc.NewActionData(UnlinkAuth{
			Account: account,
			Code:    code,
			Type:    actionName,
		}),
	}

	return a
}

// UnlinkAuth represents the native `unlinkauth` action, through the
// system contract.
type UnlinkAuth struct {
	Account pc.AccountName `json:"account"`
	Code    pc.AccountName `json:"code"`
	Type    pc.ActionName  `json:"type"`
}
