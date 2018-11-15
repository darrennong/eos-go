package system

import (
	 pc "github.com/darrennong/pc-go"
)

// NewUnregProducer returns a `unregprod` action that lives on the
// `eosio.system` contract.
func NewUnregProducer(producer pc.AccountName) *pc.Action {
	return &pc.Action{
		Account: AN("eosio"),
		Name:    ActN("unregprod"),
		Authorization: []pc.PermissionLevel{
			{Actor: producer, Permission: PN("active")},
		},
		ActionData: pc.NewActionData(UnregProducer{
			Producer: producer,
		}),
	}
}

// UnregProducer represents the `eosio.system::unregprod` action
type UnregProducer struct {
	Producer pc.AccountName `json:"producer"`
}
