package blockchain

import (
	"encoding/json"
	"fmt"
	"os"
)

type TransactionWrapper struct {
	Transactions []Transaction `json:"transactions"`
}

// Expects JSON format: [ { "from": "...", "to": "...", "amount": 0 } ]
func LoadTransactionsFromFileSimple(filename string) ([]Transaction, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	var transactions []Transaction
	err = json.Unmarshal(data, &transactions)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	fmt.Printf("Loaded %d transactions from %s\n", len(transactions), filename)
	return transactions, nil
}
