package gossip

import (
	"fmt"
)
func GetLocalAddress() InetAddress {
	return InetAddress(0)
}

type GossipDigest struct {
}

type gossiper struct {
	taskIntervalMs int64
	maxGenDiff int64
	liveEndpoints map[InetAddress]bool
	unreachableEndpoints map[InetAddress]bool
	seeds map[InetAddress]bool
	endpointStates map[InetAddress]Endpoint
	removedEndpoints map[InetAddress]bool
	isShadowRound bool
}

func (g *gossiper) run() {
	localEp := g.endpointStates[GetLocalAddress()]
	localEp.UpdateHeartbeat()
	gDigests := g.makeGossipDigest()
	if len(gDigests) > 0 {
		message := fmt.Sprintf("%#v", gDigests)
		isSeed := g.GossipToLiveMember(message)
		g.GossipToUnreachableMember(message)
		if !isSeed || len(g.liveEndpoints) < len(g.seeds) {
			g.GossipToSeed(message)
		}

		g.CheckStatus()
	}
}

func(g *gossiper) CheckStatus() {
}

func(g *gossiper) GossipToLiveMember(msg string) bool {
	return false
}

func(g *gossiper) GossipToUnreachableMember(msg string) {
}

func(g *gossiper) GossipToSeed(msg string) {
}

func (g *gossiper) makeGossipDigest() []GossipDigest {
	d := []GossipDigest{}
	return d
}
