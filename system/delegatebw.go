package system

import (
	pc "github.com/darrennong/pc-go"
)

// NewDelegateBW returns a `delegatebw` action that lives on the
// `eosio.system` contract.
func NewDelegateBW(from, receiver pc.AccountName, stakeCPU, stakeNet pc.Asset, transfer bool) *pc.Action {
	return &pc.Action{
		Account: AN("potato"),
		Name:    ActN("delegatebw"),
		Authorization: []pc.PermissionLevel{
			{Actor: from, Permission: PN("active")},
		},
		ActionData: pc.NewActionData(DelegateBW{
			From:     from,
			Receiver: receiver,
			StakeNet: stakeNet,
			StakeCPU: stakeCPU,
			Transfer: pc.Bool(transfer),
		}),
	}
}

// DelegateBW represents the `eosio.system::delegatebw` action.
type DelegateBW struct {
	From     pc.AccountName `json:"from"`
	Receiver pc.AccountName `json:"receiver"`
	StakeNet pc.Asset       `json:"stake_net"`
	StakeCPU pc.Asset       `json:"stake_cpu"`
	Transfer pc.Bool        `json:"transfer"`
}
