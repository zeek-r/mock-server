package main

import (
	"encoding/json"

	"github.com/rolorin/pd-apollo-loki/config"
	"github.com/rolorin/pd-apollo-loki/utility/path"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rolorin/pd-apollo-loki/responders/patch"
)

// Handler for http...
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	responsePath := path.GetResponsePath(config.GetConfig("mockDirectory"), request.HTTPMethod, request.Path)
	response := patch.Response(responsePath)
	res, _ := json.Marshal(response)
	return events.APIGatewayProxyResponse{Body: string(res), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
