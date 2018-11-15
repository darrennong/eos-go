package msig

import (
	 pc "github.com/darrennong/pc-go"
)

// NewUnapprove returns a `unapprove` action that lives on the
// `eosio.msig` contract.
func NewUnapprove(proposer pc.AccountName, proposalName pc.Name, level pc.PermissionLevel) *pc.Action {
	return &pc.Action{
		Account:       pc.AccountName("eosio.msig"),
		Name:          pc.ActionName("unapprove"),
		Authorization: []pc.PermissionLevel{level},
		ActionData:    pc.NewActionData(Unapprove{proposer, proposalName, level}),
	}
}

type Unapprove struct {
	Proposer     pc.AccountName     `json:"proposer"`
	ProposalName pc.Name            `json:"proposal_name"`
	Level        pc.PermissionLevel `json:"level"`
}
