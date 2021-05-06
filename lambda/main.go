package main

import (
	"context"
	"fmt"
	"log"

	"go_bitcoin/lambda/bitflyer"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	t, err := bitflyer.GetTicker(bitflyer.Btcjpy)
	if err != nil {
		return "", nil
	}
	log.Print(t)

	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
