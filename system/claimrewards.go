package system

import (
	pc "github.com/darrennong/pc-go"
)

// NewClaimRewards will buy at current market price a given number of
// bytes of RAM, and grant them to the `receiver` account.
func NewClaimRewards(owner pc.AccountName) *pc.Action {
	a := &pc.Action{
		Account: AN("pcio"),
		Name:    ActN("claimrewards"),
		Authorization: []pc.PermissionLevel{
			{Actor: owner, Permission: pc.PermissionName("active")},
		},
		ActionData: pc.NewActionData(ClaimRewards{
			Owner: owner,
		}),
	}
	return a
}

// ClaimRewards represents the `pcio.system::claimrewards` action.
type ClaimRewards struct {
	Owner pc.AccountName `json:"owner"`
}
