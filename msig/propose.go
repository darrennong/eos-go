package msig

import (
	 pc "github.com/darrennong/pc-go"
)

// NewPropose returns a `propose` action that lives on the
// `eosio.msig` contract.
func NewPropose(proposer pc.AccountName, proposalName pc.Name, requested []pc.PermissionLevel, transaction *pc.Transaction) *pc.Action {
	return &pc.Action{
		Account: pc.AccountName("eosio.msig"),
		Name:    pc.ActionName("propose"),
		Authorization: []pc.PermissionLevel{
			{Actor: proposer, Permission: pc.PermissionName("active")},
		},
		ActionData: pc.NewActionData(Propose{proposer, proposalName, requested, transaction}),
	}
}

type Propose struct {
	Proposer     pc.AccountName       `json:"proposer"`
	ProposalName pc.Name              `json:"proposal_name"`
	Requested    []pc.PermissionLevel `json:"requested"`
	Transaction  *pc.Transaction      `json:"trx"`
}
