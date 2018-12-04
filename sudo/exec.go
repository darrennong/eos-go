package sudo

import (
	 pc "github.com/darrennong/pc-go"
)

// NewExec creates an `exec` action, found in the `eosio.sudo`
// contract.
//
// Given an `pc.Transaction`, call `pc.MarshalBinary` on it first,
// pass the resulting bytes as `pc.HexBytes` here.
func NewExec(executer pc.AccountName, transaction pc.HexBytes) *pc.Action {
	a := &pc.Action{
		Account: pc.AccountName("eosio.sudo"),
		Name:    pc.ActionName("exec"),
		Authorization: []pc.PermissionLevel{
			{Actor: executer, Permission: pc.PermissionName("active")},
		},
		ActionData: pc.NewActionData(Exec{
			Executer:    executer,
			Transaction: transaction,
		}),
	}
	return a
}

// Exec represents the `eosio.system::exec` action.
type Exec struct {
	Executer    pc.AccountName `json:"executer"`
	Transaction pc.HexBytes    `json:"trx"`
}
