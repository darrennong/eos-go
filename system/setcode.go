package system

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	 pc "github.com/darrennong/pc-go"
)

func NewSetContract(account pc.AccountName, wasmPath, abiPath string) (out []*pc.Action, err error) {
	codeContent, err := ioutil.ReadFile(wasmPath)
	if err != nil {
		return nil, err
	}

	abiContent, err := ioutil.ReadFile(abiPath)
	if err != nil {
		return nil, err
	}

	var abiDef pc.ABI
	if err := json.Unmarshal(abiContent, &abiDef); err != nil {
		return nil, fmt.Errorf("unmarshal ABI file: %s", err)
	}

	abiPacked, err := pc.MarshalBinary(abiDef)
	if err != nil {
		return nil, fmt.Errorf("packing ABI: %s", err)
	}

	actions := []*pc.Action{
		{
			Account: AN("potato"),
			Name:    ActN("setcode"),
			Authorization: []pc.PermissionLevel{
				{account, pc.PermissionName("active")},
			},
			ActionData: pc.NewActionData(SetCode{
				Account:   account,
				VMType:    0,
				VMVersion: 0,
				Code:      pc.HexBytes(codeContent),
			}),
		},
		{
			Account: AN("potato"),
			Name:    ActN("setabi"),
			Authorization: []pc.PermissionLevel{
				{account, pc.PermissionName("active")},
			},
			ActionData: pc.NewActionData(SetABI{
				Account: account,
				ABI:     pc.HexBytes(abiPacked),
			}),
		},
	}
	return actions, nil
}

func NewSetCode(account pc.AccountName, wasmPath string) (out *pc.Action, err error) {
	codeContent, err := ioutil.ReadFile(wasmPath)
	if err != nil {
		return nil, err
	}

	return &pc.Action{
		Account: AN("potato"),
		Name:    ActN("setcode"),
		Authorization: []pc.PermissionLevel{
			{account, pc.PermissionName("active")},
		},
		ActionData: pc.NewActionData(SetCode{
			Account:   account,
			VMType:    0,
			VMVersion: 0,
			Code:      pc.HexBytes(codeContent),
		}),
	}, nil
}

func NewSetABI(account pc.AccountName, abiPath string) (out *pc.Action, err error) {
	abiContent, err := ioutil.ReadFile(abiPath)
	if err != nil {
		return nil, err
	}

	var abiDef pc.ABI
	if err := json.Unmarshal(abiContent, &abiDef); err != nil {
		return nil, fmt.Errorf("unmarshal ABI file: %s", err)
	}

	abiPacked, err := pc.MarshalBinary(abiDef)
	if err != nil {
		return nil, fmt.Errorf("packing ABI: %s", err)
	}

	return &pc.Action{
		Account: AN("potato"),
		Name:    ActN("setabi"),
		Authorization: []pc.PermissionLevel{
			{account, pc.PermissionName("active")},
		},
		ActionData: pc.NewActionData(SetABI{
			Account: account,
			ABI:     pc.HexBytes(abiPacked),
		}),
	}, nil
}

// NewSetCodeTx is _deprecated_. Use NewSetContract instead, and build
// your transaction yourself.
func NewSetCodeTx(account pc.AccountName, wasmPath, abiPath string) (out *pc.Transaction, err error) {
	actions, err := NewSetContract(account, wasmPath, abiPath)
	if err != nil {
		return nil, err
	}
	return &pc.Transaction{Actions: actions}, nil
}

// SetCode represents the hard-coded `setcode` action.
type SetCode struct {
	Account   pc.AccountName `json:"account"`
	VMType    byte            `json:"vmtype"`
	VMVersion byte            `json:"vmversion"`
	Code      pc.HexBytes    `json:"code"`
}

// SetABI represents the hard-coded `setabi` action.
type SetABI struct {
	Account pc.AccountName `json:"account"`
	ABI     pc.HexBytes    `json:"abi"`
}
