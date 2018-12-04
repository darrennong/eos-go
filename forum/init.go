package forum

import  pc "github.com/darrennong/pc-go"

func init() {
	pc.RegisterAction(ForumAN, ActN("clnproposal"), CleanProposal{})
	pc.RegisterAction(ForumAN, ActN("expire"), Expire{})
	pc.RegisterAction(ForumAN, ActN("post"), Post{})
	pc.RegisterAction(ForumAN, ActN("propose"), Propose{})
	pc.RegisterAction(ForumAN, ActN("status"), Status{})
	pc.RegisterAction(ForumAN, ActN("unpost"), UnPost{})
	pc.RegisterAction(ForumAN, ActN("unvote"), UnVote{})
	pc.RegisterAction(ForumAN, ActN("vote"), Vote{})
}

var AN = pc.AN
var PN = pc.PN
var ActN = pc.ActN

var ForumAN = AN("eosforumrcpp")
