package forum

import  pc "github.com/darrennong/pc-go"

func init() {
	eos.RegisterAction(ForumAN, ActN("clnproposal"), CleanProposal{})
	eos.RegisterAction(ForumAN, ActN("expire"), Expire{})
	eos.RegisterAction(ForumAN, ActN("post"), Post{})
	eos.RegisterAction(ForumAN, ActN("propose"), Propose{})
	eos.RegisterAction(ForumAN, ActN("status"), Status{})
	eos.RegisterAction(ForumAN, ActN("unpost"), UnPost{})
	eos.RegisterAction(ForumAN, ActN("unvote"), UnVote{})
	eos.RegisterAction(ForumAN, ActN("vote"), Vote{})
}

var AN = pc.AN
var PN = pc.PN
var ActN = pc.ActN

var ForumAN = AN("eosforumrcpp")
