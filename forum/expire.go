package forum

import (
	 pc "github.com/darrennong/pc-go"
)

// NewExpire is an action to expire a proposal ahead of its natural death.
func NewExpire(proposer pc.AccountName, proposalName pc.Name) *pc.Action {
	a := &pc.Action{
		Account: ForumAN,
		Name:    ActN("expire"),
		Authorization: []pc.PermissionLevel{
			{Actor: proposer, Permission: pc.PermissionName("active")},
		},
		ActionData: pc.NewActionData(Expire{
			ProposalName: proposalName,
		}),
	}
	return a
}

// Expire represents the `eosio.forum::propose` action.
type Expire struct {
	ProposalName pc.Name `json:"proposal_name"`
}
