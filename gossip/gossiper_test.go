package  gossip

import (
    "testing"
)

func Test_NewGossipDigest(t *testing.T) {
    var gDigests []GossipDigest
    gDigests = append(gDigests, GossipDigest{0x00000001, 0, 0})
    gDigests = append(gDigests, GossipDigest{0x00000002, 0, 0})
    gDigests = append(gDigests, GossipDigest{0x00000003, 0, 0})
    t.Logf("%s", gDigests)
}
