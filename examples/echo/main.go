package main

import (
	"github.com/ziku/go-dssim/examples/echo/messages"
	"github.com/ziku/go-dssim/pkg/logging"
	"github.com/ziku/go-dssim/pkg/network"
	"github.com/ziku/go-dssim/pkg/simulator"
)

func Cycle() {
	simulator.SendMessage(network.Message{
		Type:    messages.HelloType,
		Source:  simulator.Nodes[0],
		Dest:    simulator.Nodes[1],
		Payload: []byte("Hello Message"),
	})
	logging.Info("Node %d sent a Hello message to node %d", simulator.Nodes[0].ID, simulator.Nodes[1].ID)
}

func Cleanup() {
	logging.Info("Simulator cleanup")
}

func main() {
	network.Init()

	cfg := simulator.Config{
		Nodes:       100,
		Latency:     100,
		Cycles:      1000,
		CycleTime:   100,
		MessageLoss: 0.1,
	}
	simulator.Init(&cfg)

	network.RegisterMessageHandler(messages.HelloType, messages.Hello)
	network.RegisterMessageHandler(messages.HelloReplyType, messages.HelloReply)

	simulator.Run(Cycle, Cleanup)
}
