package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	fmt.Println(request.RequestContext.Authorizer.JWT)

	return events.APIGatewayV2HTTPResponse{StatusCode: 200, Body: "it works"}, nil
}

func main() {
	lambda.Start(handler)
}
