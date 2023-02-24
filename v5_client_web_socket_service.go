package bybit

import "github.com/gorilla/websocket"

// V5WebsocketServiceI :
type V5WebsocketServiceI interface {
	Spot() V5WebsocketSpotServiceI
	Linear() V5WebsocketLinearServiceI
	Inverse() V5WebsocketInverseServiceI
	Option() V5WebsocketOptionServiceI
	Private() (V5WebsocketPrivateService, error)
}

// V5WebsocketService :
type V5WebsocketService struct {
	client *WebSocketClient
}

// Spot :
func (s *V5WebsocketService) Spot() V5WebsocketSpotServiceI {
	return &V5WebsocketSpotService{s.client}
}

// Linear :
func (s *V5WebsocketService) Linear() V5WebsocketLinearServiceI {
	return &V5WebsocketLinearService{s.client}
}

// Inverse :
func (s *V5WebsocketService) Inverse() V5WebsocketInverseServiceI {
	return &V5WebsocketInverseService{s.client}
}

// Option :
func (s *V5WebsocketService) Option() V5WebsocketOptionServiceI {
	return &V5WebsocketOptionService{s.client}
}

// Private :
func (s *V5WebsocketService) Private() (*V5WebsocketPrivateService, error) {
	url := s.client.baseURL + V5WebsocketPrivatePath
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	return &V5WebsocketPrivateService{
		client:          s.client,
		connection:      c,
		paramPrivateMap: map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivatePositionResponseContent) error{},
	}, nil
}

// V5 :
func (c *WebSocketClient) V5() *V5WebsocketService {
	return &V5WebsocketService{c}
}
