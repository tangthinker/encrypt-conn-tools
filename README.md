# encrypt-ecdh

[中文文档](README-zh.md)

Simple and easy-to-use Go library for common cryptographic operations. It provides high-level wrappers for ECDH key exchange, ECDSA signing/verification, and AES-256-GCM encryption.

## Features

- **ECDH (Elliptic Curve Diffie-Hellman)**: Securely generate shared secrets using P-256 curve.
- **ECDSA (Elliptic Curve Digital Signature Algorithm)**: Sign and verify data using P-256 curve keys.
- **AES-256-GCM**: Authenticated encryption and decryption with 256-bit keys.
- **Hex Encoding**: All keys, signatures, and encrypted data are handled as hex-encoded strings for easy storage and transmission.

## Installation

```bash
go get github.com/tangthinker/encrypt-conn-tools
```

## Usage

Import the package in your Go project:

```go
import "github.com/tangthinker/encrypt-conn-tools/pkg"
```

## Building Shared Library

```shell
go install mvdan.cc/garble@latest
# Build with garble
# -tiny: Optimize binary size
# -literals: Obfuscate string literals
# -seed=random: Randomize the build seed
garble -tiny -literals -seed=random build -buildmode=c-shared -o libencrypt.so main.go
```

### Cross Compilation

Cross-compilation is supported using standard `GOOS` and `GOARCH` environment variables. Note: Since `c-shared` mode is used, you must have a C cross-compiler for the target platform installed and CGO enabled.

```shell
# Example: Build for Linux (x86_64) on macOS
# Requires x86_64-linux-musl-gcc or similar
GOOS=linux GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-linux-musl-gcc \
garble -tiny -literals -seed=random build -buildmode=c-shared -o libencrypt.so main.go

# Example: Build for Windows (x86_64)
# Requires x86_64-w64-mingw32-gcc
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc \
garble -tiny -literals -seed=random build -buildmode=c-shared -o libencrypt.dll main.go
```

### ECDH Key Exchange

Generate key pairs and compute a shared secret.

```go
// 1. Generate key pairs for Alice and Bob
alicePub, alicePriv := pkg.GenerateKeyPairECDH()
bobPub, bobPriv := pkg.GenerateKeyPairECDH()

// 2. Compute shared secret (Alice uses her private key and Bob's public key)
sharedKeyAlice := pkg.GenerateSharedKey(alicePriv, bobPub)

// 3. Compute shared secret (Bob uses his private key and Alice's public key)
sharedKeyBob := pkg.GenerateSharedKey(bobPriv, alicePub)

// The shared keys should be identical
if sharedKeyAlice == sharedKeyBob {
    println("Shared secret established successfully!")
}
```

### ECDSA Signing & Verification

Sign data with a private key and verify it with a public key.

```go
// 1. Generate ECDSA key pair
pubKey, privKey := pkg.GenerateKeyPairECDSA()

data := "important transaction data"

// 2. Sign data
signature := pkg.SignECDSA(data, privKey)

// 3. Verify signature
valid := pkg.VerifyECDSA(data, pubKey, signature)
if valid {
    println("Signature is valid!")
}
```

### AES-256 Encryption

Encrypt and decrypt data using a 32-byte hex-encoded key.

```go
// Derive a 32-byte key (e.g., from an ECDH shared secret)
// The DeriveKey function helps create a suitable key from input factors
key := pkg.DeriveKey("some-shared-secret", "salt") 

plaintext := "Hello, World!"

// 1. Encrypt
ciphertext := pkg.Encrypt(plaintext, key)
println("Encrypted:", ciphertext)

// 2. Decrypt
decrypted := pkg.Decrypt(ciphertext, key)
println("Decrypted:", decrypted)
```

## License

MIT
