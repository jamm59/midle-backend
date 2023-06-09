package main

import (
	"sync"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

  var wg sync.WaitGroup

  var result string

  for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(result *string) {
      defer wg.Done()
      *result += "Hello, World! " 
    }( &result )

  }

  wg.Wait()

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
