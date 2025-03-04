package wallet 

import (
	"crypto/ecdsa"    // ECDSA (Elliptic Curve Digital Signature Algorithm) keys ke liye
	"crypto/elliptic" // Elliptic Curve ka model define karne ke liye (P-256)
	"crypto/rand"     // Random number generator for secure key generation
	"fmt"
)




type Wallet struct {
	privateKey *ecdsa.PrivateKey  // Private Key (Secret key, should NOT be shared)
	publicKey  *ecdsa.PublicKey   // Public Key (Can be shared with others)
}

func NewWallet() *Wallet {
	// private key and public key ke baare mei hai
    w := new(Wallet) // Ek naya Wallet struct banaya
    privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader) 
    w.privateKey = privateKey  // Private key ko wallet me store kiya
    w.publicKey = &w.privateKey.PublicKey // Public key assign ki
    return w
}


func (w *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey 
}

func (w *Wallet) PrivateKeyStr() string {
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
 }

func (w *Wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey 
}
	
func (w *Wallet) PublicKeyStr() string {
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes()) //  Public key as hex string
}