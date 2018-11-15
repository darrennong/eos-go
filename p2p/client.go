package p2p

import (
	"fmt"
	"log"
	"math"

	"time"

	"github.com/darrennong/pc-go"
)

type Client struct {
	peer        *Peer
	handlers    []Handler
	readTimeout time.Duration
	catchup     *Catchup
}

func NewClient(peer *Peer, catchup bool) *Client {
	client := &Client{
		peer: peer,
	}
	if catchup {
		client.catchup = &Catchup{
			headBlock: peer.handshakeInfo.HeadBlockNum,
		}
	}
	return client
}

func (c *Client) CloseConnection() error {
	if c.peer.connection == nil {
		return nil
	}
	return c.peer.connection.Close()
}

func (c *Client) SetReadTimeout(readTimeout time.Duration) {
	c.readTimeout = readTimeout
}

func (c *Client) RegisterHandler(handler Handler) {

	c.handlers = append(c.handlers, handler)
}

func (c *Client) read(peer *Peer, errChannel chan error) {
	for {
		packet, err := peer.Read()
		if err != nil {
			errChannel <- fmt.Errorf("read message from %s: %s", peer.Address, err)
			break
		}

		envelope := NewEnvelope(peer, peer, packet)
		for _, handle := range c.handlers {
			handle.Handle(envelope)
		}

		switch m := packet.P2PMessage.(type) {
		case *pc.GoAwayMessage:
			errChannel <- fmt.Errorf("GoAwayMessage reason [%s]: %s", m.Reason, err)

		case *pc.HandshakeMessage:
			if c.catchup == nil {
				m.NodeID = peer.NodeID
				m.P2PAddress = peer.Name
				err = peer.WriteP2PMessage(m)
				if err != nil {
					errChannel <- fmt.Errorf("HandshakeMessage: %s", err)
					break
				}
				fmt.Println("Handshake resent!")

			} else {

				c.catchup.originHeadBlock = m.HeadNum
				err := c.catchup.sendSyncRequest(peer)
				if err != nil {
					errChannel <- fmt.Errorf("handshake: sending sync request: %s", err)
				}
				c.catchup.IsCatchingUp = true
			}
		case *pc.NoticeMessage:
			if c.catchup != nil {
				pendingNum := m.KnownBlocks.Pending
				if pendingNum > 0 {
					c.catchup.originHeadBlock = pendingNum
					err = c.catchup.sendSyncRequest(peer)
					if err != nil {
						errChannel <- fmt.Errorf("noticeMessage: sending sync request: %s", err)
					}
				}
			}
		case *pc.SignedBlock:

			if c.catchup != nil {

				blockNum := m.BlockNumber()
				c.catchup.headBlock = blockNum
				if c.catchup.requestedEndBlock == blockNum {

					if c.catchup.originHeadBlock <= blockNum {
						fmt.Println("In sync with last handshake")
						blockID, err := m.BlockID()
						if err != nil {
							errChannel <- fmt.Errorf("getting block id: %s", err)
						}
						peer.handshakeInfo.HeadBlockNum = blockNum
						peer.handshakeInfo.HeadBlockID = blockID
						peer.handshakeInfo.HeadBlockTime = m.SignedBlockHeader.Timestamp.Time
						peer.SendHandshake(peer.handshakeInfo)
						fmt.Println("Sent new handshake with info:", peer.handshakeInfo)
					} else {
						err = c.catchup.sendSyncRequest(peer)
						if err != nil {
							errChannel <- fmt.Errorf("signed block: sending sync request: %s", err)
						}
					}
				}
			}
		}
	}
}

func (c *Client) Start() error {

	fmt.Println("Starting client")

	errorChannel := make(chan error, 1)

	readyChannel := c.peer.Connect(errorChannel)

	for {
		select {
		case <-readyChannel:
			go c.read(c.peer, errorChannel)
			if c.peer.handshakeInfo != nil {

				err := triggerHandshake(c.peer)
				if err != nil {
					return fmt.Errorf("connect and start: trigger handshake: %s", err)
				}
			}
		case err := <-errorChannel:
			log.Println("Start got ERROR:", err)
			return err
		}
	}
}

type Catchup struct {
	IsCatchingUp        bool
	requestedStartBlock uint32
	requestedEndBlock   uint32
	headBlock           uint32
	originHeadBlock     uint32
}

func (c *Catchup) sendSyncRequest(peer *Peer) error {

	c.IsCatchingUp = true

	delta := c.originHeadBlock - c.headBlock

	c.requestedStartBlock = c.headBlock
	c.requestedEndBlock = c.headBlock + uint32(math.Min(float64(delta), 100))

	fmt.Printf("Sending sync request to origin: start block [%d] end block [%d]\n", c.requestedStartBlock, c.requestedEndBlock)
	err := peer.SendSyncRequest(c.requestedStartBlock, c.requestedEndBlock+1)

	if err != nil {
		return fmt.Errorf("send sync request: %s", err)
	}

	return nil

}
