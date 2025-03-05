package wallet

import (
	"crypto/ecdsa"    // Elliptic Curve Digital Signature Algorithm (Private-Public Key Pair)
	"crypto/elliptic" // Elliptic Curve Cryptography (P-256 Curve)
	"crypto/rand"     // Randomness for Secure Key Generation
	"crypto/sha256"   // SHA-256 Hashing Algorithm
	"fmt"             // Formatting for String Conversion
	"math/big"

	"github.com/btcsuite/btcutil/base58" // Base58 Encoding (Bitcoin Style Address)
	"golang.org/x/crypto/ripemd160"      // RIPEMD-160 Hashing (Bitcoin Address)
)

// 🚀 **Wallet Struct: Yeh Tera Digital Wallet Hai!**
type Wallet struct {
	privateKey        *ecdsa.PrivateKey  // Tera Secret Key (Don't Share!)
	publicKey         *ecdsa.PublicKey   // Tera Public Key (Sabko Dikh Sakti Hai)
	blockchainAddress string             // Final Wallet Address (Bitcoin Style)
}

// 🚀 **NewWallet(): Yeh Tera Naya Wallet Generate Karega!**
func NewWallet() *Wallet {
	w := new(Wallet) // Naya Wallet Bana Diya!

	// 🔥 **Step 1: Private & Public Key Generate Karna**
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader) // P-256 Elliptic Curve
	w.privateKey = privateKey
	w.publicKey = &w.privateKey.PublicKey // Private Key Se Public Key Nikali

	// 🔥 **Step 2: Public Key Ka SHA-256 Hash Banana (Security Badhane Ke Liye)**
	h2 := sha256.New()
	h2.Write(w.publicKey.X.Bytes()) // X Coordinate Hash
	h2.Write(w.publicKey.Y.Bytes()) // Y Coordinate Hash
	digest2 := h2.Sum(nil) // Final SHA-256 Hash

	// 🔥 **Step 3: RIPEMD-160 Hash (Bitcoin Wala Short Hash)**
	h3 := ripemd160.New()
	h3.Write(digest2) // SHA-256 Ka Output -> RIPEMD-160 Input
	digest3 := h3.Sum(nil) // Final RIPEMD-160 Hash

	// 🔥 **Step 4: Version Byte Add Karna (Bitcoin Ke Liye 0x00)**
	vd4 := make([]byte, 21) // 21 Bytes: 1 Byte Version + 20 Bytes RIPEMD-160
	vd4[0] = 0x00 // 0x00 -> Bitcoin Wale Addresses Ke Liye Prefix
	copy(vd4[1:], digest3) // RIPEMD-160 Hash Copy Kar Diya

	// 🔥 **Step 5: Double SHA-256 Lagana (Checksum Generate Karne Ke Liye)**
	h5 := sha256.New()
	h5.Write(vd4) // Version + RIPEMD-160 Hash
	digest5 := h5.Sum(nil) // SHA-256 Round 1

	h6 := sha256.New()
	h6.Write(digest5) // SHA-256 Round 2
	digest6 := h6.Sum(nil)

	// 🔥 **Step 7: First 4 Bytes (Checksum) Nikalna**
	chsum := digest6[:4] // Pehle 4 Bytes = Checksum

	// 🔥 **Step 8: Checksum Ko Address Ke Saath Attach Karna**
	dc8 := make([]byte, 25) // 25 Bytes: 21 Bytes Data + 4 Bytes Checksum
	copy(dc8[:21], vd4[:])  // Pehle 21 Bytes: Version + RIPEMD-160
	copy(dc8[21:], chsum[:]) // Last 4 Bytes: Checksum

	// 🔥 **Step 9: Base58 Encoding (Bitcoin Style Address)**
	address := base58.Encode(dc8) // Base58 Encoding
	w.blockchainAddress = address // Address Wallet Me Store Kar Diya

	return w // Wallet Wapas Return Karo 🚀
}

// 🚀 **Private Key Getter: Yeh Tera Secret Key Hai!**
func (w *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

// 🚀 **Private Key as String: Debug Ke Liye (Don't Share!)**
func (w *Wallet) PrivateKeyStr() string {
	return fmt.Sprintf("%x", w.privateKey.D.Bytes()) // Hexadecimal Format Me Convert
}

// 🚀 **Public Key Getter: Public Key Wapas Dega**
func (w *Wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

// 🚀 **Public Key as String: Hexadecimal Format Me!**
func (w *Wallet) PublicKeyStr() string {
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes()) // X + Y Coordinate Ka Combo
}

// 🚀 **Blockchain Address Getter: Final Wallet Address De Raha Hai!**
func (w *Wallet) BlockchainAddress() string {
	return w.blockchainAddress
}






	type Transaction struct {
		senderPrivateKey         *ecdsa.PrivateKey  // 🚀 Sender Ki Secret Private Key (Signing Ke Liye Use Hogi)
		senderPublicKey          *ecdsa.PublicKey   // 🔑 Sender Ki Public Key (Verify Karne Ke Liye)
		senderBlockchainAddress  string             // 🏦 Sender Ka Wallet Address (Jis Se Paise Katne Hain)
		recipientBlockchainAddress string           // 🎯 Receiver Ka Wallet Address (Jisko Paise Milenge)
		value                    float32            // 💰 Transfer Hone Wali Amount (Bitcoin Style)
	}
	
func NewTransaction(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey,
sender string, recipient string, value float32) *Transaction {
	return &Transaction{privateKey,  publicKey, sender, recipient, value}

}
func (t *Transaction) GenerateSignature() *Signature {
  m, - := json.Marshall(t)
  h := sha265.Sum265([]byte(m))
  r, s, _ := ecdsa.Sign(rand.Reader, t.senderPrivateKey, h[:])
  return  &Signature{r, s}
}
func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender string `json:"sender_blockchain_address"`
		Recipient string `json:"recipient_blockchain_address"`
		Value float32 `json:"value"`
	}{
		Sender; t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress
	})
}
type Signature struct {
	R *big.Int
	S *big.Int
}
func (s *Signature) String() string {
	return fmt.Sprint("%x\x",s.R,s.S)
}