package main

import (
	"evochain/blockchain"
	"evochain/constants"
	"log"
	"sync"
	"time"
)

func init() {
	log.SetPrefix(constants.BLOCKCHAIN_NAME + ": ")
}

func main() {

	var wg sync.WaitGroup

	genesisBlock := blockchain.NewBlock("0x0", 0)
	firstTransaction := blockchain.NewTransaction("0x0", "0x1", 100_000, []byte{})
	blockchain := blockchain.NewBlockchain(genesisBlock)
	log.Println(blockchain)

	wg.Add(1)
	go blockchain.ProofOfWorkMining("alice")
	time.Sleep(2000)
	blockchain.AddTransaction(firstTransaction)
	log.Println("Transaction Pool Log", blockchain, "\n\n")
	wg.Wait()
}
