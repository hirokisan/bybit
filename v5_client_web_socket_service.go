package bybit

import (
	"github.com/gorilla/websocket"
)

// V5WebsocketServiceI :
type V5WebsocketServiceI interface {
	Public(CategoryV5) (V5WebsocketPublicService, error)
	Private() (V5WebsocketPrivateService, error)
	Trade() (V5WebsocketTradeService, error)
}

// V5WebsocketService :
type V5WebsocketService struct {
	client *WebSocketClient
}

// Public :
func (s *V5WebsocketService) Public(category CategoryV5) (V5WebsocketPublicServiceI, error) {
	url := s.client.baseURL + V5WebsocketPublicPathFor(category)
	var c *websocket.Conn
	var err error
	if s.client.dialer != nil {
		c, _, err = s.client.dialer.Dial(url, nil)
	} else {
		c, _, err = websocket.DefaultDialer.Dial(url, nil)
	}
	if err != nil {
		return nil, err
	}
	return &V5WebsocketPublicService{
		client:              s.client,
		connection:          c,
		category:            category,
		paramOrderBookMap:   make(map[V5WebsocketPublicOrderBookParamKey]func(V5WebsocketPublicOrderBookResponse) error),
		paramKlineMap:       make(map[V5WebsocketPublicKlineParamKey]func(V5WebsocketPublicKlineResponse) error),
		paramTickerMap:      make(map[V5WebsocketPublicTickerParamKey]func(V5WebsocketPublicTickerResponse) error),
		paramTradeMap:       make(map[V5WebsocketPublicTradeParamKey]func(V5WebsocketPublicTradeResponse) error),
		paramLiquidationMap: make(map[V5WebsocketPublicLiquidationParamKey]func(V5WebsocketPublicLiquidationResponse) error),
	}, nil
}

// Private :
func (s *V5WebsocketService) Private() (V5WebsocketPrivateServiceI, error) {
	url := s.client.baseURL + V5WebsocketPrivatePath
	var c *websocket.Conn
	var err error
	if s.client.dialer != nil {
		c, _, err = s.client.dialer.Dial(url, nil)
	} else {
		c, _, err = websocket.DefaultDialer.Dial(url, nil)
	}
	if err != nil {
		return nil, err
	}
	return &V5WebsocketPrivateService{
		client:            s.client,
		connection:        c,
		paramOrderMap:     make(map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivateOrderResponse) error),
		paramPositionMap:  make(map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivatePositionResponse) error),
		paramExecutionMap: make(map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivateExecutionResponse) error),
		paramWalletMap:    make(map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivateWalletResponse) error),
	}, nil
}

// Trade :
func (s *V5WebsocketService) Trade() (V5WebsocketTradeServiceI, error) {
	url := s.client.baseURL + V5WebsocketTradePath
	var c *websocket.Conn
	var err error
	if s.client.dialer != nil {
		c, _, err = s.client.dialer.Dial(url, nil)
	} else {
		c, _, err = websocket.DefaultDialer.Dial(url, nil)
	}
	if err != nil {
		return nil, err
	}
	return &V5WebsocketTradeService{
		client:     s.client,
		connection: c,
	}, nil
}

// V5 :
func (c *WebSocketClient) V5() *V5WebsocketService {
	return &V5WebsocketService{c}
}
