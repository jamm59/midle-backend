package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
  return &events.APIGatewayProxyResponse{
    StatusCode:        200,
    Body:              "Hello, World!",
    Headers: map[string]string{
      "Access-Control-Allow-Origin":      "*",
      "Access-Control-Allow-Credentials": "true",
      "Content-Type":                     "application/json",
    },
  }, nil
}

func main() {
  lambda.Start(handler)
}
