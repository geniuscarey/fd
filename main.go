package main

import (
    "fmt"
    "fd/gossip"
)

func main() {
    fmt.Println("hello, world\n")
    new_ep := gossip.NewEndpoint()
    fmt.Println(new_ep.IsAlive())
}
