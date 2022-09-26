package bybit

const (
	// WebsocketScheme :
	WebsocketScheme = "wss"
	// WebsocketHost :
	WebsocketHost = "stream-testnet.bybit.com"
)

// WebSocketClient :
type WebSocketClient struct{}

// NewWebsocketClient :
func NewWebsocketClient() *WebSocketClient {
	return &WebSocketClient{}
}

// SpotWebsocketService :
type SpotWebsocketService struct {
}

// V1 :
func (s *SpotWebsocketService) V1() *SpotWebsocketV1Service {
	return &SpotWebsocketV1Service{}
}

// Spot :
func (c *WebSocketClient) Spot() *SpotWebsocketService {
	return &SpotWebsocketService{}
}
