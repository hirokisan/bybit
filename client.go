package bybit

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	// MainNetBaseURL :
	MainNetBaseURL = "https://api.bybit.com"
)

// Client :
type Client struct {
	httpClient *http.Client

	baseURL string
	key     string
	secret  string
}

// NewClient :
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},

		baseURL: MainNetBaseURL,
	}
}

// WithHTTPClient :
func (c *Client) WithHTTPClient(httpClient *http.Client) *Client {
	c.httpClient = httpClient

	return c
}

// WithAuth :
func (c *Client) WithAuth(key string, secret string) *Client {
	c.key = key
	c.secret = secret

	return c
}

// Request :
func (c *Client) Request(req *http.Request, dst interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch {
	case 200 <= resp.StatusCode && resp.StatusCode <= 299:
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(body, &dst); err != nil {
			return err
		}

		{
			rateLimitError := &RateLimitError{}
			if err := json.Unmarshal(body, rateLimitError); err != nil {
				return err
			}
			if rateLimitError.RetCode == 10006 {
				return rateLimitError
			}
		}

		return nil
	case resp.StatusCode == http.StatusForbidden:
		return ErrAccessDenied
	case resp.StatusCode == http.StatusNotFound:
		return ErrPathNotFound
	default:
		return errors.New("unexpected error")
	}
}

// hasAuth : check has auth key and secret
func (c *Client) hasAuth() bool {
	return c.key != "" && c.secret != ""
}

func (c *Client) populateSignature(src url.Values) url.Values {
	intNow := int(time.Now().UTC().UnixNano() / int64(time.Millisecond))
	now := strconv.Itoa(intNow)

	if src == nil {
		src = url.Values{}
	}

	src.Add("api_key", c.key)
	src.Add("timestamp", now)
	src.Add("sign", getSignature(src, c.secret))

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
	u, err := url.Parse(c.baseURL)
	if err != nil {
		return err
	}
	u.Path = path
	u.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}

	if err := c.Request(req, &dst); err != nil {
		return err
	}

	return nil
}

func (c *Client) getPrivately(path string, query url.Values, dst interface{}) error {
	if !c.hasAuth() {
		return fmt.Errorf("this is private endpoint, please set api key and secret")
	}

	u, err := url.Parse(c.baseURL)
	if err != nil {
		return err
	}
	u.Path = path
	query = c.populateSignature(query)
	u.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}

	if err := c.Request(req, &dst); err != nil {
		return err
	}

	return nil
}

func (c *Client) postJSON(path string, body []byte, dst interface{}) error {
	if !c.hasAuth() {
		return fmt.Errorf("this is private endpoint, please set api key and secret")
	}

	u, err := url.Parse(c.baseURL)
	if err != nil {
		return err
	}
	u.Path = path

	query := url.Values{}
	query = c.populateSignature(query)
	u.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	if err := c.Request(req, &dst); err != nil {
		return err
	}

	return nil
}

func (c *Client) postForm(path string, body url.Values, dst interface{}) error {
	if !c.hasAuth() {
		return fmt.Errorf("this is private endpoint, please set api key and secret")
	}

	u, err := url.Parse(c.baseURL)
	if err != nil {
		return nil
	}
	u.Path = path

	body = c.populateSignature(body)

	req, err := http.NewRequest(http.MethodPost, u.String(), strings.NewReader(body.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return err
	}

	if err := c.Request(req, &dst); err != nil {
		return err
	}

	return nil
}

func (c *Client) deletePrivately(path string, query url.Values, dst interface{}) error {
	if !c.hasAuth() {
		return fmt.Errorf("this is private endpoint, please set api key and secret")
	}

	u, err := url.Parse(c.baseURL)
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

	if err := c.Request(req, &dst); err != nil {
		return err
	}

	return nil
}
