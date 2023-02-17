package bybit

// SpotWebsocketService :
type SpotWebsocketService struct {
	client *WebSocketClient
}

// V1 :
func (s *SpotWebsocketService) V1() *SpotWebsocketV1Service {
	return &SpotWebsocketV1Service{s.client}
}

// Spot :
func (c *WebSocketClient) Spot() *SpotWebsocketService {
	return &SpotWebsocketService{c}
}
