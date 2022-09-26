package bybit

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"
)

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
