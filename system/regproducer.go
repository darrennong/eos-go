package system

import (
	 pc "github.com/darrennong/pc-go"
	"github.com/darrennong/pc-go/ecc"
)

// NewRegProducer returns a `regproducer` action that lives on the
// `eosio.system` contract.
func NewRegProducer(producer pc.AccountName, producerKey ecc.PublicKey, url string, location uint16) *pc.Action {
	return &pc.Action{
		Account: AN("eosio"),
		Name:    ActN("regproducer"),
		Authorization: []pc.PermissionLevel{
			{Actor: producer, Permission: PN("active")},
		},
		ActionData: pc.NewActionData(RegProducer{
			Producer:    producer,
			ProducerKey: producerKey,
			URL:         url,
			Location:    location,
		}),
	}
}

// RegProducer represents the `eosio.system::regproducer` action
type RegProducer struct {
	Producer    pc.AccountName `json:"producer"`
	ProducerKey ecc.PublicKey   `json:"producer_key"`
	URL         string          `json:"url"`
	Location    uint16          `json:"location"` // what,s the meaning of that anyway ?
}
