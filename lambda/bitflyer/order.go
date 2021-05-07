package bitflyer

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"go_bitcoin/lambda/shared"
	"strconv"
	"time"
)

const btcMinAmount = 0.001
const btcPlace = 4.0

type Order struct {
	ProductCode    string  `json:"product_code"`
	ChildOrderType string  `json:"child_order_type"`
	Side           string  `json:"side"`
	Price          float64 `json:"price"`
	Size           float64 `json:"size"`
	MinuteToExpire int     `json:"minute_to_expire"`
	TimeInForce    string  `json:"time_in_force"`
}

type OrderRes struct {
	ChildOrderAcceptanceId string `json:"child_order_acceptance_id"`
}

func GetByLogic(strategy int) func(float64, *Ticker) (float64, float64) {
	var logic func(budget float64, t *Ticker) (float64, float64)
	switch strategy {
	case 1:
		// LTPの98.5%での価格
		logic = func(budget float64, t *Ticker) (float64, float64) {
			var buyPrice, buySize float64
			buyPrice = shared.RoundDecimal(t.Ltp * 0.985)
			buySize = shared.CalcAmount(buyPrice, budget, btcMinAmount, btcPlace)
			return buyPrice, buySize
		}
		break
	default:
		// BestAskを注文価格とする
		logic = func(budget float64, t *Ticker) (float64, float64) {
			var buyPrice, buySize float64
			buyPrice = shared.RoundDecimal(t.BestAsk)
			buySize = shared.CalcAmount(buyPrice, budget, btcMinAmount, btcPlace)
			return buyPrice, buySize
		}
		break
	}
	return logic
}

func (order *Order) PlaceWithParams(secret *shared.Secret, price, size float64) (*OrderRes, error) {
	order.Price = price
	order.Size = size

	orderRes, err := order.PlaceOrder(secret)
	if err != nil {
		return nil, err
	}
	return orderRes, nil
}

func (order *Order) PlaceOrder(secret *shared.Secret) (*OrderRes, error) {
	method := "POST"
	path := "/v1/me/sendchildorder"
	url := baseURL + path
	data, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	header, err := getHeader(method, path, secret, data)
	if err != nil {
		return nil, err
	}

	res, err := shared.DoHttpRequest(method, url, header, nil, data)
	if err != nil {
		return nil, err
	}

	var orderRes OrderRes
	if err := json.Unmarshal(res, &orderRes); err != nil {
		return nil, err
	}
	if len(orderRes.ChildOrderAcceptanceId) == 0 {
		return nil, errors.New(string(res))
	}
	return &orderRes, nil
}

func getHeader(method, path string, secret *shared.Secret, body []byte) (map[string]string, error) {

	if method != "POST" {
		return nil, errors.New("POST is valid")
	}

	ts := strconv.FormatInt(time.Now().Unix(), 10)
	t := ts + method + path + string(body)
	mac := hmac.New(sha256.New, []byte(secret.Secret))
	if _, err := mac.Write([]byte(t)); err != nil {
		return nil, err
	}
	sign := hex.EncodeToString(mac.Sum(nil))

	return map[string]string{
		"ACCESS-KEY":       secret.Key,
		"ACCESS-TIMESTAMP": ts,
		"ACCESS-SIGN":      sign,
		"Content-Type":     "application/json",
	}, nil
}
