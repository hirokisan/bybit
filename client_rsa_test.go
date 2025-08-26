package bybit

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExampleClient_WithAuthRSA(t *testing.T) {
	// Variables to capture the request
	var capturedRequest *http.Request
	var capturedBody []byte

	// Create test server that captures the request
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Capture the request
		capturedRequest = r
		// Read and capture the body if present
		if r.Body != nil {
			body, _ := io.ReadAll(r.Body)
			capturedBody = body
		}
		// Response is bogus, just for testing
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"retCode": 0,
			"retMsg":  "OK",
			"result": map[string]interface{}{
				"list": []interface{}{
					map[string]interface{}{
						"accountType": "UNIFIED",
						"coin":        []interface{}{},
					},
				},
			},
		})
	}))
	defer server.Close()
	// instantiate w/ RSA
	client := NewClient().
		WithBaseURL(server.URL).
		WithAuthRSA(TestAPIKey, TestRSAPrivateKey)
	// useRSA is how the client determines if it should use RSA signing
	require.True(t, client.useRSA)
	require.True(t, client.hasAuth())
	// Make some random API call
	resp, err := client.V5().Account().GetWalletBalance(AccountTypeV5UNIFIED, nil)
	require.NoError(t, err)
	// Grab inputs needed to recreate signature
	tsstr := capturedRequest.Header.Get("X-BAPI-TIMESTAMP")
	ts, err := strconv.ParseInt(tsstr, 10, 64)
	if err != nil {
		t.Fatalf("failed to parse timestamp: %v", err)
	}
	rcv := capturedRequest.Header.Get("X-BAPI-RECV-WINDOW")
	// Load up RSA key
	p, err := LoadRSAPrivateKeyFromBytes([]byte(TestRSAPrivateKey))
	require.NoError(t, err)
	// Rebuild sig
	query := capturedRequest.URL.Query().Encode()
	expectedSig := getV5SignatureRSA(
		ts,
		TestAPIKey,
		rcv,
		query,
		p,
	)
	actualSig := capturedRequest.Header.Get("X-BAPI-SIGN")
	require.Equal(t, expectedSig, actualSig)
	// verify signature matches private key and payload (timestamp + api_key + recv_window + queryString)
	payload := strconv.FormatInt(ts, 10) + TestAPIKey + rcv + query
	hash := sha256.Sum256([]byte(payload))
	// Ensure signer match
	sigBytes, err := base64.StdEncoding.DecodeString(actualSig)
	require.NoError(t, err)
	err = rsa.VerifyPKCS1v15(&p.PublicKey, crypto.SHA256, hash[:], sigBytes)
	require.NoError(t, err)
	require.NotNil(t, resp)

	// The client will automatically use RSA signatures instead of HMAC
	t.Log("Client created with RSA authentication", capturedBody)
}

// Verify our helper
func TestLoadRSAPrivateKeyFromFile(t *testing.T) {
	// Write to file
	tmp := os.TempDir()
	err := os.WriteFile(tmp+"/private_key.pem", []byte(TestRSAPrivateKey), 0600)
	if err != nil {
		panic(err)
	}
	// Load private key from file
	privateKeyPEM, err := LoadRSAPrivateKeyFromFile(tmp + "/private_key.pem")
	if err != nil {
		panic(err)
	}

	// Create client with RSA authentication
	client := NewClient().
		WithAuthRSA(TestAPIKey, privateKeyPEM)
	require.True(t, client.useRSA)
	require.True(t, client.hasAuth())
	_ = client
}

func TestRSAModeWebSocket(t *testing.T) {
	// Example RSA private key in PEM format
	privateKeyPEM := TestRSAPrivateKey

	// Create WebSocket client with RSA authentication
	wsClient := NewWebsocketClient().
		WithAuthRSA(TestAPIKey, privateKeyPEM)
	require.True(t, wsClient.useRSA)

	// Use the WebSocket client
	// The client will automatically use RSA signatures for authentication
	_ = wsClient
}

