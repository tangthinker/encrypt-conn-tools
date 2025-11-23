package pkg

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
)

// GenerateKeyPairECDSA generates a P256 ECDSA key pair.
// Returns hex-encoded PKIX public key and SEC 1 private key.
func GenerateKeyPairECDSA() (pubKey, privKey string) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", ""
	}

	privBytes, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return "", ""
	}

	pubBytes, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	if err != nil {
		return "", ""
	}

	return hex.EncodeToString(pubBytes), hex.EncodeToString(privBytes)
}

// SignECDSA signs the data using the private key.
// Input private key must be hex-encoded SEC 1. Returns hex-encoded ASN.1 signature.
func SignECDSA(data string, privKey string) string {
	privBytes, err := hex.DecodeString(privKey)
	if err != nil {
		return ""
	}

	key, err := x509.ParseECPrivateKey(privBytes)
	if err != nil {
		return ""
	}

	hash := sha256.Sum256([]byte(data))

	sig, err := ecdsa.SignASN1(rand.Reader, key, hash[:])
	if err != nil {
		return ""
	}

	return hex.EncodeToString(sig)
}

// VerifyECDSA verifies the signature of the data using the public key.
// Input public key must be hex-encoded PKIX. Input signature must be hex-encoded ASN.1.
func VerifyECDSA(data string, pubKey string, signature string) bool {
	pubBytes, err := hex.DecodeString(pubKey)
	if err != nil {
		return false
	}

	pub, err := x509.ParsePKIXPublicKey(pubBytes)
	if err != nil {
		return false
	}

	ecdsaPub, ok := pub.(*ecdsa.PublicKey)
	if !ok {
		return false
	}

	sigBytes, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}

	hash := sha256.Sum256([]byte(data))

	return ecdsa.VerifyASN1(ecdsaPub, hash[:], sigBytes)
}
