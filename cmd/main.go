package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func entry(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("A user hit %s", request.Path)
	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "text/html; charset=utf-8"},
		Body:       fmt.Sprintf("<h1>Hello at %s </h1>", request.Path),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(entry)
}
