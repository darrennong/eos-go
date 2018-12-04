package msig

import (
	"github.com/darrennong/pc-go"
)

func init() {
	pc.RegisterAction(AN("eosio.msig"), ActN("propose"), &Propose{})
	pc.RegisterAction(AN("eosio.msig"), ActN("approve"), &Approve{})
	pc.RegisterAction(AN("eosio.msig"), ActN("unapprove"), &Unapprove{})
	pc.RegisterAction(AN("eosio.msig"), ActN("cancel"), &Cancel{})
	pc.RegisterAction(AN("eosio.msig"), ActN("exec"), &Exec{})
}

var AN = pc.AN
var PN = pc.PN
var ActN = pc.ActN
