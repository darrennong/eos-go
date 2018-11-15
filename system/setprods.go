package system

import (
	 pc "github.com/darrennong/pc-go"
	"github.com/darrennong/pc-go/ecc"
)

// NewSetPriv returns a `setpriv` action that lives on the
// `eosio.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `eos-bios` boot process by the
// `eosio.system` contract.
func NewSetProds(producers []ProducerKey) *pc.Action {
	a := &pc.Action{
		Account: AN("eosio"),
		Name:    ActN("setprods"),
		Authorization: []pc.PermissionLevel{
			{Actor: AN("eosio"), Permission: PN("active")},
		},
		ActionData: pc.NewActionData(SetProds{
			Schedule: producers,
		}),
	}
	return a
}

// SetProds is present in `eosio.bios` contract. Used only at boot time.
type SetProds struct {
	Schedule []ProducerKey `json:"schedule"`
}

type ProducerKey struct {
	ProducerName    pc.AccountName `json:"producer_name"`
	BlockSigningKey ecc.PublicKey   `json:"block_signing_key"`
}
