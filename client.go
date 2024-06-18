package bybit

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
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
	// MainNetBaseURL2 :
	MainNetBaseURL2 = "https://api.bytick.com"
)

// Client :
type Client struct {
	httpClient *http.Client

	debug  bool
	logger *log.Logger

	baseURL string
	key     string
	secret  string

	referer string

	checkResponseBody        checkResponseBodyFunc
	syncTimeDeltaNanoSeconds int64
}

func (c *Client) debugf(format string, v ...interface{}) {
	if c.debug {
		c.logger.Printf(format, v...)
	}
}

// NewClient :
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},

		logger: newDefaultLogger(),

		baseURL:           MainNetBaseURL,
		checkResponseBody: checkResponseBody,
	}
}

// WithHTTPClient :
func (c *Client) WithHTTPClient(httpClient *http.Client) *Client {
	c.httpClient = httpClient

	return c
}

// WithDebug :
func (c *Client) WithDebug(debug bool) *Client {
	c.debug = debug

	return c
}

// WithLogger :
func (c *Client) WithLogger(logger *log.Logger) *Client {
	c.debug = true
	c.logger = logger

	return c
}

// WithAuth :
func (c *Client) WithAuth(key string, secret string) *Client {
	c.key = key
	c.secret = secret

	return c
}

func (c Client) withCheckResponseBody(f checkResponseBodyFunc) *Client {
	c.checkResponseBody = f

	return &c
}

// WithBaseURL :
func (c *Client) WithBaseURL(url string) *Client {
	c.baseURL = url

	return c
}

func (c *Client) WithReferer(referer string) *Client {
	c.referer = referer

	return c
}

// Request :
func (c *Client) Request(req *http.Request, dst interface{}) (err error) {
	c.debugf("request: %v", req)
	resp, err := c.httpClient.Do(req)
	c.debugf("response: %v", resp)
	if err != nil {
		return err
	}
	c.debugf("response status code: %v", resp.StatusCode)
	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = cerr
		}
	}()

	switch {
	case 200 <= resp.StatusCode && resp.StatusCode <= 299:
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if c.checkResponseBody == nil {
			return errors.New("checkResponseBody func should be set")
		}
		if err := c.checkResponseBody(body); err != nil {
			return err
		}

		if err := json.Unmarshal(body, &dst); err != nil {
			return err
		}

		c.debugf("response body: %v", string(body))
		return nil
	case resp.StatusCode == http.StatusBadRequest:
		return fmt.Errorf("%v: Need to send the request with GET / POST (must be capitalized)", ErrBadRequest)
	case resp.StatusCode == http.StatusUnauthorized:
		return fmt.Errorf("%w: invalid key/secret", ErrInvalidRequest)
	case resp.StatusCode == http.StatusForbidden:
		return fmt.Errorf("%w: not permitted", ErrForbiddenRequest)
	case resp.StatusCode == http.StatusNotFound:
		return fmt.Errorf("%w: wrong path", ErrPathNotFound)
	default:
		return fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}
}

// hasAuth : check has auth key and secret
func (c *Client) hasAuth() bool {
	return c.key != "" && c.secret != ""
}

func (c *Client) populateSignature(src url.Values) url.Values {
	if src == nil {
		src = url.Values{}
	}

	src.Add("api_key", c.key)
	src.Add("timestamp", strconv.FormatInt(c.getTimestamp(), 10))
	if c.referer != "" {
		src.Add("referer", c.referer)
	}
	src.Add("sign", getSignature(src, c.secret))

	return src
}

func (c *Client) populateSignatureForBody(src []byte) []byte {
	body := map[string]interface{}{}
	if err := json.Unmarshal(src, &body); err != nil {
		panic(err)
	}

	body["api_key"] = c.key
	body["timestamp"] = strconv.FormatInt(c.getTimestamp(), 10)
	if c.referer != "" {
		body["referer"] = c.referer
	}
	body["sign"] = getSignatureForBody(body, c.secret)

	result, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	return result
}

func getV5Signature(
	timestamp int64,
	key string,
	queryString string,
	secret string,
) string {
	val := strconv.FormatInt(timestamp, 10) + key
	val = val + queryString
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(val))
	return hex.EncodeToString(h.Sum(nil))
}

func getV5SignatureForBody(
	timestamp int64,
	key string,
	body []byte,
	secret string,
) string {
	val := strconv.FormatInt(timestamp, 10) + key
	val = val + string(body)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(val))
	return hex.EncodeToString(h.Sum(nil))
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

