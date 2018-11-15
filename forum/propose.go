package forum

import (
	 pc "github.com/darrennong/pc-go"
)

// NewPropose is an action to submit a proposal for vote.
func NewPropose(proposer pc.AccountName, proposalName pc.Name, title string, proposalJSON string, expiresAt pc.JSONTime) *pc.Action {
	a := &pc.Action{
		Account: ForumAN,
		Name:    ActN("propose"),
		Authorization: []pc.PermissionLevel{
			{Actor: proposer, Permission: pc.PermissionName("active")},
		},
		ActionData: pc.NewActionData(Propose{
			Proposer:     proposer,
			ProposalName: proposalName,
			Title:        title,
			ProposalJSON: proposalJSON,
			ExpiresAt:    expiresAt,
		}),
	}
	return a
}

// Propose represents the `eosio.forum::propose` action.
type Propose struct {
	Proposer     pc.AccountName `json:"proposer"`
	ProposalName pc.Name        `json:"proposal_name"`
	Title        string          `json:"title"`
	ProposalJSON string          `json:"proposal_json"`
	ExpiresAt    pc.JSONTime    `json:"expires_at"`
}
