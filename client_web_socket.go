package bybit

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

const (
	// WebsocketScheme :
	WebsocketScheme = "wss"
	// WebsocketHost :
	//WebsocketHost = "stream-testnet.bybit.com"
	WebsocketHost = "stream-testnet.bybit.com"
)

// WebSocketClient :
type WebSocketClient struct {
	key    string
	secret string
}

// NewWebsocketClient :
func NewWebsocketClient() *WebSocketClient {
	return &WebSocketClient{}
}

// WithAuth :
func (c *WebSocketClient) WithAuth(key string, secret string) *WebSocketClient {
	c.key = key
	c.secret = secret

	return c
}

type authParam struct {
	Op   string        `json:"op"`
	Args []interface{} `json:"args"`
}

func (c *WebSocketClient) buildAuthParam() authParam {
	expires := time.Now().Unix()*1000 + 10000
	req := fmt.Sprintf("GET/realtime%d", expires)
	sig := hmac.New(sha256.New, []byte(c.secret))
	sig.Write([]byte(req))
	signature := hex.EncodeToString(sig.Sum(nil))
	param := authParam{
		Op:   "auth",
		Args: []interface{}{c.key, expires, signature},
	}
	return param
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
