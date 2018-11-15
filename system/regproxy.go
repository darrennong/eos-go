package system

import (
	 pc "github.com/darrennong/pc-go"
)

// NewRegProxy returns a `regproxy` action that lives on the
// `eosio.system` contract.
func NewRegProxy(proxy pc.AccountName, isProxy bool) *pc.Action {
	return &pc.Action{
		Account: AN("eosio"),
		Name:    ActN("regproxy"),
		Authorization: []pc.PermissionLevel{
			{Actor: proxy, Permission: PN("active")},
		},
		ActionData: pc.NewActionData(RegProxy{
			Proxy:   proxy,
			IsProxy: isProxy,
		}),
	}
}

// RegProxy represents the `eosio.system::regproxy` action
type RegProxy struct {
	Proxy   pc.AccountName `json:"proxy"`
	IsProxy bool            `json:"isproxy"`
}
