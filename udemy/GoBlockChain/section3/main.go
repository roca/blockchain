package main

import (
	"log"

	"udemy.com/goblockchain/section3/blockchain"
)

func init() {
	log.SetPrefix("Blockchain")
}

func main() {
	blockChain := blockchain.NewBlockchain()
	blockChain.Print()

	blockChain.AddTransaction("A", "B", 1.0)

	previousHash := blockChain.LastBlock().Hash()
	nonce := blockChain.ProofOfWork()
	blockChain.CreateBlock(nonce, previousHash)
	blockChain.Print()

	blockChain.AddTransaction("C", "D", 2.0)
	blockChain.AddTransaction("X", "Y", 3.0)

	previousHash = blockChain.LastBlock().Hash()
	nonce = blockChain.ProofOfWork()
	blockChain.CreateBlock(nonce, previousHash)
	blockChain.Print()
}
