package forum

import (
	 pc "github.com/darrennong/pc-go"
)

// CleanProposal is an action to flush proposal and allow RAM used by it.
func NewCleanProposal(cleaner pc.AccountName, proposalName pc.Name, maxCount uint64) *pc.Action {
	a := &pc.Action{
		Account: ForumAN,
		Name:    ActN("clnproposal"),
		Authorization: []pc.PermissionLevel{
			{Actor: cleaner, Permission: pc.PermissionName("active")},
		},
		ActionData: pc.NewActionData(CleanProposal{
			ProposalName: proposalName,
			MaxCount:     maxCount,
		}),
	}
	return a
}

// CleanProposal represents the `eosio.forum::clnproposal` action.
type CleanProposal struct {
	ProposalName pc.Name `json:"proposal_name"`
	MaxCount     uint64   `json:"max_count"`
}
