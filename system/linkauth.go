package system

import "github.com/darrennong/pc-go"

// NewLinkAuth creates an action from the `eosio.system` contract
// called `linkauth`.
//
// `linkauth` allows you to attach certain permission to the given
// `code::actionName`. With this set on-chain, you can use the
// `requiredPermission` to sign transactions for `code::actionName`
// and not rely on your `active` (which might be more sensitive as it
// can sign anything) for the given operation.
func NewLinkAuth(account, code pc.AccountName, actionName pc.ActionName, requiredPermission pc.PermissionName) *pc.Action {
	a := &pc.Action{
		Account: AN("potato"),
		Name:    ActN("linkauth"),
		Authorization: []pc.PermissionLevel{
			{account, pc.PermissionName("active")},
		},
		ActionData: pc.NewActionData(LinkAuth{
			Account:     account,
			Code:        code,
			Type:        actionName,
			Requirement: requiredPermission,
		}),
	}

	return a
}

// LinkAuth represents the native `linkauth` action, through the
// system contract.
type LinkAuth struct {
	Account     pc.AccountName    `json:"account"`
	Code        pc.AccountName    `json:"code"`
	Type        pc.ActionName     `json:"type"`
	Requirement pc.PermissionName `json:"requirement"`
}
