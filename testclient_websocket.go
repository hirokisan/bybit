package bybit

import "os"

const (
	// TestWebsocketBaseURL :
	TestWebsocketBaseURL = "wss://stream-testnet.bybit.com"
)

// TestWebSocketClient :
type TestWebSocketClient struct {
	*WebSocketClient
}

// NewTestWebsocketClient :
func NewTestWebsocketClient() *TestWebSocketClient {
	return &TestWebSocketClient{
		WebSocketClient: &WebSocketClient{
			baseURL: TestWebsocketBaseURL,
		},
	}
}

// WithAuthFromEnv :
func (c *TestWebSocketClient) WithAuthFromEnv() *TestWebSocketClient {
	key, ok := os.LookupEnv("BYBIT_TEST_KEY")
	if !ok {
		panic("need BYBIT_TEST_KEY as environment variable")
	}
	secret, ok := os.LookupEnv("BYBIT_TEST_SECRET")
	if !ok {
		panic("need BYBIT_TEST_SECRET as environment variable")
	}
	c.key = key
	c.secret = secret

	return c
}
