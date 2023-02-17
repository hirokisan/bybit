package bybit

// V5WebsocketServiceI :
type V5WebsocketServiceI interface {
	Spot() V5WebsocketSpotServiceI
	Linear() V5WebsocketLinearServiceI
	Inverse() V5WebsocketInverseServiceI
	Option() V5WebsocketOptionServiceI
	Private() V5WebsocketPrivateServiceI
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
func (s *V5WebsocketService) Private() V5WebsocketPrivateServiceI {
	return &V5WebsocketPrivateService{s.client}
}

// V5 :
func (c *WebSocketClient) V5() V5WebsocketServiceI {
	return &V5WebsocketService{c}
}
