package jsonreader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	latencycalculator "github.com/rolorin/pd-apollo-loki/utility/latencyCalculator"
)

// Latency struct is struct for Latency specified in the json file of response
type Latency struct {
	P95 float64 `json:"p95"`
	P99 float64 `json:"p99"`
}
type jsonStruct struct {
	Latency  Latency
	Response interface{} `json:"response"`
}

// JSONReader reads json from filename passed to it and responds with latency specified in file
func JSONReader(filename string) interface{} {
	file, err := ioutil.ReadFile(filename)

	JSON := jsonStruct{}

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error: Returning empty json and 0 latency")
		return "{}"
	}

	err = json.Unmarshal(file, &JSON)

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error: Returning empty json and 0 latency")
		return "{}"
	}
	latency := JSON.Latency
	fmt.Println(latency.P95, latency.P99)
	microseconds := (latencycalculator.Calculate(latency.P95, latency.P99) * int(1000))
	sleep := time.Duration(microseconds)
	time.Sleep(sleep)
	return JSON.Response
}
