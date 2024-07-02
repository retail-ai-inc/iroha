package aescbc

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func Encrypt(plaintext, key, initialvector string) (string, error) {
	k := []byte(key)
	iv := []byte(initialvector)

	bPlaintext := PKCS5Padding([]byte(plaintext), aes.BlockSize)
	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}

	// use first `aes.BlockSize`` byte if the key length is longer than `aes.BlockSize`.
	if len(iv) > aes.BlockSize {
		iv = iv[0:aes.BlockSize]
	}

	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, bPlaintext)

	return string(ciphertext), nil
}

// Decrypt decrypts cipherText using AES 256 CBC mode.
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

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
