package system

import (
	 pc "github.com/darrennong/pc-go"
)

// NewSetalimits sets the account limits. Requires signature from `eosio@active` account.
func NewSetalimits(account pc.AccountName, ramBytes, netWeight, cpuWeight int64) *pc.Action {
	a := &pc.Action{
		Account: AN("eosio"),
		Name:    ActN("setalimit"),
		Authorization: []pc.PermissionLevel{
			{Actor: pc.AccountName("eosio"), Permission: PN("active")},
		},
		ActionData: pc.NewActionData(Setalimits{
			Account:   account,
			RAMBytes:  ramBytes,
			NetWeight: netWeight,
			CPUWeight: cpuWeight,
		}),
	}
	return a
}

// Setalimits represents the `eosio.system::setalimit` action.
type Setalimits struct {
	Account   pc.AccountName `json:"account"`
	RAMBytes  int64           `json:"ram_bytes"`
	NetWeight int64           `json:"net_weight"`
	CPUWeight int64           `json:"cpu_weight"`
}
