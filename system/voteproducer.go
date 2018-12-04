package system

import "github.com/darrennong/pc-go"

// NewNonce returns a `nonce` action that lives on the
// `eosio.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `eos-bios` boot process by the
// `eosio.system` contract.
func NewVoteProducer(voter pc.AccountName, proxy pc.AccountName, producers ...pc.AccountName) *pc.Action {
	a := &pc.Action{
		Account: AN("potato"),
		Name:    ActN("voteproducer"),
		Authorization: []pc.PermissionLevel{
			{Actor: voter, Permission: PN("active")},
		},
		ActionData: pc.NewActionData(
			VoteProducer{
				Voter:     voter,
				Proxy:     proxy,
				Producers: producers,
			},
		),
	}
	return a
}

// VoteProducer represents the `eosio.system::voteproducer` action
type VoteProducer struct {
	Voter     pc.AccountName   `json:"voter"`
	Proxy     pc.AccountName   `json:"proxy"`
	Producers []pc.AccountName `json:"producers"`
}
