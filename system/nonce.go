package system

import "github.com/darrennong/pc-go"

// NewNonce returns a `nonce` action that lives on the
// `eosio.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `eos-bios` boot process by the
// `eosio.system` contract.
func NewNonce(nonce string) *pc.Action {
	a := &pc.Action{
		Account:       AN("eosio"),
		Name:          ActN("nonce"),
		Authorization: []pc.PermissionLevel{
			//{Actor: AN("eosio"), Permission: PN("active")},
		},
		ActionData: pc.NewActionData(Nonce{
			Value: nonce,
		}),
	}
	return a
}
