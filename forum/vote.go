package forum

import (
	 pc "github.com/darrennong/pc-go"
)

// NewVote is an action representing a simple vote to be broadcast
// through the chain network.
func NewVote(voter pc.AccountName, proposalName pc.Name, voteValue uint8, voteJSON string) *pc.Action {
	a := &pc.Action{
		Account: ForumAN,
		Name:    ActN("vote"),
		Authorization: []pc.PermissionLevel{
			{Actor: voter, Permission: pc.PermissionName("active")},
		},
		ActionData: pc.NewActionData(Vote{
			Voter:        voter,
			ProposalName: proposalName,
			Vote:         voteValue,
			VoteJSON:     voteJSON,
		}),
	}
	return a
}

// Vote represents the `eosio.forum::vote` action.
type Vote struct {
	Voter        pc.AccountName `json:"voter"`
	ProposalName pc.Name        `json:"proposal_name"`
	Vote         uint8           `json:"vote"`
	VoteJSON     string          `json:"vote_json"`
}
