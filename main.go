package main


import(
	"log"
	"fmt"
	"blockchain/wallet"  
)

func init() {
	log.SetPrefix("BlockChain : ")
}

func main() {
  w := wallet.NewWallet()
  fmt.Println(w.PrivateKeyStr())
  fmt.Println(w.PublicKeyStr())
  fmt.Println(w.BlockchainAddress())


  t := wallet.NewTransaction(w.PrivateKey(), w.PublicKey(), w.BlockchainAddress(), "B", 1.0)
   fmt.Printf("signature %s\n", t.GenerateSignature())
}