func getSignatureForBody(src map[string]interface{}, key string) string {
	keys := make([]string, len(src))
	i := 0
	_val := ""
	for k := range src {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		_val += k + "=" + fmt.Sprintf("%v", src[k]) + "&"
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

func (c *Client) getPubliclyWithContext(ctx context.Context, path string, query url.Values, dst interface{}) error {
	u, err := url.Parse(c.baseURL)
	if err != nil {
		return err
	}
	u.Path = path
	u.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
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

func (c *Client) getV5Privately(path string, query url.Values, dst interface{}) error {
	if !c.hasAuth() {
		return fmt.Errorf("this is private endpoint, please set api key and secret")
	}

	u, err := url.Parse(c.baseURL)
	if err != nil {
		return err
	}
	u.Path = path
	u.RawQuery = query.Encode()

	timestamp := c.getTimestamp()
	sign := getV5Signature(timestamp, c.key, query.Encode(), c.secret)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-BAPI-API-KEY", c.key)
	req.Header.Set("X-BAPI-TIMESTAMP", strconv.FormatInt(timestamp, 10))
	req.Header.Set("X-BAPI-SIGN", sign)

	if err := c.Request(req, &dst); err != nil {
		return err
	}

	return nil
}

func (c *Client) getV5PrivatelyWithContext(ctx context.Context, path string, query url.Values, dst interface{}) error {
	if !c.hasAuth() {
		return fmt.Errorf("this is private endpoint, please set api key and secret")
	}

	u, err := url.Parse(c.baseURL)
	if err != nil {
		return err
	}
	u.Path = path
	u.RawQuery = query.Encode()

	timestamp := c.getTimestamp()
	sign := getV5Signature(timestamp, c.key, query.Encode(), c.secret)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-BAPI-API-KEY", c.key)
	req.Header.Set("X-BAPI-TIMESTAMP", strconv.FormatInt(timestamp, 10))
	req.Header.Set("X-BAPI-SIGN", sign)

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

	body = c.populateSignatureForBody(body)

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

func (c *Client) postV5JSON(path string, body []byte, dst interface{}) error {
	if !c.hasAuth() {
		return fmt.Errorf("this is private endpoint, please set api key and secret")
	}

	u, err := url.Parse(c.baseURL)
	if err != nil {
		return err
	}
	u.Path = path

	timestamp := c.getTimestamp()
	sign := getV5SignatureForBody(timestamp, c.key, body, c.secret)

	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-BAPI-API-KEY", c.key)
	req.Header.Set("X-BAPI-TIMESTAMP", strconv.FormatInt(timestamp, 10))
	req.Header.Set("X-BAPI-SIGN", sign)
	if c.referer != "" {
		req.Header.Set("X-Referer", c.referer)
	}

	if err := c.Request(req, &dst); err != nil {
		return err
	}

	return nil
}

func (c *Client) postV5JSONWithContext(ctx context.Context, path string, body []byte, dst interface{}) error {
	if !c.hasAuth() {
		return fmt.Errorf("this is private endpoint, please set api key and secret")
	}

	u, err := url.Parse(c.baseURL)
	if err != nil {
		return err
	}
	u.Path = path

	timestamp := c.getTimestamp()
	sign := getV5SignatureForBody(timestamp, c.key, body, c.secret)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-BAPI-API-KEY", c.key)
	req.Header.Set("X-BAPI-TIMESTAMP", strconv.FormatInt(timestamp, 10))
	req.Header.Set("X-BAPI-SIGN", sign)
	if c.referer != "" {
		req.Header.Set("X-Referer", c.referer)
	}

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

func (c *Client) getTimestamp() int64 {
	return (time.Now().UnixNano() - c.syncTimeDeltaNanoSeconds) / 1000000
}

func (c *Client) updateSyncTimeDelta(
	remoteServerTimeRaw string,
	localTimestampNanoseconds int64,
) error {
	remoteServerTimeNS, err := strconv.ParseInt(remoteServerTimeRaw, 10, 64)
	if err != nil {
		return fmt.Errorf("parse server time: %w", err)
	}

	c.syncTimeDeltaNanoSeconds = localTimestampNanoseconds - remoteServerTimeNS
	return nil
}

func (c *Client) SyncServerTime(ctx context.Context) error {
	r, err := c.NewTimeService().GetServerTime(ctx)
	if err != nil {
		return fmt.Errorf("get server time: %w", err)
	}

	if r.Result.TimeNano == "" {
		return errors.New("server time is empty")
	}

	return c.updateSyncTimeDelta(r.Result.TimeNano, time.Now().UnixNano())
}
