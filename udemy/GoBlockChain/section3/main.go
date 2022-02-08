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

	// Wallet A sends some coins to Wallet B
	t := wallet.NewTransaction(
		walletA.PrivateKey(),
		walletA.PublicKey(),
		walletA.BlockchainAddress(),
		walletB.BlockchainAddress(),
		1.0)

	// Blockchain
	blockchain := blockchain.NewBlockchain(walletM.BlockchainAddress())
	isAdded := blockchain.AddTransaction(
		walletA.BlockchainAddress(),
		walletB.BlockchainAddress(),
		1.0,
		walletA.PublicKey(),
		t.GenerateSignature())

	fmt.Println("Added?", isAdded)

	blockchain.Mining()
	blockchain.Print()

	fmt.Printf("A %.1f\n", blockchain.CalculateTotalAmount(walletA.BlockchainAddress()))
	fmt.Printf("B %.1f\n", blockchain.CalculateTotalAmount(walletB.BlockchainAddress()))
	fmt.Printf("M %.1f\n", blockchain.CalculateTotalAmount(walletM.BlockchainAddress()))
}
