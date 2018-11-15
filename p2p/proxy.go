package p2p

import (
	"fmt"

	"log"

	"github.com/darrennong/pc-go"
)

type Proxy struct {
	Peer1                       *Peer
	Peer2                       *Peer
	handlers                    []Handler
	waitingOriginHandShake      bool
	waitingDestinationHandShake bool
}

func NewProxy(peer1 *Peer, peer2 *Peer) *Proxy {
	return &Proxy{
		Peer1: peer1,
		Peer2: peer2,
	}
}

func (p *Proxy) RegisterHandler(handler Handler) {
	p.handlers = append(p.handlers, handler)
}

func (p *Proxy) RegisterHandlers(handlers []Handler) {
	p.handlers = append(p.handlers, handlers...)
}

func (p *Proxy) read(sender *Peer, receiver *Peer, errChannel chan error) {
	for {

		log.Println("Waiting for packet")
		packet, err := sender.Read()
		log.Println("Received for packet")
		if err != nil {
			errChannel <- fmt.Errorf("read message from %s: %s", sender.Address, err)
			return
		}
		err = p.handle(packet, sender, receiver)
		if err != nil {
			errChannel <- err
		}
	}
}

func (p *Proxy) handle(packet *pc.Packet, sender *Peer, receiver *Peer) error {

	_, err := receiver.Write(packet.Raw)
	if err != nil {
		return fmt.Errorf("handleDefault: %s", err)
	}

	switch m := packet.P2PMessage.(type) {
	case *pc.GoAwayMessage:
		return fmt.Errorf("handling message: go away: reason [%d]", m.Reason)
	}

	envelope := NewEnvelope(sender, receiver, packet)

	for _, handle := range p.handlers {
		handle.Handle(envelope)
	}

	return nil
}

func triggerHandshake(peer *Peer) error {
	fmt.Printf("Sending handshake [%s] to: %s\n", peer.handshakeInfo, peer.Address)
	return peer.SendHandshake(peer.handshakeInfo)
}

func (p *Proxy) ConnectAndStart() error {

	log.Println("Connecting and starting proxy")

	errorChannel := make(chan error)

	peer1ReadyChannel := p.Peer1.Connect(errorChannel)
	peer2ReadyChannel := p.Peer2.Connect(errorChannel)

	peer1Ready := false
	peer2Ready := false
	for {

		select {
		case <-peer1ReadyChannel:
			peer1Ready = true
		case <-peer2ReadyChannel:
			peer2Ready = true
		case err := <-errorChannel:
			return err
		}
		if peer1Ready && peer2Ready {
			break
		}
	}

	return p.Start()

}

func (p *Proxy) Start() error {

	log.Println("Starting readers")
	errorChannel := make(chan error)
	go p.read(p.Peer1, p.Peer2, errorChannel)
	go p.read(p.Peer2, p.Peer1, errorChannel)

	if p.Peer2.handshakeInfo != nil {

		err := triggerHandshake(p.Peer2)
		if err != nil {
			return fmt.Errorf("connect and start: trigger handshake: %s", err)
		}
	}

	log.Println("Started")
	return <-errorChannel
}
