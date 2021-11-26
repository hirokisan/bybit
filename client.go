package bybit

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/url"
	"os"
	"sort"
	"strconv"
	"time"
)

const (
	// MainNetBaseURL :
	MainNetBaseURL = "https://api.bybit.com"
	// TestNetBaseURL :
	TestNetBaseURL = "https://api-testnet.bybit.com"
)

// Client :
type Client struct {
	BaseURL string
	Key     string
	Secret  string
}

// NewClient :
func NewClient() *Client {
	return &Client{
		BaseURL: MainNetBaseURL,
	}
}

// NewTestClient :
func NewTestClient() *Client {
	return &Client{
		BaseURL: TestNetBaseURL,
	}
}

// WithAuth :
func (c *Client) WithAuth(key string, secret string) *Client {
	c.Key = key
	c.Secret = secret

	return c
}

// WithAuthFromEnv :
func (c *Client) WithAuthFromEnv() *Client {
	key, ok := os.LookupEnv("BYBIT_TEST_KEY")
	if !ok {
		panic("need BYBIT_TEST_KEY as environment variable")
	}
	secret, ok := os.LookupEnv("BYBIT_TEST_SECRET")
	if !ok {
		panic("need BYBIT_TEST_SECRET as environment variable")
	}
	c.Key = key
	c.Secret = secret

	return c
}

// HasAuth : check has auth key and secret
func (c *Client) HasAuth() bool {
	return c.Key != "" && c.Secret != ""
}

// Wallet :
func (c *Client) Wallet() *WalletService {
	return &WalletService{c}
}

// Account :
func (c *Client) Account() *AccountService {
	return &AccountService{c}
}

// Market :
func (c *Client) Market() *MarketService {
	return &MarketService{c}
}

// BuildPublicURL :
func (c *Client) BuildPublicURL(path string, params map[string]string) (string, error) {
	if params == nil {
		params = map[string]string{}
	}
	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return "", nil
	}
	u.Path = path

	q := u.Query()
	for key, param := range params {
		q.Set(key, param)
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}

// BuildPrivateURL :
func (c *Client) BuildPrivateURL(path string, params map[string]string) (string, error) {
	if !c.HasAuth() {
		return "", fmt.Errorf("this is private endpoint, please set api key and secret")
	}

	if params == nil {
		params = map[string]string{}
	}
	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return "", err
	}
	u.Path = path

	intNow := int(time.Now().UTC().UnixNano() / int64(time.Millisecond))
	now := strconv.Itoa(intNow)

	params["api_key"] = c.Key
	params["timestamp"] = now

	params["sign"] = getSignature(params, c.Secret)

	q := u.Query()
	for key, param := range params {
		q.Set(key, param)
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}

func getSignature(params map[string]string, key string) string {
	keys := make([]string, len(params))
	i := 0
	_val := ""
	for k := range params {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		_val += k + "=" + params[k] + "&"
	}
	_val = _val[0 : len(_val)-1]
	h := hmac.New(sha256.New, []byte(key))
	io.WriteString(h, _val)
	return fmt.Sprintf("%x", h.Sum(nil))
}
