package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	putresponder "github.com/rolorin/pd-apollo-loki/responders/put"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Request here", request.Body)
	response := putresponder.Put("./mocks/vendor-management/vendor.json")
	res, _ := json.Marshal(response)
	return events.APIGatewayProxyResponse{Body: string(res), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
