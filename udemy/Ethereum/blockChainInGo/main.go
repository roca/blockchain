package main

import (
	"fmt"

	b "udemy.com/Ethereum/blockChainInGo/blockchain"
)

func main() {
	// Create a new Block
	newBlock := b.Block{[]byte("Hey!"), []byte("Hey!"), []byte("Hey!")}

	// Print the block
	fmt.Println(newBlock)
}
