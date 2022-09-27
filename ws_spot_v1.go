package bybit

import (
	"net/url"

	"github.com/gorilla/websocket"
)

// SpotWebsocketV1Service :
type SpotWebsocketV1Service struct {
	client *WebSocketClient
}

// PublicV1 :
func (s *SpotWebsocketV1Service) PublicV1() (*SpotWebsocketV1PublicV1Service, error) {
	u := url.URL{Scheme: WebsocketScheme, Host: s.client.baseHost, Path: SpotWebsocketV1PublicV1Path}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}
	return &SpotWebsocketV1PublicV1Service{
		connection:    c,
		paramTradeMap: map[SpotWebsocketV1PublicV1TradeParamKey]func(SpotWebsocketV1PublicV1TradeResponse) error{},
	}, nil
}

// PublicV2 :
func (s *SpotWebsocketV1Service) PublicV2() (*SpotWebsocketV1PublicV2Service, error) {
	u := url.URL{Scheme: WebsocketScheme, Host: s.client.baseHost, Path: SpotWebsocketV1PublicV2Path}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}
	return &SpotWebsocketV1PublicV2Service{
		connection:    c,
		paramTradeMap: map[SpotWebsocketV1PublicV2TradeParamKey]func(SpotWebsocketV1PublicV2TradeResponse) error{},
	}, nil
}

// Private :
func (s *SpotWebsocketV1Service) Private() (*SpotWebsocketV1PrivateService, error) {
	u := url.URL{Scheme: WebsocketScheme, Host: s.client.baseHost, Path: SpotWebsocketV1PrivatePath}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}
	return &SpotWebsocketV1PrivateService{
		client:                      s.client,
		connection:                  c,
		paramOutboundAccountInfoMap: map[SpotWebsocketV1PrivateParamKey]func(SpotWebsocketV1PrivateOutboundAccountInfoResponse) error{},
	}, nil
}
