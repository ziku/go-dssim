package messages

import (
	"github.com/ziku/go-dssim/pkg/logging"
	"github.com/ziku/go-dssim/pkg/network"
	"github.com/ziku/go-dssim/pkg/simulator"
)

func Hello(src *network.Peer, dest *network.Peer) {
	logging.Info("Node %d received a Hello message from node %d", dest.ID, src.ID)
	dest.MessagesCount["received"] += 1

	simulator.SendMessage(network.Message{
		Type:    HelloReplyType,
		Source:  dest,
		Dest:    src,
		Payload: []byte("HelloReply Message"),
	})
	logging.Info("Node %d sent a HelloReply message to node %d", dest.ID, src.ID)
}

func HelloReply(src *network.Peer, dest *network.Peer) {
	logging.Info("Node %d received a HelloReply message from node %d", dest.ID, src.ID)
	dest.MessagesCount["received"] += 1
}
