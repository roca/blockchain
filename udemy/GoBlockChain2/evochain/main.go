package main

import (
	"evochain/blockchain"
	"evochain/constants"
	"log"
	"sync"
)

func init() {
	log.SetPrefix(constants.BLOCKCHAIN_NAME + ": ")
}

func main() {

	var wg sync.WaitGroup

	genesisBlock := blockchain.NewBlock("0x0", 0)
	blockchain := blockchain.NewBlockchain(genesisBlock)
	log.Println(blockchain)
	log.Print("Starting the mining process..", "\n\n")
	wg.Add(1)
	go blockchain.ProofOfWorkMining("alice")
	wg.Wait()
}
