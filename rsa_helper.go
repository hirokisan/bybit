package bybit

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// LoadRSAPrivateKeyFromFile loads an RSA private key from a PEM file
func LoadRSAPrivateKeyFromFile(path string) (string, error) {
	keyData, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read private key file: %w", err)
	}
	
	return string(keyData), nil
}

// LoadRSAPrivateKeyFromBytes loads an RSA private key from PEM bytes
func LoadRSAPrivateKeyFromBytes(keyData []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the private key")
	}
	
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		// Try PKCS8 format if PKCS1 fails
		parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse private key: %w", err)
		}
		var ok bool
		privateKey, ok = parsedKey.(*rsa.PrivateKey)
		if !ok {
			return nil, fmt.Errorf("not an RSA private key")
		}
	}
	
	return privateKey, nil
}