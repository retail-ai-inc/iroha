package chacha20poly1305

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/chacha20poly1305"
)

const (
	KeySize    = chacha20poly1305.KeySize
	nonceSizeX = chacha20poly1305.NonceSizeX
)

// Encrypt encrypts plainText using XChaCha20-Poly1305.
func Encrypt(plainText, key string) (string, error) {
	k := []byte(key)

	// use first 32byte if the key length is longer than 32byte.
	if len(k) > KeySize {
		k = k[0:KeySize]
	}

	aead, err := chacha20poly1305.NewX(k)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, nonceSizeX)
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}

	cipherText := aead.Seal(nil, nonce, []byte(plainText), nil)
	cipherText = append(nonce, cipherText...)
	return string(cipherText), nil
}

// Decrypt decrypts cipherText using XChaCha20-Poly1305.
func Decrypt(cipherText, key string) (string, error) {
	k := []byte(key)
	ct := []byte(cipherText)

	// use first 32byte if the key length is longer than 32byte.
	if len(k) > KeySize {
		k = k[0:KeySize]
	}

	aead, err := chacha20poly1305.NewX(k)
	if err != nil {
		return "", err
	}

	if len(ct) < nonceSizeX {
		return "", fmt.Errorf("cipherText is too short: textsize=[%d], noncesize=[%d]", len(ct), nonceSizeX)
	}

	nonce := ct[:nonceSizeX]
	plainByte, err := aead.Open(nil, nonce, ct[nonceSizeX:], nil)
	if err != nil {
		return "", err
	}

	return string(plainByte), nil
}
