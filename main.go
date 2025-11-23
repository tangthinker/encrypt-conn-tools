package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"encoding/json"
	"unsafe"

	"github.com/tangthinker/encrypt-conn-tools/pkg"
)

// KeyPair holds public and private keys for JSON marshaling
type KeyPair struct {
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

//export GenerateKeyPairECDH
func GenerateKeyPairECDH() *C.char {
	pub, priv := pkg.GenerateKeyPairECDH()
	kp := KeyPair{PublicKey: pub, PrivateKey: priv}
	jsonData, _ := json.Marshal(kp)
	return C.CString(string(jsonData))
}

//export GenerateSharedKey
func GenerateSharedKey(privKey *C.char, pubKey *C.char) *C.char {
	shared, err := pkg.GenerateSharedKey(C.GoString(privKey), C.GoString(pubKey))
	if err != nil {
		return C.CString("")
	}
	return C.CString(shared)
}

//export GenerateKeyPairECDSA
func GenerateKeyPairECDSA() *C.char {
	pub, priv := pkg.GenerateKeyPairECDSA()
	kp := KeyPair{PublicKey: pub, PrivateKey: priv}
	jsonData, _ := json.Marshal(kp)
	return C.CString(string(jsonData))
}

//export SignECDSA
func SignECDSA(data *C.char, privKey *C.char) *C.char {
	sig := pkg.SignECDSA(C.GoString(data), C.GoString(privKey))
	return C.CString(sig)
}

//export VerifyECDSA
func VerifyECDSA(data *C.char, pubKey *C.char, signature *C.char) C.int {
	valid := pkg.VerifyECDSA(C.GoString(data), C.GoString(pubKey), C.GoString(signature))
	if valid {
		return 1
	}
	return 0
}

//export Encrypt
func Encrypt(plaintext *C.char, key *C.char) *C.char {
	res := pkg.Encrypt(C.GoString(plaintext), C.GoString(key))
	return C.CString(res)
}

//export Decrypt
func Decrypt(ciphertext *C.char, key *C.char) *C.char {
	res := pkg.Decrypt(C.GoString(ciphertext), C.GoString(key))
	return C.CString(res)
}

//export DeriveKey
func DeriveKey(factor *C.char) *C.char {
	res := pkg.DeriveKey(C.GoString(factor))
	return C.CString(res)
}

//export FreeString
func FreeString(str *C.char) {
	C.free(unsafe.Pointer(str))
}

func main() {}
