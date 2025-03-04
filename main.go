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
}
