package gossip

import (
    "math"
)

const (
    SAMPLE_SIZE = 1000
    INITIAL_VALUE_NANOS = 100000000 //0.1s 
    MAX_INTERVAL_NANOS = 100000000 //0.1s 
)

var (
    PHI_FACTOR = 1.0/math.Log(10.0)
)

type InetAddress int64

type ArrayStats struct {
    intervals [SAMPLE_SIZE]int64
    sum int64
    size int64
    index int64
    isFilled bool
    mean float64
}

func (as *ArrayStats) Add(interval_ns int64) {
    if as.index == SAMPLE_SIZE {
        as.index = 0
        as.isFilled = true
    }
    //delete overrided item's value
    if as.isFilled {
        as.sum -= as.intervals[as.index]
    }

    as.intervals[as.index] = interval_ns
    as.sum += interval_ns
    as.index += 1
    as.mean = float64(as.sum/as.Size())
}

func (as *ArrayStats) Size() int64{
    if as.isFilled {
        return SAMPLE_SIZE
    } else {
        return as.index
    }
}

func (as *ArrayStats) GetArrivalIntervals() [SAMPLE_SIZE]int64{
    return as.intervals
}

type ArrivalWindow struct {
    arrivalIntervals *ArrayStats
    lastReportedPhi float64
    tLast int64
}

func (aw *ArrivalWindow) Add(value int64, ep InetAddress) {
    if aw.tLast == 0 {
        aw.tLast = INITIAL_VALUE_NANOS
    } else {
        interval_ns  := value - aw.tLast
        if interval_ns <= MAX_INTERVAL_NANOS {
            aw.arrivalIntervals.Add(interval_ns)
        }
    }
    aw.tLast = value
}

type FailureDetector struct {
}
