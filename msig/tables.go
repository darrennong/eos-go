package msig

import  pc "github.com/darrennong/pc-go"

type ProposalRow struct {
	ProposalName       pc.Name              `json:"proposal_name"`
	RequestedApprovals []pc.PermissionLevel `json:"requested_approvals"`
	ProvidedApprovals  []pc.PermissionLevel `json:"provided_approvals"`
	PackedTransaction  pc.HexBytes          `json:"packed_transaction"`
}
