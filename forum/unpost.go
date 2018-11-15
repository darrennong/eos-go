package forum

import (
	 pc "github.com/darrennong/pc-go"
)

// NewUnPost is an action undoing a post that is active
func NewUnPost(poster pc.AccountName, postUUID string) *pc.Action {
	a := &pc.Action{
		Account: ForumAN,
		Name:    ActN("post"),
		Authorization: []pc.PermissionLevel{
			{Actor: poster, Permission: pc.PermissionName("active")},
		},
		ActionData: pc.NewActionData(UnPost{
			Poster:   poster,
			PostUUID: postUUID,
		}),
	}
	return a
}

// UnPost represents the `eosio.forum::unpost` action.
type UnPost struct {
	Poster   pc.AccountName `json:"poster"`
	PostUUID string          `json:"post_uuid"`
}
