package forum

import (
	 pc "github.com/darrennong/pc-go"
)

// NewPost is an action representing a simple message to be posted
// through the chain network.
func NewPost(poster pc.AccountName, postUUID, content string, replyToPoster pc.AccountName, replyToPostUUID string, certify bool, jsonMetadata string) *pc.Action {
	a := &pc.Action{
		Account: ForumAN,
		Name:    ActN("post"),
		Authorization: []pc.PermissionLevel{
			{Actor: poster, Permission: pc.PermissionName("active")},
		},
		ActionData: pc.NewActionData(Post{
			Poster:          poster,
			PostUUID:        postUUID,
			Content:         content,
			ReplyToPoster:   replyToPoster,
			ReplyToPostUUID: replyToPostUUID,
			Certify:         certify,
			JSONMetadata:    jsonMetadata,
		}),
	}
	return a
}

// Post represents the `eosio.forum::post` action.
type Post struct {
	Poster          pc.AccountName `json:"poster"`
	PostUUID        string          `json:"post_uuid"`
	Content         string          `json:"content"`
	ReplyToPoster   pc.AccountName `json:"reply_to_poster"`
	ReplyToPostUUID string          `json:"reply_to_post_uuid"`
	Certify         bool            `json:"certify"`
	JSONMetadata    string          `json:"json_metadata"`
}
