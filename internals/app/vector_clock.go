package app

import (
	"sync"
	tt "vectorclock/transport"
)

type VectorClock struct {
	vector map[int]tt.Peers
	mutex  sync.Mutex
}

func NewVectorClock() *VectorClock {
	return &VectorClock{
		vector: make(map[int]tt.Peers),
	}
}

func (v *VectorClock) UpdateLocalClock(host string, port int) {

}

func (v *VectorClock) MergeVectorClocks(senderVectorClock map[int]tt.Peers) {

}
