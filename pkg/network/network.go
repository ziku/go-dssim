package network

import (
	"github.com/ziku/go-dssim/pkg/logging"
)

var (
	messageHandlers = map[int]func(src *Peer, dest *Peer){}
)

type Message struct {
	Type    int
	Source  *Peer
	Dest    *Peer
	Payload []byte
}

type Peer struct {
	ID            int
	MessagesCount map[string]int
}

func (p *Peer) Init(id int) {
	p.ID = id
	p.MessagesCount = map[string]int{
		"total":    0,
		"sent":     0,
		"received": 0,
		"dropped":  0,
	}
}

func RegisterMessageHandler(messageType int, handler func(src *Peer, dest *Peer)) {
	messageHandlers[messageType] = handler
}

func HandleMessage(messageType int, src *Peer, dest *Peer, payload []byte) {
	if handler, ok := messageHandlers[messageType]; ok {
		handler(src, dest)
	} else {
		logging.Error("No handler for message type %d", messageType)
	}
}

func Init() {
	logging.Info("Network initialized")
}
