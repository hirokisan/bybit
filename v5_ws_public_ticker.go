package bybit

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/gorilla/websocket"
)

// SubscribeTicker :
func (s *V5WebsocketPublicService) SubscribeTicker(
	key V5WebsocketPublicTickerParamKey,
	f func(V5WebsocketPublicTickerResponse) error,
) (func() error, error) {
	if err := s.addParamTickerFunc(key, f); err != nil {
		return nil, err
	}
	param := struct {
		Op   string        `json:"op"`
		Args []interface{} `json:"args"`
	}{
		Op:   "subscribe",
		Args: []interface{}{key.Topic()},
	}
	buf, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	if err := s.writeMessage(websocket.TextMessage, buf); err != nil {
		return nil, err
	}
	return func() error {
		param := struct {
			Op   string        `json:"op"`
			Args []interface{} `json:"args"`
		}{
			Op:   "unsubscribe",
			Args: []interface{}{key.Topic()},
		}
		buf, err := json.Marshal(param)
		if err != nil {
			return err
		}
		if err := s.writeMessage(websocket.TextMessage, []byte(buf)); err != nil {
			return err
		}
		s.removeParamTickerFunc(key)
		return nil
	}, nil
}

// V5WebsocketPublicTickerParamKey :
type V5WebsocketPublicTickerParamKey struct {
	Symbol SymbolV5
}

// Topic :
func (k *V5WebsocketPublicTickerParamKey) Topic() string {
	return fmt.Sprintf("%s.%s", V5WebsocketPublicTopicTicker, k.Symbol)
}

// V5WebsocketPublicTickerResponse :
type V5WebsocketPublicTickerResponse struct {
	Topic     string                      `json:"topic"`
	Type      string                      `json:"type"`
	TimeStamp int64                       `json:"ts"`
	CrossSeq  int64                       `json:"cs,omitempty"`
	ID        string                      `json:"id,omitempty"`
	Data      V5WebsocketPublicTickerData `json:"data"`
}

// V5WebsocketPublicTickerData :
type V5WebsocketPublicTickerData struct {
	LinearInverse *V5WebsocketPublicTickerLinearInverseResult
	Option        *V5WebsocketPublicTickerOptionResult
	Spot          *V5WebsocketPublicTickerSpotResult
}

// V5WebsocketPublicTickerLinearInverseResult :
type V5WebsocketPublicTickerLinearInverseResult struct {
	Symbol                 SymbolV5      `json:"symbol"`
	TickDirection          TickDirection `json:"tickDirection"`
	Price24hPercent        string        `json:"price24hPcnt"`
	LastPrice              string        `json:"lastPrice"`
	PrevPrice24h           string        `json:"prevPrice24h"`
	HighPrice24h           string        `json:"highPrice24h"`
	LowPrice24h            string        `json:"lowPrice24h"`
	PrevPrice1h            string        `json:"prevPrice1h"`
	MarkPrice              string        `json:"markPrice"`
	IndexPrice             string        `json:"indexPrice"`
	OpenInterest           string        `json:"openInterest"`
	OpenInterestValue      string        `json:"openInterestValue"`
	Turnover24h            string        `json:"turnover24h"`
	Volume24h              string        `json:"volume24h"`
	NextFundingTime        string        `json:"nextFundingTime"`
	FundingRate            string        `json:"fundingRate"`
	Bid1Price              string        `json:"bid1Price"`
	Bid1Size               string        `json:"bid1Size"`
	Ask1Price              string        `json:"ask1Price"`
	Ask1Size               string        `json:"ask1Size"`
	DeliveryTime           string        `json:"deliveryTime,omitempty"`
	BasisRate              string        `json:"basisRate,omitempty"`
	DeliveryFeeRate        string        `json:"deliveryFeeRate,omitempty"`
	PredictedDeliveryPrice string        `json:"predictedDeliveryPrice,omitempty"`
}

// V5WebsocketPublicTickerOptionResult :
type V5WebsocketPublicTickerOptionResult struct {
	Symbol                 SymbolV5 `json:"symbol"`
	BidPrice               string   `json:"bidPrice"`
	BidSize                string   `json:"bidSize"`
	BidIv                  string   `json:"bidIv"`
	AskPrice               string   `json:"askPrice"`
	AskSize                string   `json:"askSize"`
	AskIv                  string   `json:"askIv"`
	LastPrice              string   `json:"lastPrice"`
	HighPrice24H           string   `json:"highPrice24h"`
	LowPrice24H            string   `json:"lowPrice24h"`
	MarkPrice              string   `json:"markPrice"`
	IndexPrice             string   `json:"indexPrice"`
	MarkPriceIv            string   `json:"markPriceIv"`
	UnderlyingPrice        string   `json:"underlyingPrice"`
	OpenInterest           string   `json:"openInterest"`
	Turnover24H            string   `json:"turnover24h"`
	Volume24H              string   `json:"volume24h"`
	TotalVolume            string   `json:"totalVolume"`
	TotalTurnover          string   `json:"totalTurnover"`
	Delta                  string   `json:"delta"`
	Gamma                  string   `json:"gamma"`
	Vega                   string   `json:"vega"`
	Theta                  string   `json:"theta"`
	PredictedDeliveryPrice string   `json:"predictedDeliveryPrice"`
	Change24H              string   `json:"change24h"`
}

// V5WebsocketPublicTickerSpotResult :
type V5WebsocketPublicTickerSpotResult struct {
	Symbol        SymbolV5 `json:"symbol"`
	LastPrice     string   `json:"lastPrice"`
	HighPrice24H  string   `json:"highPrice24h"`
	LowPrice24H   string   `json:"lowPrice24h"`
	PrevPrice24H  string   `json:"prevPrice24h"`
	Volume24H     string   `json:"volume24h"`
	Turnover24H   string   `json:"turnover24h"`
	Price24HPcnt  string   `json:"price24hPcnt"`
	UsdIndexPrice string   `json:"usdIndexPrice"`
}

// Key :
func (r *V5WebsocketPublicTickerResponse) Key() V5WebsocketPublicTickerParamKey {
	topic := r.Topic
	arr := strings.Split(topic, ".")
	if arr[0] != V5WebsocketPublicTopicTicker.String() || len(arr) != 2 {
		return V5WebsocketPublicTickerParamKey{}
	}

	return V5WebsocketPublicTickerParamKey{
		Symbol: SymbolV5(arr[1]),
	}
}

// addParamTickerFunc :
func (s *V5WebsocketPublicService) addParamTickerFunc(key V5WebsocketPublicTickerParamKey, f func(V5WebsocketPublicTickerResponse) error) error {
	if _, exist := s.paramTickerMap[key]; exist {
		return errors.New("already registered for this key")
	}
	s.paramTickerMap[key] = f
	return nil
}

// removeParamTickerFunc :
func (s *V5WebsocketPublicService) removeParamTickerFunc(key V5WebsocketPublicTickerParamKey) {
	delete(s.paramTickerMap, key)
}

// retrieveTickerFunc :
func (s *V5WebsocketPublicService) retrieveTickerFunc(key V5WebsocketPublicTickerParamKey) (func(V5WebsocketPublicTickerResponse) error, error) {
	f, exist := s.paramTickerMap[key]
	if !exist {
		return nil, errors.New("func not found")
	}
	return f, nil
}
