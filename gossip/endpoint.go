package gossip

import (
    "time"
)

type Heartbeat struct {
    gen int64
    ver int64
}

func (h *Heartbeat) Update() {
    h.ver += 1
}

func (h *Heartbeat) GetVersion() int64 {
    return h.ver
}

func (h *Heartbeat) GetGeneration() int64 {
    return h.gen
}

type Endpoint struct {
    hb Heartbeat
    updateAt time.Time
    isAlive bool
}

func NewEndpoint() *Endpoint {
    return &Endpoint{Heartbeat{0,0}, time.Now(), false}
}

func (e *Endpoint) GetHeartbeat() Heartbeat {
    return e.hb
}

func (e *Endpoint) SetHeartbeat(nhb Heartbeat) {
    e.hb = nhb
}

func (e *Endpoint) UpdateTimestamp() {
    e.updateAt = time.Now()
}

func (e *Endpoint) GetTimestamp() time.Time {
    return e.updateAt
}

func (e *Endpoint) MarkAlive() {
    e.isAlive = true
}

func (e *Endpoint) MarkDead() {
    e.isAlive = false
}

func (e *Endpoint) IsAlive() bool {
    return e.isAlive
}
