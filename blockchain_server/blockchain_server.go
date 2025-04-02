 package main

import (
	"blockchain/block"
	"blockchain/utils"
	"blockchain/wallet"
	"encoding/json"

	"io"
	"log"
	"net/http"
	"strconv"
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
func(bcs *BlockchainServer) Transactions(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet :
      w.Header().Add("Conten-Type", "application/json")
	  bc := bcs.GetBlockchain()
	  transactions := bc.TransactionPool()
	  m , _ := json.Marshal(struct {
		Transactions []*block.Transaction `json:"transactions"`
		Length int `json:"Length"`
	  }{
		Transactions : transactions,
		Length: len(transactions), 
	  })
	  io.WriteString(w, string(m))
	case http.MethodPost :
        decoder := json.NewDecoder(req.Body)
		var t block.TransactionRequest
		err := decoder.Decode(&t)
		if err != nil {
			log.Printf("ERROR: %v", err)
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return 
		}
		if !t.Validate() {
			log.Println("ERROR : missing fields(s)")
		}
		publicKey := utils.PublicKeyFromString(*t.SenderPublicKey)
		signature := utils.SignatureFromString(*t.Signature)
		bc := bcs.GetBlockchain()
		newTransaction := block.NewTransaction(*t.SenderBlockchainAddress, *t.RecipientBlockchainAddress, *t.Value)
		isCreated := bc.CreateTransaction(*t.SenderBlockchainAddress, *t.RecipientBlockchainAddress, *t.Value, publicKey, signature, newTransaction)
		


		w.Header().Add("Content-Type", "application/json")
		var m []byte
		if !isCreated {
			w.WriteHeader(http.StatusBadRequest)
			m = utils.JsonStatus("fail")
		} else {
			w.WriteHeader(http.StatusCreated)
			m = utils.JsonStatus("success")
		}
		io.WriteString(w, string(m))
	default:
		log.Println("ERROR:Invalid HTTP Method")
		w.WriteHeader(http.StatusBadRequest)
	}
}
func (bcs *BlockchainServer) Transaction(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		// GET request handle karna baaki hai
		w.WriteHeader(http.StatusNotImplemented)
		io.WriteString(w, string(utils.JsonStatus("fail")))
		return

	case http.MethodPost:
		decoder := json.NewDecoder(req.Body)
		var t block.TransactionRequest

		err := decoder.Decode(&t)
		if err != nil {
			log.Printf("ERROR: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}

		// Validate transaction request
		if !t.Validate() {
			log.Println("ERROR: missing field(s)")
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}

		// Check for nil values before parsing
		if t.SenderPublicKey == nil || *t.SenderPublicKey == "" {
			log.Println("ERROR: Sender Public Key is missing")
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}

		if t.Signature == nil || *t.Signature == "" {
			log.Println("ERROR: Signature is missing")
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}

		publicKey := utils.PublicKeyFromString(*t.SenderPublicKey)
		signature := utils.SignatureFromString(*t.Signature)

		if publicKey == nil {
			log.Println("ERROR: Invalid Sender Public Key")
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}

		if signature == nil {
			log.Println("ERROR: Invalid Signature")
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}

		// Blockchain object le
		bc := bcs.GetBlockchain()

		// Transaction banakar blockchain me daal
		isCreated := bc.CreateTransaction(
			*t.SenderBlockchainAddress,
			*t.RecipientBlockchainAddress,
			*t.Value,
			publicKey,
			signature,
			block.NewTransaction(*t.SenderBlockchainAddress, *t.RecipientBlockchainAddress, *t.Value),
		)

		w.Header().Add("Content-Type", "application/json")
		var m []byte
		if !isCreated {
			w.WriteHeader(http.StatusBadRequest)
			m = utils.JsonStatus("fail")
		} else {
			w.WriteHeader(http.StatusCreated)
			m = utils.JsonStatus("success")
		}
		io.WriteString(w, string(m))

	default:
		log.Println("ERROR: Invalid HTTP Method")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, string(utils.JsonStatus("fail")))
	}
}

func (bcs *BlockchainServer) Mine(w http.ResponseWriter, req *http.Request)  {
	switch req.Method {
	case http.MethodGet:
		bc := bcs.GetBlockchain()
		isMined := bc.Mining()


		var m []byte
		if !isMined {
			w.WriteHeader(http.StatusBadRequest)
			m = utils.JsonStatus("fail")
		} else {
			m = utils.JsonStatus("success")
		}
		w.Header().Add("content_type", "application/json")
		io.WriteString(w, string(m))
	default:
		log.Println("error: invalid http method")
		w.WriteHeader(http.StatusBadRequest)
	}
}
func (bcs *BlockchainServer) StartMine(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		bc := bcs.GetBlockchain()
		bc.StartMining() 

		m := utils.JsonStatus("success")
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, string(m))

	default:
		log.Println("ERROR: Invalid HTTP Method")
		w.WriteHeader(http.StatusBadRequest)
	}
}





func (bcs *BlockchainServer) Run() {

	http.HandleFunc("/",  bcs.GetChain )
	http.HandleFunc("/transaction", bcs.Transactions)
	http.HandleFunc("/mine", bcs.Mine)
	http.HandleFunc("/Mine/Start", bcs.StartMine)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcs.Port())), nil ))

}

