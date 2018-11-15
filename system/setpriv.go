package system

import  pc "github.com/darrennong/pc-go"

// NewSetPriv returns a `setpriv` action that lives on the
// `eosio.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `eos-bios` boot process by the
// `eosio.system` contract.
func NewSetPriv(account pc.AccountName) *pc.Action {
	a := &pc.Action{
		Account: AN("eosio"),
		Name:    ActN("setpriv"),
		Authorization: []pc.PermissionLevel{
			{Actor: AN("eosio"), Permission: PN("active")},
		},
		ActionData: pc.NewActionData(SetPriv{
			Account: account,
			IsPriv:  pc.Bool(true),
		}),
	}
	return a
}

// SetPriv sets privileged account status. Used in the bios boot mechanism.
type SetPriv struct {
	Account pc.AccountName `json:"account"`
	IsPriv  pc.Bool        `json:"is_priv"`
}
