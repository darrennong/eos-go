package msig

import (
	 pc "github.com/darrennong/pc-go"
)

// NewExec returns a `exec` action that lives on the
// `eosio.msig` contract.
func NewExec(proposer pc.AccountName, proposalName pc.Name, executer pc.AccountName) *pc.Action {
	return &pc.Action{
		Account: pc.AccountName("eosio.msig"),
		Name:    pc.ActionName("exec"),
		// TODO: double check in this package that the `Actor` is always the `proposer`..
		Authorization: []pc.PermissionLevel{
			{Actor: executer, Permission: pc.PermissionName("active")},
		},
		ActionData: pc.NewActionData(Exec{proposer, proposalName, executer}),
	}
}

type Exec struct {
	Proposer     pc.AccountName `json:"proposer"`
	ProposalName pc.Name        `json:"proposal_name"`
	Executer     pc.AccountName `json:"executer"`
}
