package gossip

import (
	"fmt"
	"math/rand"
)
func GetLocalAddress() InetAddress {
	return InetAddress(0)
}

type GossipDigest struct {
    Endpoint InetAddress `json:"endpoint"`
    Generation int64 `json:"generation"`
    Version int64 `json:"version"`
}

func (gd *GossipDigest) GetEndpoint() InetAddress {
    return gd.Endpoint
}

func (gd *GossipDigest) GetGeneration() int64{
    return gd.Generation
}

func (gd *GossipDigest) GetVersion() int64{
    return gd.Version
}

func (gd *GossipDigest) CompareTo(other *GossipDigest) int64{
    if gd.Generation == other.Generation {
        return gd.Version - other.Version
    } else  {
        return gd.Version - other.Version
    }
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
    //TODO, do nothing now
}

func(g *gossiper) SendGossip(msg string, toEndpoints map[InetAddress]bool) bool {
    nSend := len(toEndpoints)
    if nSend < 1 {
        return false
    }

    index := 0
    if nSend > 1 {
        index = rand.Intn(nSend)
    }

    var to InetAddress
    i := 0
    for k,_ := range(toEndpoints) {
        if i == index {
            to = k
            break
        }
        i += 1
    }

    if _, ok := g.seeds[to]; ok {
        return true
    } else {
        return false
    }
}

func(g *gossiper) GossipToLiveMember(msg string) bool {
    if len(g.liveEndpoints) == 0 {
	    return false
    }

    return g.SendGossip(msg, g.liveEndpoints)
}

func(g *gossiper) GossipToUnreachableMember(msg string) {
    nLive := len(g.liveEndpoints)
    nUnreachable := len(g.unreachableEndpoints)
    if nUnreachable > 0 {
        prob := float64(nUnreachable/(nLive+1))
        randomFloat := rand.Float64()
        if randomFloat < prob {
            g.SendGossip(msg, g.unreachableEndpoints)
        }
    }
}

func(g *gossiper) GossipToSeed(msg string) {
    nSeeds := len(g.seeds)
    if nSeeds > 0 {
        if len(g.liveEndpoints) == 0 {
            g.SendGossip(msg, g.seeds)
        }

        prob := float64(nSeeds/(len(g.liveEndpoints)+len(g.unreachableEndpoints)))
        randomFloat := rand.Float64()
        if randomFloat < prob{
            g.SendGossip(msg, g.seeds)
        }
    }
}

func (g *gossiper) makeGossipDigest() []GossipDigest {
    var (
        version int64 = 0
        generation int64 = 0
        i int64 = 0
    )

	gDigests := make([]GossipDigest, len(g.endpointStates))
    for epAddr,epState := range g.endpointStates {
        epHeartbeat := epState.GetHeartbeat()
        version = epHeartbeat.GetVersion()
        generation= epHeartbeat.GetGeneration()
        gDigests[i] = GossipDigest{epAddr, generation, version}
        i += 1
    }

	return gDigests
}

