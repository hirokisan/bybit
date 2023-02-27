package bybit

import (
	"github.com/gorilla/websocket"
)

// V5WebsocketServiceI :
type V5WebsocketServiceI interface {
	Public(CategoryV5) (V5WebsocketPublicService, error)
	Private() (V5WebsocketPrivateService, error)
}

// V5WebsocketService :
type V5WebsocketService struct {
	client *WebSocketClient
}

// Public :
func (s *V5WebsocketService) Public(category CategoryV5) (V5WebsocketPublicServiceI, error) {
	url := s.client.baseURL + V5WebsocketPublicPathFor(category)
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	return &V5WebsocketPublicService{
		client:            s.client,
		connection:        c,
		paramOrderBookMap: map[V5WebsocketPublicOrderBookParamKey]func(V5WebsocketPublicOrderBookResponse) error{},
	}, nil
}

// Private :
func (s *V5WebsocketService) Private() (V5WebsocketPrivateServiceI, error) {
	url := s.client.baseURL + V5WebsocketPrivatePath
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	return &V5WebsocketPrivateService{
		client:           s.client,
		connection:       c,
		paramOrderMap:    map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivateOrderResponse) error{},
		paramPositionMap: map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivatePositionResponse) error{},
		paramWalletMap:   map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivateWalletResponse) error{},
	}, nil
}

// V5 :
func (c *WebSocketClient) V5() *V5WebsocketService {
	return &V5WebsocketService{c}
}
