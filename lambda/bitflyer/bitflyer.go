package bitflyer

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"go_bitcoin/lambda/shared"
	"strconv"
	"time"
)

type Ticker struct {
	ProductCode     string  `json:"product_code"`
	State           string  `json:"state"`
	Timestamp       string  `json:"timestamp"`
	TickID          int     `json:"tick_id"`
	BestBid         float64 `json:"best_bid"`
	BestAsk         float64 `json:"best_ask"`
	BestBidSize     float64 `json:"best_bid_size"`
	BestAskSize     float64 `json:"best_ask_size"`
	TotalBidDepth   float64 `json:"total_bid_depth"`
	TotalAskDepth   float64 `json:"total_ask_depth"`
	MarketBidSize   float64 `json:"market_bid_size"`
	MarketAskSize   float64 `json:"market_ask_size"`
	Ltp             float64 `json:"ltp"`
	Volume          float64 `json:"volume"`
	VolumeByProduct float64 `json:"volume_by_product"`
}

const baseURL = "https://api.bitflyer.com"
const productCodeKey = "product_code"

func GetTicker(code ProductCode) (*Ticker, error) {
	url := baseURL + "/v1/ticker"
	res, err := shared.DoHttpRequest("GET", url, nil, map[string]string{productCodeKey: code.String()}, nil)
	if err != nil {
		return nil, err
	}

	var ticker Ticker
	err = json.Unmarshal(res, &ticker)
	if err != nil {
		return nil, err
	}
	return &ticker, nil
}

func getHeader(method, path, apiKey, apiSecret string, body []byte) (map[string]string, error) {
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
