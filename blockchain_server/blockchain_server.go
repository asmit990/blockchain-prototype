package main

import(
	"io"
	"net/http"
	"log"
	"strconv"
   "blockchain/block"
   "blockchain/wallet"
  "flag"
    
)

func init() {
    log.SetPrefix("BlockChain : ")
}
var cache map[string]*block.Blockchain = make(map[string]*block.Blockchain)
type BlockchainServer struct {
	port uint16
}

func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{port}
}
func( bcs *BlockchainServer)  Port() uint16 {
	return bcs.port 
}
func(bcs *BlockchainServer) GetBlockchain() *block.Blockchain {
	bc, ok := cache["blockchain"]
	if !ok {
		minerWallet := wallet.NewWallet()
		bc = block.NewBlockchain(minerWallet.BlockchainAddress(), bcs.Port()) 
		log.Printf("private_key %v", minerWallet.PrivateKeyStr())
		log.Printf("public_key %v", minerWallet.PublicKeyStr())
		log.Printf("blockchain_address %v", minerWallet.BlockchainAddress())

	}
	return bc
}

func (bcs *BlockchainServer) GetChain(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
	bc := bcs.GetBlockchain()
		m , _ := bc.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		log.Printf("ERROR: Invalid HTTP Method")
	}
}
func (bcs *BlockchainServer) Run() {
	http.HandleFunc("/",  bcs.GetChain )
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcs.Port())), nil ))

}
func main() {
    port := flag.Uint("port", 6000, "TCP Port Number for Blockchain Server")
    flag.Parse()

    app := NewBlockchainServer(uint16(*port)) // Ensure this function is defined in blockchain_server.go
    app.Run()
}
