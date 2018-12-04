package system

import (
	"github.com/darrennong/pc-go"
	"github.com/darrennong/pc-go/ecc"
)

// NewNewAccount returns a `newaccount` action that lives on the
// `eosio.system` contract.
func NewNewAccount(creator, newAccount pc.AccountName, publicKey ecc.PublicKey) *pc.Action {
	return &pc.Action{
		Account: AN("potato"),
		Name:    ActN("newaccount"),
		Authorization: []pc.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: pc.NewActionData(NewAccount{
			Creator: creator,
			Name:    newAccount,
			Owner: pc.Authority{
				Threshold: 1,
				Keys: []pc.KeyWeight{
					{
						PublicKey: publicKey,
						Weight:    1,
					},
				},
				Accounts: []pc.PermissionLevelWeight{},
			},
			Active: pc.Authority{
				Threshold: 1,
				Keys: []pc.KeyWeight{
					{
						PublicKey: publicKey,
						Weight:    1,
					},
				},
				Accounts: []pc.PermissionLevelWeight{},
			},
		}),
	}
}

// NewDelegatedNewAccount returns a `newaccount` action that lives on the
// `eosio.system` contract. It is filled with an authority structure that
// delegates full control of the new account to an already existing account.
func NewDelegatedNewAccount(creator, newAccount pc.AccountName, delegatedTo pc.AccountName) *pc.Action {
	return &pc.Action{
		Account: AN("potato"),
		Name:    ActN("newaccount"),
		Authorization: []pc.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: pc.NewActionData(NewAccount{
			Creator: creator,
			Name:    newAccount,
			Owner: pc.Authority{
				Threshold: 1,
				Keys:      []pc.KeyWeight{},
				Accounts: []pc.PermissionLevelWeight{
					pc.PermissionLevelWeight{
						Permission: pc.PermissionLevel{
							Actor:      delegatedTo,
							Permission: PN("active"),
						},
						Weight: 1,
					},
				},
			},
			Active: pc.Authority{
				Threshold: 1,
				Keys:      []pc.KeyWeight{},
				Accounts: []pc.PermissionLevelWeight{
					pc.PermissionLevelWeight{
						Permission: pc.PermissionLevel{
							Actor:      delegatedTo,
							Permission: PN("active"),
						},
						Weight: 1,
					},
				},
			},
		}),
	}
}

// NewCustomNewAccount returns a `newaccount` action that lives on the
// `eosio.system` contract. You can specify your own `owner` and
// `active` permissions.
func NewCustomNewAccount(creator, newAccount pc.AccountName, owner, active pc.Authority) *pc.Action {
	return &pc.Action{
		Account: AN("potato"),
		Name:    ActN("newaccount"),
		Authorization: []pc.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: pc.NewActionData(NewAccount{
			Creator: creator,
			Name:    newAccount,
			Owner:   owner,
			Active:  active,
		}),
	}
}

// NewAccount represents a `newaccount` action on the `eosio.system`
// contract. It is one of the rare ones to be hard-coded into the
// blockchain.
type NewAccount struct {
	Creator pc.AccountName `json:"creator"`
	Name    pc.AccountName `json:"name"`
	Owner   pc.Authority   `json:"owner"`
	Active  pc.Authority   `json:"active"`
}
