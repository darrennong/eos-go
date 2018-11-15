package system

import (
	 pc "github.com/darrennong/pc-go"
)

func NewSetRAM(maxRAMSize uint64) *pc.Action {
	a := &pc.Action{
		Account: AN("eosio"),
		Name:    ActN("setram"),
		Authorization: []pc.PermissionLevel{
			{AN("eosio"), pc.PermissionName("active")},
		},
		ActionData: pc.NewActionData(SetRAM{
			MaxRAMSize: maxRAMSize,
		}),
	}
	return a
}

// SetRAM represents the hard-coded `setram` action.
type SetRAM struct {
	MaxRAMSize uint64 `json:"max_ram_size"`
}
