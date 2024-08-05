package transport

type Peers struct {
	host       string
	port       int
	nodeID     int
	localClock int
}

func NewPeers() *Peers {
	return &Peers{
		// empty one
	}
}

func (p *Peers) AddPeer(host string, port, nodeID int) {

}

func (p *Peers) SearchPeer(host string, port int) bool {
	return false
}
