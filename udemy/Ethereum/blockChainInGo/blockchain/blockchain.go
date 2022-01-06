package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	// This will join our previous block hash with current block data

	hash := sha256.Sum256(info)
	//The actual hashing algorithm

	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	// block := &Block{[]byte{}, []byte(data), prevHash}
	// //It is simple subtituing value to block
	// block.DeriveHash()
	// return block

	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	// Don't forget to add the 0 at the end for the nonce!
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

type BlockChain struct {
	Blocks []*Block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
