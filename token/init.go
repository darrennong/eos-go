package token

import "github.com/darrennong/pc-go"

func init() {
	pc.RegisterAction(AN("pc.token"), ActN("transfer"), Transfer{})
	pc.RegisterAction(AN("pc.token"), ActN("issue"), Issue{})
	pc.RegisterAction(AN("pc.token"), ActN("create"), Create{})
}
