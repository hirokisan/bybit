package bybit

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/url"
	"sort"
	"strconv"
	"time"
)

const (
	// MainNetBaseURL :
	MainNetBaseURL = "https://api.bybit.com"
)

// Client :
type Client struct {
	BaseURL string
	Key     string
	Secret  string
}

// NewClient :
func NewClient(key string, secret string) *Client {
	return &Client{
		BaseURL: MainNetBaseURL,
		Key:     key,
		Secret:  secret,
	}
}

// Wallet :
func (c *Client) Wallet() *WalletService {
	return &WalletService{c}
}

// BuildURL :
func (c *Client) BuildURL(path string, params map[string]string) string {
	u, err := url.Parse(c.BaseURL)
	if err != nil {
		panic(err)
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

	return u.String()
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
