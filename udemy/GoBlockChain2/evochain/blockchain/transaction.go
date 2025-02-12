package blockchain

import (
	"encoding/json"
	"math"
)

type Transaction struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value uint64  `json:"value"`
	Data  []byte `json:"data"`
	Status string `json:"status"`
}

func NewTransaction(from, to string, value uint64, data []byte) *Transaction {
	t := &Transaction{
		From:  from,
		To:    to,
		Value: value,
		Data:  data,
	}
	
	return t
}

func (t Transaction) String() string {
	nb, err := json.Marshal(t)
	if err != nil {
		return err.Error()
	}

	return string(nb)
}

func (t *Transaction) IsValid() bool {
	if t.Value == 0 {
		return false
	}

	if t.Value > math.MaxUint64 {
		return false
	}

	// TODO: check the signatures
	
	return true
}
