package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	curve := elliptic.P256()
	privateKey, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	publicKey := append(x.Bytes(), y.Bytes()...)
	h := sha256.New()
	h.Write(publicKey)
	publicKeyHash := h.Sum(nil)

	// Double hash the public key hash using SHA-256
	h = sha256.New()
	h.Write(publicKeyHash)
	doubleHash := h.Sum(nil)

	// Truncate the double hash to 20 bytes
	addressBytes := doubleHash[:20]

	// Encode the address bytes in hexadecimal format
	address := hex.EncodeToString(addressBytes)

	fmt.Println("Private Key: ", privateKey)
	fmt.Println("Public Key (X, Y): ", x, y)
	fmt.Println("Address: ", address)
}
