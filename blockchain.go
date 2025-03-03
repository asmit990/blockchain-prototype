package main
import (
	"fmt"
	"log"
	"time"
	"strings"
	"crypto/sha256"
	"encoding/json"
)
type Block struct {
	nonce int
	previousHash string
	timestamp int64
	transactions []string
}

func NewBlock(nonce int , previousHash string) *Block{
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b 
	
}

func (b *Block) Print(){
	fmt.Printf("timestamp   %d\n", b.timestamp)
	fmt.Printf("nonce   %d\n", b.nonce)
	fmt.Printf("previous_hash  %s\n", b.previousHash)
	fmt.Printf("transactions  %v\n", b.transactions)

}
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64 `json:"timestamp"`
		Nonce        int  `json:"nonce"`
		PreviousHash string`json:"previous_hash"`
		Transactions []string`json:"transactions"`
	}{
		Timestamp:    b.timestamp,  
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}
func NewBlockChain() *Blockchain{
	bc := new(Blockchain)

	bc.CreateBlock(0, "Init hash")
	return bc
}

func(bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}
func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain) -1 ]
}
func (bc *Blockchain) Print(){
	for i, block := range bc.chain {
		fmt.Printf("%s chain %d  %s\n",  strings.Repeat("=", 25), i,
		strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("=",25))
}

type Transactions struxt {
	senderBlockchainAddress string
	recipientBlockAddress   string
	value                   float32

}
func NewTransaction(sender string, recipient string, value float32) *Transaction{
  return &Transation{sender, recipient, value}	
}
func (t *Transaction) Print() {
	fmt.Printf("%s\n", string.Repeat("-", 40) )
	fmt.Printf("sender_blockchain_address      %s\n", t.senderBlockchainAddress)
	fmt.Printf("recipient_blockchain_address   %s\n", t.recipientBlockAddress)
    fmt.Printf("value                          %.1f\n", t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error)
func init(){
	log.SetPrefix("BlockChain : ")
}

func main() {  
 blockChain := NewBlockChain()
blockChain.Print()
blockChain.CreateBlock(5, "hash 1")
blockChain.Print()
blockChain.CreateBlock(5, "hash 2")
blockChain.Print()

}

