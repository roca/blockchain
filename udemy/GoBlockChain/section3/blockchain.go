package main

import (
	"fmt"
	"log"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func (b Block) Print(){
    fmt.Printf("Timestamp\t%d\n", b.timestamp)
    fmt.Printf("Nonce\t\t%d\n", b.nonce)
    fmt.Printf("PreviousHash\t%s\n", b.previousHash)
    fmt.Printf("Transactions\t%s\n", b.transactions)
}

func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}

func init() {
	log.SetPrefix("Blockchain")
}

func main() {
	b := NewBlock(0, "init hash")
	b.Print()
}
