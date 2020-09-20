package latencygenerator

import (
	"time"

	latencycalculator "github.com/rolorin/pd-apollo-loki/utility/latencyCalculator"
)

// Latency struct...
type Latency struct {
	P95 float64 `json:"p95"`
	P99 float64 `json:"p99"`
}

// LatencyGenerator generates fake latency...
func Sleep(latency Latency, response interface{}) interface{} {
	microseconds := (latencycalculator.Calculate(latency.P95, latency.P99) * int(1000))
	sleep := time.Duration(microseconds)
	time.Sleep(sleep)
	return response
}
