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
}
