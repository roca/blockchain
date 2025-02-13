package blockchain

import (
	"encoding/json"
	"evochain/constants"
	"strings"
)

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

func (bc *Blockchain) String() string {
	nb, err := json.Marshal(bc)
	if err != nil {
		return err.Error()
	}

	return string(nb)
}

func (bc *Blockchain) AddBlock(b *Block) {
	m := map[string]bool{}
	for _, txn := range b.Transactions {
		m[txn.TransactionHash] = true
	}

	// remove the transactions from the pool``
	for idx,txn := range bc.TransactionPool {
		_, ok := m[txn.TransactionHash]
		if ok {
			bc.TransactionPool = append(bc.TransactionPool[:idx], bc.TransactionPool[idx+1:]...)
		}
	}

	bc.Blocks = append(bc.Blocks, b)
}

func (bc *Blockchain) AddTransaction(t *Transaction) {
	bc.TransactionPool = append(bc.TransactionPool, t)
}

func (bc *Blockchain) ProofOfWorkMining(minersAddress string) {
	// calulate the PrevHash
	prevHash := bc.Blocks[len(bc.Blocks)-1].Hash()
	// start with a 0 nonce
	nonce := 0
	for {
		// create a new block
		guessBlock := NewBlock(prevHash, nonce)
		// copy the transaction pool
		for _, txn := range bc.TransactionPool {
			newTxn := NewTransaction(txn.From, txn.To, txn.Value, txn.Data)
			guessBlock.AddTransaction(newTxn)
		}
		// guess the Hash
		guessHash := guessBlock.Hash()
		desiredHash := strings.Repeat("0", constants.MINING_DIFFICULTY)
		ourSolutionHash := guessHash[2 : 2+constants.MINING_DIFFICULTY]
		// compare this hash with the mining difficulty
		if ourSolutionHash == desiredHash {
			rewardTxn := NewTransaction(constants.BLOCKCHAIN_ADDRESS, minersAddress, constants.MINING_REWARD, []byte{})
			guessBlock.Transactions = append(guessBlock.Transactions, rewardTxn)
			bc.AddBlock(guessBlock)
			prevHash = bc.Blocks[len(bc.Blocks)-1].Hash()
			nonce = 0
			continue
		}
		nonce++
	}

}
