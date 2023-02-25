package bybit

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

// V5WebsocketPrivateService :
type V5WebsocketPrivateService struct {
	client     *WebSocketClient
	connection *websocket.Conn

	paramPrivateMap map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivatePositionResponseContent) error
}

const (
	// V5WebsocketPrivatePath :
	V5WebsocketPrivatePath = "/v5/private"
)

// V5WebsocketPrivateTopic :
type V5WebsocketPrivateTopic string

const (
	// V5WebsocketPrivateTopicPosition :
	V5WebsocketPrivateTopicPosition = "position"
)

// V5WebsocketPrivateParamKey :
type V5WebsocketPrivateParamKey struct {
	Topic V5WebsocketPrivateTopic
}

type V5WebsocketPrivatePositionResponseContent struct {
	ID           string                                   `json:"id"`
	Topic        V5WebsocketPrivateTopic                  `json:"topic"`
	CreationTime int64                                    `json:"creationTime"`
	Data         []V5WebsocketPrivatePositionResponseData `json:"data"`
}

// V5WebsocketPrivatePositionResponseData :
type V5WebsocketPrivatePositionResponseData struct {
	AutoAddMargin   int        `json:"autoAddMargin"`
	PositionIdx     int        `json:"positionIdx"`
	TpSlMode        TpSlMode   `json:"tpSlMode"`
	TradeMode       int        `json:"tradeMode"`
	RiskID          int        `json:"riskId"`
	RiskLimitValue  string     `json:"riskLimitValue"`
	Symbol          SymbolV5   `json:"symbol"`
	Side            Side       `json:"side"`
	Size            string     `json:"size"`
	EntryPrice      string     `json:"entryPrice"`
	Leverage        string     `json:"leverage"`
	PositionValue   string     `json:"positionValue"`
	MarkPrice       string     `json:"markPrice"`
	PositionBalance string     `json:"positionBalance"`
	PositionIM      string     `json:"positionIM"`
	PositionMM      string     `json:"positionMM"`
	TakeProfit      string     `json:"takeProfit"`
	StopLoss        string     `json:"stopLoss"`
	TrailingStop    string     `json:"trailingStop"`
	UnrealisedPnl   string     `json:"unrealisedPnl"`
	CumRealisedPnl  string     `json:"cumRealisedPnl"`
	CreatedTime     string     `json:"CreatedTime"`
	UpdatedTime     string     `json:"updatedTime"`
	TpslMode        TpSlMode   `json:"tpslMode"`
	LiqPrice        string     `json:"liqPrice"`
	BustPrice       string     `json:"bustPrice"`
	Category        CategoryV5 `json:"category"`
	PositionStatus  string     `json:"positionStatus"`
}

// Key :
func (r *V5WebsocketPrivatePositionResponseContent) Key() V5WebsocketPrivateParamKey {
	return V5WebsocketPrivateParamKey{
		Topic: r.Topic,
	}
}

// addParamPositionFunc :
func (s *V5WebsocketPrivateService) addParamPositionFunc(param V5WebsocketPrivateParamKey, f func(V5WebsocketPrivatePositionResponseContent) error) error {
	if _, exist := s.paramPrivateMap[param]; exist {
		return errors.New("already registered for this param")
	}
	s.paramPrivateMap[param] = f
	return nil
}

// retrievePositionFunc :
func (s *V5WebsocketPrivateService) retrievePositionFunc(key V5WebsocketPrivateParamKey) (func(V5WebsocketPrivatePositionResponseContent) error, error) {
	f, exist := s.paramPrivateMap[key]
	if !exist {
		return nil, errors.New("func not found")
	}
	return f, nil
}

// judgeTopic :
func (s *V5WebsocketPrivateService) judgeTopic(respBody []byte) (V5WebsocketPrivateTopic, error) {
	parsedData := map[string]interface{}{}
	if err := json.Unmarshal(respBody, &parsedData); err == nil {
		if topic, ok := parsedData["topic"].(string); ok {
			return V5WebsocketPrivateTopic(topic), nil
		}
		if authStatus, ok := parsedData["success"].(bool); ok {
			if !authStatus {
				return "", errors.New("auth failed: " + parsedData["ret_msg"].(string))
			}
		}
		return "", nil
	} else {
		return "", err
	}
}

// parseResponse :
func (s *V5WebsocketPrivateService) parseResponse(respBody []byte, response interface{}) error {
	if err := json.Unmarshal(respBody, &response); err != nil {
		return err
	}
	return nil
}

// Subscribe :
func (s *V5WebsocketPrivateService) Subscribe() error {
	param, err := s.client.buildAuthParam()
	if err != nil {
		return err
	}
	if err := s.connection.WriteMessage(websocket.TextMessage, param); err != nil {
		return err
	}
	return nil
}

// RegisterFuncPosition :
func (s *V5WebsocketPrivateService) RegisterFuncPosition(f func(V5WebsocketPrivatePositionResponseContent) error) error {
	key := V5WebsocketPrivateParamKey{
		Topic: V5WebsocketPrivateTopicPosition,
	}
	if err := s.addParamPositionFunc(key, f); err != nil {
		return err
	}
	param := struct {
		Op   string        `json:"op"`
		Args []interface{} `json:"args"`
	}{
		Op:   "subscribe",
		Args: []interface{}{V5WebsocketPrivateTopicPosition},
	}
	buf, err := json.Marshal(param)
	if err != nil {
		return err
	}
	if err := s.connection.WriteMessage(websocket.TextMessage, buf); err != nil {
		return err
	}
	return nil
}

// ErrHandler :
type ErrHandler func(isWebsocketClosed bool, err error)

// Start :
func (s *V5WebsocketPrivateService) Start(ctx context.Context, errHandler ErrHandler) error {
	done := make(chan struct{})

	go func() {
		defer close(done)

		for {
			if err := s.Run(); err != nil {
				errHandler(IsErrWebsocketClosed(err), err)
				return
			}
		}
	}()

	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	for {
		select {
		case <-done:
			return nil
		case <-ticker.C:
			if err := s.Ping(); err != nil {
				return err
			}
		case <-ctx.Done():
			log.Println("interrupt")

			if err := s.Close(); err != nil {
				return err
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return nil
		}
	}
}

// Run :
func (s *V5WebsocketPrivateService) Run() error {
	_, message, err := s.connection.ReadMessage()
	if err != nil {
		return err
	}

	topic, err := s.judgeTopic(message)
	if err != nil {
		return err
	}
	switch topic {
	case V5WebsocketPrivateTopicPosition:
		var resp V5WebsocketPrivatePositionResponseContent
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}
		f, err := s.retrievePositionFunc(resp.Key())
		if err != nil {
			return err
		}
		if err := f(resp); err != nil {
			return err
		}
	}
	return nil
}

// Ping :
func (s *V5WebsocketPrivateService) Ping() error {
	if err := s.connection.WriteMessage(websocket.PingMessage, nil); err != nil {
		return err
	}
	return nil
}

// Close :
func (s *V5WebsocketPrivateService) Close() error {
	if err := s.connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
		return err
	}
	return nil
}
