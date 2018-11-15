package forum

import (
	 pc "github.com/darrennong/pc-go"
)

// Status is an action to set a status update for a given account on the forum contract.
func NewStatus(account pc.AccountName, content string) *pc.Action {
	a := &pc.Action{
		Account: ForumAN,
		Name:    ActN("status"),
		Authorization: []pc.PermissionLevel{
			{Actor: account, Permission: pc.PermissionName("active")},
		},
		ActionData: pc.NewActionData(Status{
			Account: account,
			Content: content,
		}),
	}
	return a
}

// Status represents the `eosio.forum::status` action.
type Status struct {
	Account pc.AccountName `json:"account_name"`
	Content string          `json:"content"`
}
