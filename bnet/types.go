package bnet

import (
	 pc "github.com/darrennong/pc-go"
	"github.com/darrennong/pc-go/ecc"
)

type BNetMessageType byte

const (
	HelloType                    BNetMessageType = iota // 0
	TransactionNoticeType                               // 1
	BlockNoticeType                                     // 2
	SignedBlockPointerType                              // 3
	PackedTransactionPointerType                        // 4
	PingType                                            // 5
	PongType                                            // 6
)

//
// Message types
//

type Hello struct {
	PeerID                   ecc.PublicKey     `json:"peer_id"`
	NetworkVersion           string            `json:"network_version"`
	User                     string            `json:"user"`
	Password                 string            `json:"password"`
	Agent                    string            `json:"agent"`
	ProtocolVersion          string            `json:"protocol_version"`
	ChainID                  pc.SHA256Bytes   `json:"chain_id"`
	RequestTransactions      pc.Bool          `json:"request_transactions"`
	LastIrreversibleBlockNum uint32            `json:"last_irr_block_num"`
	PendingBlockIDs          []pc.SHA256Bytes `json:"pending_block_ids"`
}

/**
 * From bnet_plugin.cpp's `trx_notice`:
 *
 * This message is sent upon successful speculative application of a transaction
 * and informs a peer not to send this message.
 */
type TransactionNotice struct {
	SignedTransactionIDs []pc.SHA256Bytes ///< hash of trx + sigs
}

/**
 * From bnet_plugin.cpp's `block_notice`:
 *
 * This message is sent upon successfully adding a transaction to the fork database
 * and informs the remote peer that there is no need to send this block.
 */
type BlockNotice struct {
	BlockIDs []pc.SHA256Bytes `json:"block_ids"`
}

type Ping struct {
	Sent                  pc.Tstamp      `json:"sent"`
	Code                  pc.SHA256Bytes `json:"code"`
	LastIrreversibleBlock uint32          `json:"lib"` /// last irreversible block
}

type Pong struct {
	Sent pc.Tstamp      `json:"sent"`
	Code pc.SHA256Bytes `json:"code"`
}

// Also use `pc.SignedBlock`
// Also use `pc.SignedTransaction`
var SignedBlock = pc.SignedTransaction{}
