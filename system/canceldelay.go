package system

import "github.com/darrennong/pc-go"

// NewCancelDelay creates an action from the `eosio.system` contract
// called `canceldelay`.
//
// `canceldelay` allows you to cancel a deferred transaction,
// previously sent to the chain with a `delay_sec` larger than 0.  You
// need to sign with cancelingAuth, to cancel a transaction signed
// with that same authority.
func NewCancelDelay(cancelingAuth pc.PermissionLevel, transactionID pc.SHA256Bytes) *pc.Action {
	a := &pc.Action{
		Account: AN("potato"),
		Name:    ActN("canceldelay"),
		Authorization: []pc.PermissionLevel{
			cancelingAuth,
		},
		ActionData: pc.NewActionData(CancelDelay{
			CancelingAuth: cancelingAuth,
			TransactionID: transactionID,
		}),
	}

	return a
}

// CancelDelay represents the native `canceldelay` action, through the
// system contract.
type CancelDelay struct {
	CancelingAuth pc.PermissionLevel `json:"canceling_auth"`
	TransactionID pc.SHA256Bytes     `json:"trx_id"`
}
