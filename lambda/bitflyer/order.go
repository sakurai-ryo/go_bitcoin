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

func (order *Order) PlaceOrder(apiKey, apiSecret string) (*OrderRes, error) {
	method := "POST"
	path := "/v1/me/sendchildorder"
	url := baseURL + path
	data, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	header, err := getHeader(method, path, apiKey, apiSecret, data)
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

func getHeader(method, path, apiKey, apiSecret string, body []byte) (map[string]string, error) {

	if method != "POST" {
		return nil, errors.New("POST is valid")
	}

	ts := strconv.FormatInt(time.Now().Unix(), 10)
	t := ts + method + path + string(body)
	mac := hmac.New(sha256.New, []byte(apiSecret))
	if _, err := mac.Write([]byte(t)); err != nil {
		return nil, err
	}
	sign := hex.EncodeToString(mac.Sum(nil))

	return map[string]string{
		"ACCESS-KEY":       apiKey,
		"ACCESS-TIMESTAMP": ts,
		"ACCESS-SIGN":      sign,
		"Content-Type":     "application/json",
	}, nil
}
