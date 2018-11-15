package system

import (
	 pc "github.com/darrennong/pc-go"
)

// NewRemoveProducer returns a `rmvproducer` action that lives on the
// `eosio.system` contract.  This is to be called by the consortium of
// BPs, to oust a BP from its place.  If you want to unregister
// yourself as a BP, use `unregprod`.
func NewRemoveProducer(producer pc.AccountName) *pc.Action {
	return &pc.Action{
		Account: AN("eosio"),
		Name:    ActN("rmvproducer"),
		Authorization: []pc.PermissionLevel{
			{Actor: AN("eosio"), Permission: PN("active")},
		},
		ActionData: pc.NewActionData(RemoveProducer{
			Producer: producer,
		}),
	}
}

// RemoveProducer represents the `eosio.system::rmvproducer` action
type RemoveProducer struct {
	Producer pc.AccountName `json:"producer"`
}
