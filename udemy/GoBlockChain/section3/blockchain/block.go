package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []*Transaction
}

func (b *Block) Hash() [32]byte {
	m, e := json.Marshal(b)
	if e != nil {
		panic(e)
	}
	return sha256.Sum256(m)
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		Nonce        int            `json:"nonce"`
		PreviousHash string      `json:"previous_hash"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: fmt.Sprintf("%x",b.previousHash),
		Transactions: b.transactions,
	})
}

func (b *Block) Print() {
	fmt.Printf("Timestamp\t%d\n", b.timestamp)
	fmt.Printf("Nonce\t\t%d\n", b.nonce)
	fmt.Printf("PreviousHash\t%x\n", b.previousHash)
	for _, t := range b.transactions {
		t.Print()
	}
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	b.transactions = transactions
	return b
}
