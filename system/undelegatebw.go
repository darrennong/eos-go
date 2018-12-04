package system

import (
	 pc "github.com/darrennong/pc-go"
)

// NewUndelegateBW returns a `undelegatebw` action that lives on the
// `eosio.system` contract.
func NewUndelegateBW(from, receiver pc.AccountName, unstakeCPU, unstakeNet pc.Asset) *pc.Action {
	return &pc.Action{
		Account: AN("potato"),
		Name:    ActN("undelegatebw"),
		Authorization: []pc.PermissionLevel{
			{Actor: from, Permission: PN("active")},
		},
		ActionData: pc.NewActionData(UndelegateBW{
			From:     from,
			Receiver: receiver,
			UnstakeNet: unstakeNet,
			UnstakeCPU: unstakeCPU,
		}),
	}
}

// UndelegateBW represents the `eosio.system::undelegatebw` action.
type UndelegateBW struct {
	From         pc.AccountName `json:"from"`
	Receiver     pc.AccountName `json:"receiver"`
	UnstakeNet   pc.Asset       `json:"unstake_net_quantity"`
	UnstakeCPU   pc.Asset       `json:"unstake_cpu_quantity"`
}
