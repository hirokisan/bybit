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

// SpotWebsocketV1PrivateService :
type SpotWebsocketV1PrivateService struct {
	client     *WebSocketClient
	connection *websocket.Conn

	paramOutboundAccountInfoMap map[SpotWebsocketV1PrivateParamKey]func(SpotWebsocketV1PrivateOutboundAccountInfoResponse) error
}

const (
	// SpotWebsocketV1PrivatePath :
	SpotWebsocketV1PrivatePath = "/spot/ws"
)

// SpotWebsocketV1PrivateEventType :
type SpotWebsocketV1PrivateEventType string

const (
	// SpotWebsocketV1PrivateEventTypeOutboundAccountInfo :
	SpotWebsocketV1PrivateEventTypeOutboundAccountInfo = "outboundAccountInfo"
)

// SpotWebsocketV1PrivateParamKey :
type SpotWebsocketV1PrivateParamKey struct {
	EventType SpotWebsocketV1PrivateEventType
}

// SpotWebsocketV1PrivateOutboundAccountInfoResponse :
type SpotWebsocketV1PrivateOutboundAccountInfoResponse struct {
	Content SpotWebsocketV1PrivateOutboundAccountInfoResponseContent
}

// SpotWebsocketV1PrivateOutboundAccountInfoResponseContent :
type SpotWebsocketV1PrivateOutboundAccountInfoResponseContent struct {
	EventType            SpotWebsocketV1PrivateEventType                                        `json:"e"`
	Timestamp            string                                                                 `json:"E"`
	AllowTrade           bool                                                                   `json:"T"`
	AllowWithdraw        bool                                                                   `json:"W"`
	AllowWDeposit        bool                                                                   `json:"D"`
	WalletBalanceChanges []SpotWebsocketV1PrivateOutboundAccountInfoResponseWalletBalanceChange `json:"B"`
}

// SpotWebsocketV1PrivateOutboundAccountInfoResponseWalletBalanceChange :
type SpotWebsocketV1PrivateOutboundAccountInfoResponseWalletBalanceChange struct {
	SymbolName       string `json:"a"`
	AvailableBalance string `json:"f"`
	ReservedBalance  string `json:"l"`
}

// UnmarshalJSON :
func (r *SpotWebsocketV1PrivateOutboundAccountInfoResponse) UnmarshalJSON(data []byte) error {
	parsedArrayData := []map[string]interface{}{}
	if err := json.Unmarshal(data, &parsedArrayData); err != nil {
		return err
	}
	if len(parsedArrayData) != 1 {
		return errors.New("unexpected response")
	}
	buf, err := json.Marshal(parsedArrayData[0])
	if err != nil {
		return err
	}
	if err := json.Unmarshal(buf, &r.Content); err != nil {
		return err
	}
	return nil
}

// MarshalJSON :
func (r *SpotWebsocketV1PrivateOutboundAccountInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Content)
}

// Key :
func (r *SpotWebsocketV1PrivateOutboundAccountInfoResponse) Key() SpotWebsocketV1PrivateParamKey {
	return SpotWebsocketV1PrivateParamKey{
		EventType: r.Content.EventType,
	}
}

// addParamOutboundAccountInfoFunc :
func (s *SpotWebsocketV1PrivateService) addParamOutboundAccountInfoFunc(param SpotWebsocketV1PrivateParamKey, f func(SpotWebsocketV1PrivateOutboundAccountInfoResponse) error) error {
	if _, exist := s.paramOutboundAccountInfoMap[param]; exist {
		return errors.New("already registered for this param")
	}
	s.paramOutboundAccountInfoMap[param] = f
	return nil
}

// retrieveOutboundAccountInfoFunc :
func (s *SpotWebsocketV1PrivateService) retrieveOutboundAccountInfoFunc(key SpotWebsocketV1PrivateParamKey) (func(SpotWebsocketV1PrivateOutboundAccountInfoResponse) error, error) {
	f, exist := s.paramOutboundAccountInfoMap[key]
	if !exist {
		return nil, errors.New("func not found")
	}
	return f, nil
}

type spotWebsocketV1PrivateEventJudge struct {
	EventType SpotWebsocketV1PrivateEventType
}

func (r *spotWebsocketV1PrivateEventJudge) UnmarshalJSON(data []byte) error {
	parsedData := map[string]interface{}{}
	if err := json.Unmarshal(data, &parsedData); err == nil {
		if event, ok := parsedData["e"].(string); ok {
			r.EventType = SpotWebsocketV1PrivateEventType(event)
		}
		if authStatus, ok := parsedData["auth"].(string); ok {
			if authStatus != "success" {
				return errors.New("auth failed")
			}
		}
		return nil
	}

	parsedArrayData := []map[string]interface{}{}
	if err := json.Unmarshal(data, &parsedArrayData); err != nil {
		return err
	}
	if len(parsedArrayData) != 1 {
		return errors.New("unexpected response")
	}
	r.EventType = SpotWebsocketV1PrivateEventType(parsedArrayData[0]["e"].(string))
	return nil
}

// judgeEventType :
func (s *SpotWebsocketV1PrivateService) judgeEventType(respBody []byte) (SpotWebsocketV1PrivateEventType, error) {
	var result spotWebsocketV1PrivateEventJudge
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", err
	}
	return result.EventType, nil
}

// parseResponse :
func (s *SpotWebsocketV1PrivateService) parseResponse(respBody []byte, response interface{}) error {
	if err := json.Unmarshal(respBody, &response); err != nil {
		return err
	}
	return nil
}

// Subscribe :
func (s *SpotWebsocketV1PrivateService) Subscribe() error {
	param, err := s.client.buildAuthParam()
	if err != nil {
		return err
	}
	if err := s.connection.WriteMessage(websocket.TextMessage, param); err != nil {
		return err
	}
	return nil
}

// RegisterFuncOutboundAccountInfo :
func (s *SpotWebsocketV1PrivateService) RegisterFuncOutboundAccountInfo(f func(SpotWebsocketV1PrivateOutboundAccountInfoResponse) error) error {
	key := SpotWebsocketV1PrivateParamKey{
		EventType: SpotWebsocketV1PrivateEventTypeOutboundAccountInfo,
	}
	if err := s.addParamOutboundAccountInfoFunc(key, f); err != nil {
		return err
	}
	return nil
}

// Start :
func (s *SpotWebsocketV1PrivateService) Start(ctx context.Context) {
	done := make(chan struct{})

	go func() {
		defer close(done)

		for {
			if err := s.Run(); err != nil {
				if IsErrWebsocketClosed(err) {
					return
				}
				log.Println(err)
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
			return
		case <-ticker.C:
			if err := s.Ping(); err != nil {
				return
			}
		case <-ctx.Done():
			log.Println("interrupt")

			if err := s.Close(); err != nil {
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

// Run :
func (s *SpotWebsocketV1PrivateService) Run() error {
	_, message, err := s.connection.ReadMessage()
	if err != nil {
		return err
	}

	topic, err := s.judgeEventType(message)
	if err != nil {
		return err
	}
	switch topic {
	case SpotWebsocketV1PrivateEventTypeOutboundAccountInfo:
		var resp SpotWebsocketV1PrivateOutboundAccountInfoResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}
		f, err := s.retrieveOutboundAccountInfoFunc(resp.Key())
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
func (s *SpotWebsocketV1PrivateService) Ping() error {
	if err := s.connection.WriteMessage(websocket.PingMessage, nil); err != nil {
		return err
	}
	return nil
}

// Close :
func (s *SpotWebsocketV1PrivateService) Close() error {
	if err := s.connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
		return err
	}
	return nil
}
