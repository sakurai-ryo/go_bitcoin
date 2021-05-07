package main

import (
	"context"
	"log"
	"math"

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
	buyPrice := roundDecimal(t.Ltp * 0.95)

	order := bitflyer.Order{
		ProductCode:    bitflyer.Btcjpy.String(),
		ChildOrderType: bitflyer.Limit.String(),
		Side:           bitflyer.Buy.String(),
		Price:          buyPrice,
		Size:           0.001,
		MinuteToExpire: 4320,
		TimeInForce:    bitflyer.Gtc.String(),
	}
	orderRes, err := order.PlaceOrder(secret.Key, secret.Secret)
	if err != nil {
		return "", err
	}
	log.Print("OrderResult: ", *orderRes)
	return "", nil
}

func roundDecimal(num float64) float64 {
	return math.Round(num)
}

func main() {
	lambda.Start(handler)
}
