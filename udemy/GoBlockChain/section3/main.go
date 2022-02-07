package main

import (
	"log"

	"udemy.com/goblockchain/section3/blockchain"
)

func init() {
	log.SetPrefix("Blockchain")
}

func main() {
	MinerBlockchainAddress := "miner-blockchain-address"
	blockChain := blockchain.NewBlockchain(MinerBlockchainAddress)
	blockChain.Print()

	blockChain.AddTransaction("A", "B", 1.0)
	blockChain.Mining()
	blockChain.Print()

	blockChain.AddTransaction("C", "D", 2.0)
	blockChain.AddTransaction("X", "Y", 3.0)
	blockChain.Mining()
	blockChain.Print()
}
