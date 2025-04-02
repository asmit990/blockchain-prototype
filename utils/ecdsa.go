package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"math/big"
	"fmt"
)

// Signature struct
type Signature struct {
	R *big.Int
	S *big.Int
}

func (s *Signature) String() string {
	return fmt.Sprintf("%064x%064x", s.R, s.S)
}

// Converts string to BigInt tuple (X, Y)
func String2BigIntTuple(s string) (big.Int, big.Int) {
	if len(s) < 128 { // Ensure string is valid
		panic("Invalid string length for BigInt tuple")
	}
	bx, _ := hex.DecodeString(s[:64])
	by, _ := hex.DecodeString(s[64:])
	var bix, biy big.Int
	bix.SetBytes(bx)
	biy.SetBytes(by)
	return bix, biy
}

// Convert Public Key from string
func PublicKeyFromString(s string) *ecdsa.PublicKey {
	x, y := String2BigIntTuple(s)
	return &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     &x,
		Y:     &y,
	}
}

// Convert Signature from string
func SignatureFromString(s string) *Signature {
	x, y := String2BigIntTuple(s)
	return &Signature{R: &x, S: &y}
}

// Convert Private Key from string
func PrivateKeyFromString(s string, publicKey *ecdsa.PublicKey) *ecdsa.PrivateKey {
	b, _ := hex.DecodeString(s) // Decode full string
	var bi big.Int
	bi.SetBytes(b)

	return &ecdsa.PrivateKey{
		PublicKey: *publicKey,
		D:         &bi,
	}
}
