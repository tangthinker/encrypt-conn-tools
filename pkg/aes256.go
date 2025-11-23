package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

// Encrypt encrypts plaintext using AES-256-GCM.
// key must be a hex-encoded 32-byte string (64 characters).
// Returns a hex-encoded string containing the nonce and ciphertext.
func Encrypt(plaintext string, key string) string {
	keyBytes, err := hex.DecodeString(key)
	if err != nil || len(keyBytes) != 32 {
		return ""
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return ""
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return ""
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return ""
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)
	return hex.EncodeToString(ciphertext)
}

// Decrypt decrypts a hex-encoded ciphertext using AES-256-GCM.
// key must be a hex-encoded 32-byte string (64 characters).
func Decrypt(ciphertext string, key string) string {
	keyBytes, err := hex.DecodeString(key)
	if err != nil || len(keyBytes) != 32 {
		return ""
	}

	data, err := hex.DecodeString(ciphertext)
	if err != nil {
		return ""
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return ""
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return ""
	}

	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return ""
	}

	nonce, ciphertextBytes := data[:nonceSize], data[nonceSize:]
	plaintextBytes, err := aesGCM.Open(nil, nonce, ciphertextBytes, nil)
	if err != nil {
		return ""
	}

	return string(plaintextBytes)
}
