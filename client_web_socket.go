package bybit

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

const (
	// WebsocketBaseURL :
	WebsocketBaseURL = "wss://stream.bybit.com"
)

// WebSocketClient :
type WebSocketClient struct {
	baseURL string
	key     string
	secret  string
}

// NewWebsocketClient :
func NewWebsocketClient() *WebSocketClient {
	return &WebSocketClient{
		baseURL: WebsocketBaseURL,
	}
}

// WithAuth :
func (c *WebSocketClient) WithAuth(key string, secret string) *WebSocketClient {
	c.key = key
	c.secret = secret

	return c
}

func (c *WebSocketClient) buildAuthParam() ([]byte, error) {
	expires := time.Now().Unix()*1000 + 10000
	req := fmt.Sprintf("GET/realtime%d", expires)
	s := hmac.New(sha256.New, []byte(c.secret))
	if _, err := s.Write([]byte(req)); err != nil {
		return nil, err
	}
	signature := hex.EncodeToString(s.Sum(nil))
	param := struct {
		Op   string        `json:"op"`
		Args []interface{} `json:"args"`
	}{
		Op:   "auth",
		Args: []interface{}{c.key, expires, signature},
	}
	buf, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

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

// WebsocketExecutor :
type WebsocketExecutor interface {
	Run() error
	Close() error
	Ping() error
}

// Start :
func (c *WebSocketClient) Start(ctx context.Context, executors []WebsocketExecutor) {
	done := make(chan struct{})

	go func() {
		defer close(done)

		for {
			for _, executor := range executors {
				if err := executor.Run(); err != nil {
					if IsErrWebsocketClosed(err) {
						return
					}
					log.Println(err)
					return
				}
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
			for _, executor := range executors {
				if err := executor.Ping(); err != nil {
					return
				}
			}
		case <-ctx.Done():
			log.Println("interrupt")

			for _, executor := range executors {
				if err := executor.Close(); err != nil {
					return
				}
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
