package bybit

import (
	"net/url"

	"github.com/gorilla/websocket"
)

// SpotWebsocketV1Service :
type SpotWebsocketV1Service struct{}

// PublicV1 :
func (s *SpotWebsocketV1Service) PublicV1() (*SpotWebsocketV1PublicV1Service, error) {
	u := url.URL{Scheme: WebsocketScheme, Host: WebsocketHost, Path: SpotWebsocketV1PublicV1Path}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}
	return &SpotWebsocketV1PublicV1Service{
		connection:    c,
		paramTradeMap: map[SpotWebsocketV1PublicV1TradeParamKey]func(SpotWebsocketV1PublicV1TradeResponse) error{},
	}, nil
}
