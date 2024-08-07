package aesgcm

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

// Encrypt encrypts plainText using AES GCM mode.
func Encrypt(plainText, key string) (string, error) {
	k := []byte(key)

	// use first 32byte if the key length is longer than 32byte.
	if len(k) > 32 {
		k = k[0:32]
	}

	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return "", err
	}

	cipherText := gcm.Seal(nil, nonce, []byte(plainText), nil)
	cipherText = append(nonce, cipherText...)

	return string(cipherText), nil
}

// Decrypt decrypts cipherText using AES GCM mode.
func Decrypt(cipherText, key string) (string, error) {
	k := []byte(key)
	ct := []byte(cipherText)

	// use first 32byte if the key length is longer than 32byte.
	if len(k) > 32 {
		k = k[0:32]
	}

	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ct) < nonceSize {
		return "", fmt.Errorf("cipherText is too short: textsize=[%d], noncesize=[%d]", len(ct), nonceSize)
	}

	nonce := ct[:nonceSize]
	plainByte, err := gcm.Open(nil, nonce, ct[nonceSize:], nil)
	if err != nil {
		return "", err
	}

	return string(plainByte), nil
}
