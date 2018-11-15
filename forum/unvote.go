package forum

import (
	 pc "github.com/darrennong/pc-go"
)

// NewUnVote is an action representing the action to undoing a current vote
func NewUnVote(voter pc.AccountName, proposalName pc.Name) *pc.Action {
	a := &pc.Action{
		Account: ForumAN,
		Name:    ActN("unvote"),
		Authorization: []pc.PermissionLevel{
			{Actor: voter, Permission: pc.PermissionName("active")},
		},
		ActionData: pc.NewActionData(UnVote{
			Voter:        voter,
			ProposalName: proposalName,
		}),
	}
	return a
}

// UnVote represents the `eosio.forum::unvote` action.
type UnVote struct {
	Voter        pc.AccountName `json:"voter"`
	ProposalName pc.Name        `json:"proposal_name"`
}
