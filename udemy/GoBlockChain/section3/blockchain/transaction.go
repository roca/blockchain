package blockchain

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_blockchain_address\t%s\n", t.senderBlockchainAddress)
	fmt.Printf(" recipient_blockchain_address\t%s\n", t.recipientBlockchainAddress)
	fmt.Printf(" value\t%.1f\n", t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderBlockchainAddress    string  `json:"sender_blockchain_address"`
		RecipientBlockchainAddress string  `json:"recipient_blockchain_address"`
		Value                      float32 `json:"value"`
	}{
		SenderBlockchainAddress:    t.senderBlockchainAddress,
		RecipientBlockchainAddress: t.recipientBlockchainAddress,
		Value:                      t.value,
	})
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

type TransactionRequest struct {
	SenderBlockchainAddress *string `json:"sender_blockchain_address"`
	RecipientBlockchainAddress *string `json:"recipient_blockchain_address"`
	SenderPublicKey *string `json:"sender_public_key"`
	Value *float32 `json:"value"`
	Signature *string `json:"signature"`
}

func (tr *TransactionRequest) Validate() bool{
	if tr.SenderBlockchainAddress == nil ||
	tr.RecipientBlockchainAddress == nil ||
	tr.SenderPublicKey == nil ||
	tr.Value == nil ||
	tr.Signature == nil {
		return false
	}
	return true
}
