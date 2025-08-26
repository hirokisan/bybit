package bybit

import (
	"context"
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// WebsocketBaseURL :
	WebsocketBaseURL = "wss://stream.bybit.com"
	// WebsocketBaseURL2 :
	WebsocketBaseURL2 = "wss://stream.bytick.com"
)

// WebSocketClient :
type WebSocketClient struct {
	debug  bool
	logger *log.Logger

	baseURL string
	key     string
	secret  string
	
	// RSA signing support
	useRSA     bool
	privateKey *rsa.PrivateKey
	
	dialer  *websocket.Dialer
}

func (c *WebSocketClient) debugf(format string, v ...interface{}) {
	if c.debug {
		c.logger.Printf(format, v...)
	}
}

// NewWebsocketClient :
func NewWebsocketClient() *WebSocketClient {
	return &WebSocketClient{
		logger: newDefaultLogger(),

		baseURL: WebsocketBaseURL,
	}
}

// WithDebug :
func (c *WebSocketClient) WithDebug(debug bool) *WebSocketClient {
	c.debug = debug

	return c
}

// WithLogger :
func (c *WebSocketClient) WithLogger(logger *log.Logger) *WebSocketClient {
	c.debug = true
	c.logger = logger

	return c
}

// WithAuth :
func (c *WebSocketClient) WithAuth(key string, secret string) *WebSocketClient {
	c.key = key
	c.secret = secret
	c.useRSA = false

	return c
}

// WithAuthRSA sets up authentication using RSA private key
func (c *WebSocketClient) WithAuthRSA(key string, privateKeyPEM string) *WebSocketClient {
	c.key = key
	c.useRSA = true
	
	// Parse the private key
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil {
		panic("failed to parse PEM block containing the private key")
	}
	
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		// Try PKCS8 format if PKCS1 fails
		parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			panic(fmt.Sprintf("failed to parse private key: %v", err))
		}
		var ok bool
		privateKey, ok = parsedKey.(*rsa.PrivateKey)
		if !ok {
			panic("not an RSA private key")
		}
	}
	
	c.privateKey = privateKey
	
	return c
}

// WithBaseURL :
func (c *WebSocketClient) WithBaseURL(url string) *WebSocketClient {
	c.baseURL = url

	return c
}

// WithDialer :
func (c *WebSocketClient) WithDialer(dialer *websocket.Dialer) *WebSocketClient {
	c.dialer = dialer

	return c
}

// hasAuth : check has auth key and secret
func (c *WebSocketClient) hasAuth() bool {
	return c.key != "" && c.secret != ""
}

func (c *WebSocketClient) buildAuthParam() ([]byte, error) {
	if !c.hasAuth() {
		return nil, fmt.Errorf("this is private endpoint, please set api key and secret")
	}

	expires := time.Now().Unix()*1000 + 10000
	req := fmt.Sprintf("GET/realtime%d", expires)
	
	var signature string
	if c.useRSA {
		// For RSA signatures
		hashed := sha256.Sum256([]byte(req))
		sig, err := rsa.SignPKCS1v15(rand.Reader, c.privateKey, crypto.SHA256, hashed[:])
		if err != nil {
			return nil, fmt.Errorf("failed to sign with RSA: %w", err)
		}
		signature = base64.StdEncoding.EncodeToString(sig)
	} else {
		// For HMAC signatures
		s := hmac.New(sha256.New, []byte(c.secret))
		if _, err := s.Write([]byte(req)); err != nil {
			return nil, err
		}
		signature = hex.EncodeToString(s.Sum(nil))
	}
	
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
					c.debugf("websocket executor error: %s", err)
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
			c.debugf("caught websocket interrupt signal")

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
