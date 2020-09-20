package jsonreader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	latencygenerator "github.com/rolorin/pd-apollo-loki/utility/latencyGenerator"
)

// Latency struct is struct for Latency specified in the json file of response

type jsonStruct struct {
	Latency  latencygenerator.Latency
	Response interface{} `json:"response"`
}

// JSONReader reads json from filename passed to it and responds with latency specified in file
func JSONReader(filename string) (latencygenerator.Latency, interface{}) {
	file, err := ioutil.ReadFile(filename)
	zeroLatency := latencygenerator.Latency{
		P95: float64(0),
		P99: float64(0),
	}

	JSON := jsonStruct{}

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error: Returning empty json and 0 latency")
		return zeroLatency, "{}"
	}

	err = json.Unmarshal(file, &JSON)

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error: Returning empty json and 0 latency")
		return zeroLatency, "{}"
	}
	return JSON.Latency, JSON.Response
}
