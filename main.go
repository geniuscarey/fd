package main

import (
    "fmt"
    "github.com/geniuscarey/fd/gossip"
)

func main() {
    var gDigests []GossipDigest
    gDigests = append(gDigests, GossipDigest{0x00000001, 0, 0})
    gDigests = append(gDigests, GossipDigest{0x00000002, 0, 0})
    gDigests = append(gDigests, GossipDigest{0x00000003, 0, 0})
    fmt.Println(gDigests)
}
