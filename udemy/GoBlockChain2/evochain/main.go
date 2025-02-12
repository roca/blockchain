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
	log.Println("Hash of the Block",block.Hash())

	transaction := blockchain.NewTransaction("0x1", "0x2", 12, []byte{})
	log.Println(transaction)

	genesisBlock := blockchain.NewBlock("0x0", 0)
	firstTransaction := blockchain.NewTransaction("0x0", "0x1", 100_000, []byte{})
	genesisBlock.Transactions = append(genesisBlock.Transactions, firstTransaction)
	blockchain := blockchain.NewBlockchain(genesisBlock)
	log.Println(blockchain)
}
