package blockchain

import "encoding/json"

type Transaction struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value uint64  `json:"value"`
	Data  []byte `json:"data"`
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