// TestRSAPrivateKey for testing only DO NOT USE ON YOUR OWN ACCOUNT
var (
	TestRSAPrivateKey = `-----BEGIN PRIVATE KEY-----
MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDGpWdZZMSosRfq
9vl1Gbz8Op8olzAjp6DYIjaDarlMbCa4AjcxD1bAVCqsIIHdmkAGceTWqZ2wcq7q
9h9diGr5Ai76HnA74ssIRd9aiKOV4wIF5GYkIrfmvj0XlRfMShgs3vHwFgtSBtVc
9UyNI8OhvAvoOuJpXMB8MPXanrluKgTeZ+Msh1c3z1ZYP/bEgPJtvxxppwuxXOFk
EzWwga8+VYt5sddAxC9xe3Tze1aC+2U+MQJXUYBEtiJg83J98DuiARkUXcfzEvWc
pKCBxzy8yg7K2ATplt50kgXpANyJSds3k22BwDTuBFOCXchjJ5uOjd/NDH65UALA
Gp8RdJ2pAgMBAAECggEAEtpn8KgLsib8qiQ+kMqckSPPLz8KJpgmP90fZ2GdrI8n
LZeoPVPDXAWg9y0upiWZpgLxQR5gBPsddtddKrdjz7ZRVxPEeyqPMn3PMPx+6h7d
dkmzGALjiVbM3Ywc8hqmKLwQIkU1VdeoMGZnmY2nZ5Y9WxEbo/xVmHPdIZ6m4rmX
YZ6XDnmHMDDcJFRqQQcFqWqxA9qcgfhbI4hcV34E0dUXOl2aSFO6oBfcTru8B1Cf
nCgk37HyqNwgVpqwDype/VKvj3PQW4x5bobzGm8JdeW4Py/xEX5FeYIsbFXrRZ0w
n3HS2Olj3Fw5UgjaO4CAMmZQX6m2VLFElr8tZffkAQKBgQDoUfcUxv1AKQtLCi3N
B7R6YfqKih59e/1j7vIVFT7IkNAZFQUIHBnr/t9yquSJIzGjTi2uoJ3w1DWm42Cx
E+aQMs/QEaW8jqZASaINOC87sT3FNJD4wPQR/88IIqr87JZ2LYFr4iXITZx5xSet
o8R0QqWpEImFOFcwmf0HWpFRSQKBgQDa5MQU0IZmVrUMkrLmgk1tp/zF9Mk9hAfc
dCmHV4BLmi5QVlFpps4IBgQJ6hbvfeN78AXrLhtYVmAqYUwFAD7BrAloiBhcaqV3
IpfT7iLKrb4VCFvHxzO3OlUae1URBTjRVwQ/aW4bZE/o9O56tmKB422dYl9/P8t+
rkZAFX9JYQKBgC6ci2NazWL7GS30G95gJmDLmbYEIjvxDZToUx/RxGf/ThFKO1k2
Mik8WN6r1PCC5CmsvNOlnCq+mQkj47mDkaXq2/EWKVeck1SgsWfPlwJ1/Du94TxX
kmCuH361XfMjEMkjNi1MEWKP185CtURMcFUXLh+ulrjo5e11Z+P60t8pAoGAYFGl
zprraQDibgPnYPMZaxUub1UFcGI2q1UaKQnh1GKl2ogBDwJtSq6K/Gnbacr1XMYD
dLc3JSns6vkhYFn5Q3OWOD8aqR/sa333XTQ+bv0A1XR2HnSTVx7978cxaWno1IT7
w4N7BeagGxwcDDdRJWKUC1sMNow12SqKkwxilEECgYBBZsGnzoskGylfEJ9TrlpG
6WHrgCCcg6Gltr7IctyCfwzCmynEyo6/vei7sqAsp5YcomRaPd2I4pqeFiSMvdtn
HyF4dlEMZ7mC+fZjkD9EVyPzBKJ2mGm3hF5mfVmASm93ielzf3DPVbd/eiKQZ5Uu
0JHHWp6wvtJeC833k4iXAw==
-----END PRIVATE KEY-----`
	// This is a test API key, DO NOT USE ON YOUR OWN ACCOUNT
	TestAPIKey = "0s4stkea3xIc44bYoq"
)
