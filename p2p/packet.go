package p2p

import (
	"github.com/darrennong/pc-go"
)

type Envelope struct {
	Sender   *Peer
	Receiver *Peer
	Packet   *pc.Packet `json:"envelope"`
}

func NewEnvelope(sender *Peer, receiver *Peer, packet *pc.Packet) *Envelope {
	return &Envelope{
		Sender:   sender,
		Receiver: receiver,
		Packet:   packet,
	}
}
