package system

import (
	 pc "github.com/darrennong/pc-go"
)

func NewSetRAMRate(bytesPerBlock uint16) *pc.Action {
	a := &pc.Action{
		Account: AN("potato"),
		Name:    ActN("setram"),
		Authorization: []pc.PermissionLevel{
			{AN("potato"), pc.PermissionName("active")},
		},
		ActionData: pc.NewActionData(SetRAMRate{
			BytesPerBlock: bytesPerBlock,
		}),
	}
	return a
}

// SetRAMRate represents the system contract's `setramrate` action.
type SetRAMRate struct {
	BytesPerBlock uint16 `json:"bytes_per_block"`
}
