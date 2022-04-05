package simulator

import (
	"container/heap"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ziku/go-dssim/pkg/logging"
	"github.com/ziku/go-dssim/pkg/network"
	"github.com/ziku/go-dssim/pkg/simulator/events"
)

type Config struct {
	Nodes       int
	Latency     int
	Cycles      int
	CycleTime   int
	MessageLoss float64
}

var config *Config

var Nodes []*network.Peer
var Heap events.EventsHeap
var CurrentTime int

func SendMessage(message network.Message) {
	if rand.Float64() > config.MessageLoss {
		message.Source.MessagesCount["sent"] += 1
		lat := time.Duration(rand.Intn(config.Latency))
		Heap.Push(&events.Event{
			Time:    CurrentTime + int(lat),
			Message: message,
		})
		//network.HandleMessage(message.Type, message.Source, message.Dest, message.Payload)
	} else {
		message.Source.MessagesCount["dropped"] += 1
	}
}

func HandleHeap() {
	for {
		var event *events.Event = heap.Pop(&Heap).(*events.Event)
		if event.Time > CurrentTime {
			logging.Info("Times: %d %d", event.Time, CurrentTime)
			heap.Push(&Heap, event)
			break
		}
		network.HandleMessage(event.Message.Type, event.Message.Source, event.Message.Dest, event.Message.Payload)
	}
}

func Init(cfg *Config) {
	config = cfg
	logging.Info("Simulator initialized")
}

func Run(cycle func(), cleanup func()) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(0)
	}()
	defer func() {
		if r := recover(); r != nil {
			logging.Error("Simulator crashed: ", r)
		}
	}()

	for i := 0; i < config.Nodes; i++ {
		peer := network.Peer{}
		peer.Init(i)
		Nodes = append(Nodes, &peer)
		logging.Info("Node %d initialized", peer.ID)
	}

	for i := 0; i < config.Cycles; i++ {
		logging.Info("Cycle %d started", i)
		logging.Info("Current time: %d", CurrentTime)
		logging.Info("Heap size: %d", Heap.Len())

		cycle()
		HandleHeap()
		//time.Sleep(1 * time.Second)

		CurrentTime += config.CycleTime
	}
}
