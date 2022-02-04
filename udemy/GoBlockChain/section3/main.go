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
	blockChain.CreateBlock(5, "hash 1")
	blockChain.Print()
	blockChain.CreateBlock(2, "hash 2")
	blockChain.Print()
}
