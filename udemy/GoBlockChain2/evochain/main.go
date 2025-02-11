package main

import (
	"evochain/blockchain"
	"evochain/constants"
	"log"
)

func init() {
	log.SetPrefix(constants.BLOCKCHAIN_NAME + ": ")
}

func main() {
	block := blockchain.NewBlock("0x", 1)
	log.Println(block)

	transaction := blockchain.NewTransaction("0x1", "0x2", 12, []byte{})
	log.Println(transaction)

	genesisBlock := blockchain.NewBlock("0x", 0)
	blockchain := blockchain.NewBlockchain(genesisBlock)
	log.Println(blockchain)
}
