package post

import (
	jsonreader "github.com/rolorin/pd-apollo-loki/utility/jsonReader"
	latencygenerator "github.com/rolorin/pd-apollo-loki/utility/latencyGenerator"
)

// Response responder...
func Response(filePath string) interface{} {
	return latencygenerator.Sleep(jsonreader.JSONReader(filePath))
}
