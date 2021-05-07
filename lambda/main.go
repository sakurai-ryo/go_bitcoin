package main

import (
	"context"
	"log"

	"go_bitcoin/lambda/bitflyer"
	"go_bitcoin/lambda/shared"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func handler(ctx context.Context, name MyEvent) (string, error) {

	secret, err := shared.GetSecret()
	if err != nil {
		return "", err
	}

	t, err := bitflyer.GetTicker(bitflyer.Btcjpy)
	if err != nil {
		return "", nil
	}
	log.Println("TickerResult: ", *t)

	price, size := bitflyer.GetByLogic(1)(10000, t)
	order := bitflyer.Order{
		ProductCode:    bitflyer.Btcjpy.String(),
		ChildOrderType: bitflyer.Limit.String(),
		Side:           bitflyer.Buy.String(),
		Price:          price,
		Size:           size,
		MinuteToExpire: 4320,
		TimeInForce:    bitflyer.Gtc.String(),
	}
	orderRes, err := order.PlaceOrder(secret)
	if err != nil {
		return "", err
	}
	log.Print("OrderResult: ", *orderRes)
	return "", nil
}

func main() {
	lambda.Start(handler)
}
