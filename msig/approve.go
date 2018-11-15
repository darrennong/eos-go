package msig

import (
	 pc "github.com/darrennong/pc-go"
)

// NewApprove returns a `approve` action that lives on the
// `eosio.msig` contract.
func NewApprove(proposer pc.AccountName, proposalName pc.Name, level pc.PermissionLevel) *pc.Action {
	return &pc.Action{
		Account:       pc.AccountName("eosio.msig"),
		Name:          pc.ActionName("approve"),
		Authorization: []pc.PermissionLevel{level},
		ActionData:    pc.NewActionData(Approve{proposer, proposalName, level}),
	}
}

type Approve struct {
	Proposer     pc.AccountName     `json:"proposer"`
	ProposalName pc.Name            `json:"proposal_name"`
	Level        pc.PermissionLevel `json:"level"`
}
