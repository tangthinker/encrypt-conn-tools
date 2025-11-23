package pkg

import (
	"crypto/ecdh"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

// GenerateKeyPair generates a P256 ECDH key pair.
// Returns hex-encoded public and private keys.
func GenerateKeyPairECDH() (pubKey, privKey string) {
	k, err := ecdh.P256().GenerateKey(rand.Reader)
	if err != nil {
		return "", ""
	}
	return hex.EncodeToString(k.PublicKey().Bytes()), hex.EncodeToString(k.Bytes())
}

// GenerateSharedKey computes the ECDH shared secret from a private key and a peer's public key.
// Input keys must be hex-encoded. Returns the hex-encoded shared secret.
func GenerateSharedKey(privKey string, pubKey string) string {
	privBytes, err := hex.DecodeString(privKey)
	if err != nil {
		return ""
	}
	pubBytes, err := hex.DecodeString(pubKey)
	if err != nil {
		return ""
	}

	pKey, err := ecdh.P256().NewPrivateKey(privBytes)
	if err != nil {
		return ""
	}
	pub, err := ecdh.P256().NewPublicKey(pubBytes)
	if err != nil {
		return ""
	}

	secret, err := pKey.ECDH(pub)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(secret)
}

// DeriveKey derives a key by hashing the provided factors using SHA256.
// Returns a hex-encoded string.
func DeriveKey(factors ...string) string {
	h := sha256.New()
	for _, f := range factors {
		h.Write([]byte(f))
	}
	return hex.EncodeToString(h.Sum(nil))
}
