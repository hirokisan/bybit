package bybit

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
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

// WithBaseURL :
func (c *Client) WithBaseURL(url string) *Client {
	c.BaseURL = url

	return c
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

func (c *Client) populateSignature(src url.Values) url.Values {
	intNow := int(time.Now().UTC().UnixNano() / int64(time.Millisecond))
	now := strconv.Itoa(intNow)

	if src == nil {
		src = url.Values{}
	}

	src.Add("api_key", c.Key)
	src.Add("timestamp", now)
	src.Add("sign", getSignature(src, c.Secret))

	return src
}

func getSignature(src url.Values, key string) string {
	keys := make([]string, len(src))
	i := 0
	_val := ""
	for k := range src {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		_val += k + "=" + src.Get(k) + "&"
	}
	_val = _val[0 : len(_val)-1]
	h := hmac.New(sha256.New, []byte(key))
	_, err := io.WriteString(h, _val)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

func (c *Client) getPublicly(path string, query url.Values, dst interface{}) error {
	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return err
	}
	u.Path = path
	u.RawQuery = query.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&dst); err != nil {
		return err
	}

	return nil
}

func (c *Client) getPrivately(path string, query url.Values, dst interface{}) error {
	if !c.HasAuth() {
		return fmt.Errorf("this is private endpoint, please set api key and secret")
	}

	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return err
	}
	u.Path = path
	query = c.populateSignature(query)
	u.RawQuery = query.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&dst); err != nil {
		return err
	}

	return nil
}

func (c *Client) postJSON(path string, body []byte, dst interface{}) error {
	if !c.HasAuth() {
		return fmt.Errorf("this is private endpoint, please set api key and secret")
	}

	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return err
	}
	u.Path = path

	query := url.Values{}
	query = c.populateSignature(query)
	u.RawQuery = query.Encode()

	resp, err := http.Post(u.String(), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&dst); err != nil {
		return err
	}

	return nil
}

func (c *Client) postForm(path string, body url.Values, dst interface{}) error {
	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil
	}
	u.Path = path

	body = c.populateSignature(body)

	resp, err := http.Post(u.String(), "application/x-www-form-urlencoded", strings.NewReader(body.Encode()))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&dst); err != nil {
		return err
	}

	return nil
}

func (c *Client) deletePrivately(path string, query url.Values, dst interface{}) error {
	if !c.HasAuth() {
		return fmt.Errorf("this is private endpoint, please set api key and secret")
	}

	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return err
	}
	u.Path = path
	query = c.populateSignature(query)
	u.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodDelete, u.String(), nil)
	if err != nil {
		return err
	}
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&dst); err != nil {
		return err
	}

	return nil
}
