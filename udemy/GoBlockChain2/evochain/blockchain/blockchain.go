package blockchain

import "encoding/json"

type Blockchain struct {
	TransactionPool []*Transaction `json:"transaction_pool"`
	Blocks          []*Block       `json:"block_chain"`
}

func NewBlockchain(genesisBlock *Block) *Blockchain {
	bc := &Blockchain{
		TransactionPool: []*Transaction{},
		Blocks:          []*Block{genesisBlock},
	}

	return bc
}

func (bc Blockchain) String() string {
	nb, err := json.Marshal(bc)
	if err != nil {
		return err.Error()
	}

	return string(nb)
}
