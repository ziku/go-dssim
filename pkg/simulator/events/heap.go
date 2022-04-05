package events

import "github.com/ziku/go-dssim/pkg/network"

type Event struct {
	Time    int
	Message network.Message
}

type EventsHeap []*Event

func (h EventsHeap) Len() int {
	return len(h)
}

func (h EventsHeap) Less(i, j int) bool {
	return h[i].Time < h[j].Time
}

func (h EventsHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *EventsHeap) Push(x interface{}) {
	*h = append(*h, x.(*Event))
}

func (h *EventsHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
