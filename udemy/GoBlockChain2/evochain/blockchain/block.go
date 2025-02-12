package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"evochain/constants"
	"time"
)

type Block struct {
	PrevHash           string         `json:"prev_hash"`
	Timestamp          int64          `json:"timestamp"`
	Nonce              int            `json:"nonce"`
	Transactions       []*Transaction `json:"transactions"`
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

func (b *Block) Hash() string {

	bs := []byte(b.String())
	sum := sha256.Sum256(bs)
	hexRep := hex.EncodeToString(sum[:32])
	formattedHexRep := constants.HEX_PREFIX + hexRep

	return formattedHexRep
}

func (b *Block) AddTransaction(t *Transaction) {
	// verify the transaction
	t.Status = constants.FAILED
	if t.IsValid() {
		t.Status = constants.SUCCESS
	}
	b.Transactions = append(b.Transactions, t)
}
