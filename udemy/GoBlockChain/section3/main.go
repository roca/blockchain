package main

import (
	"fmt"
	"log"

	"udemy.com/goblockchain/section3/blockchain"
	"udemy.com/goblockchain/section3/wallet"
)

func init() {
	log.SetPrefix("Blockchain")
}

func main() {
	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	t := blockchain.NewTransaction(
		walletA.PrivateKey(),
		walletA.PublicKey(),
		walletA.BlockchainAddress(),
		walletB.BlockchainAddress(),
		1.0,
	)

	blockchain := blockchain.NewBlockchain(walletM.BlockchainAddress())
	isAdded := blockchain.AddTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0, t.GenerateSignature())
	fmt.Println("Added?", isAdded)
}
