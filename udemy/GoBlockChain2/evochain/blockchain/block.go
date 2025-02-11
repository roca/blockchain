package blockchain

import (
	"encoding/json"
	"time"
)

type Block struct {
	PrevHash     string        `json:"prevHash"`
	Timestamp    int64         `json:"timestamp"`
	Nonce        int           `json:"nonce"`
	Transactions []*Transaction `json:"transactions"`
}

func NewBlock(prevHash string, nonce int) *Block {
	block := new(Block)
	block.PrevHash = prevHash
	block.Timestamp = time.Now().Unix()
	block.Nonce = nonce
	block.Transactions = []*Transaction{}
	return block
}

func (b Block) ToJson() string {
	nb, err := json.Marshal(b)
	if err != nil {
		return err.Error()
	}

	return string(nb)
}
