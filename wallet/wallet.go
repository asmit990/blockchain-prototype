package wallet

import (
	"crypto/ecdsa"    
	"crypto/elliptic" 
	"crypto/rand"     
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/ripemd160"
	"github.com/btcsuite/btcutil/base58"
)

type Wallet struct {
	privateKey        *ecdsa.PrivateKey
	publicKey         *ecdsa.PublicKey
	blockchainAddress string
}

func NewWallet() *Wallet {
	w := new(Wallet)

	// Step 1: Generate Private and Public Keys
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.privateKey = privateKey
	w.publicKey = &w.privateKey.PublicKey

	// Step 2: Perform SHA-256 hashing on public key
	h2 := sha256.New()
	h2.Write(w.publicKey.X.Bytes())
	h2.Write(w.publicKey.Y.Bytes())
	digest2 := h2.Sum(nil)

	// Step 3: Perform RIPEMD-160 hashing on SHA-256 result
	h3 := ripemd160.New()
	h3.Write(digest2)
	digest3 := h3.Sum(nil)

	// Step 4: Add version byte in front of RIPEMD-160 hash
	vd4 := make([]byte, 21)
	vd4[0] = 0x00 // Bitcoin version prefix
	copy(vd4[1:], digest3)

	// Step 5: Perform SHA-256 hash on the extended RIPEMD-160 result
	h5 := sha256.New()
	h5.Write(vd4)
	digest5 := h5.Sum(nil)

	// Step 6: Perform SHA-256 hash again
	h6 := sha256.New()
	h6.Write(digest5)
	digest6 := h6.Sum(nil)

	// Step 7: Take first 4 bytes as checksum
	chsum := digest6[:4]

	// Step 8: Append checksum to extended RIPEMD-160 hash
	dc8 := make([]byte, 25)
	copy(dc8[:21], vd4[:])
	copy(dc8[21:], chsum[:])

	// Step 9: Encode in Base58
	address := base58.Encode(dc8)
	w.blockchainAddress = address

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
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
}

func (w *Wallet) BlockchainAddress() string {
	return w.blockchainAddress
}
