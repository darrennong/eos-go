package msig

import (
	 pc "github.com/darrennong/pc-go"
)

// NewCancel returns a `cancel` action that lives on the
// `eosio.msig` contract.
func NewCancel(proposer pc.AccountName, proposalName pc.Name, canceler pc.AccountName) *pc.Action {
	return &pc.Action{
		Account: pc.AccountName("eosio.msig"),
		Name:    pc.ActionName("cancel"),
		// TODO: double check in this package that the `Actor` is always the `proposer`..
		Authorization: []pc.PermissionLevel{
			{Actor: canceler, Permission: pc.PermissionName("active")},
		},
		ActionData: pc.NewActionData(Cancel{proposer, proposalName, canceler}),
	}
}

type Cancel struct {
	Proposer     pc.AccountName `json:"proposer"`
	ProposalName pc.Name        `json:"proposal_name"`
	Canceler     pc.AccountName `json:"canceler"`
}
