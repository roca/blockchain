package blockchain

import (
	"encoding/json"
	"time"
)

type Block struct {
	PrevHash     string         `json:"prev_hash"`
	Timestamp    int64          `json:"timestamp"`
	Nonce        int            `json:"nonce"`
	Transactions []*Transaction `json:"transactions"`
}

func NewBlock(prevHash string, nonce int) *Block {
	block := &Block{
		PrevHash:     prevHash,
		Timestamp:    time.Now().Unix(),
		Nonce:        nonce,
		Transactions: []*Transaction{},
	}

	return block
}

func (b Block) String() string {
	nb, err := json.Marshal(b)
	if err != nil {
		return err.Error()
	}

	return string(nb)
}
