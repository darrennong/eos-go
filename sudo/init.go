package sudo

import  pc "github.com/darrennong/pc-go"

func init() {
	eos.RegisterAction(AN("eosio.sudo"), ActN("exec"), Exec{})
}

var AN = pc.AN
var ActN = pc.ActN
