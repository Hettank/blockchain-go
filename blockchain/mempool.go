package blockchain

import (
	"fmt"
)

type Mempool struct {
	Transactions []Transaction
	MaxSize      int
}

func NewMempool(maxSize int) *Mempool {
	return &Mempool{
		Transactions: []Transaction{},
		MaxSize:      maxSize,
	}
}

func (m *Mempool) AddTransaction(tx Transaction) bool {
	if len(m.Transactions) >= m.MaxSize {
		fmt.Printf("Mempool is FULL! (Limit: %d)\n", m.MaxSize)
		fmt.Printf("   Transaction %s→%s rejected\n", tx.From, tx.To)
		return false
	}

	if tx.Amount <= 0 {
		fmt.Println("Invalid transaction: amount must be positive")
		return false
	}

	if tx.From == tx.To {
		fmt.Println("Invalid transaction: sender and receiver same")
		return false
	}

	// Add to mempool
	m.Transactions = append(m.Transactions, tx)
	return true
}

func (m *Mempool) AddTransactions(txs []Transaction) int {
	added := 0

	for _, tx := range txs {
		if m.AddTransaction(tx) {
			added++
		}
	}

	if added > 0 {
		fmt.Printf("Added %d transactions. Mempool: %d/%d\n",
			added, len(m.Transactions), m.MaxSize)
	}

	return added
}

func (m *Mempool) MinePendingTransactions(bc *Blockchain) {
	bunch := 5
	transactions := []Transaction{}

	if len(m.Transactions) < bunch {
		bunch = len(m.Transactions)
	}

	fmt.Printf("Mining block with %d transactions...\n\n", len(transactions))

	for i := 0; i < len(m.Transactions); i++ {
		for j := 0; j < bunch; j++ {
			transactions = m.Transactions[:bunch]

			m.Transactions = m.Transactions[bunch:]

			// Get previous block hash
			latestBlock := bc.GetLatestBlock().Hash

			// Create a new block
			newBlock := NewBlock(transactions, latestBlock)

			// Mine the block
			newBlock.MineBlock(Difficulty)

			// append block in a blockchain
			bc.Blocks = append(bc.Blocks, newBlock)
		}
	}

	fmt.Printf("Block mined! Total blocks: %d, Mempool remaining: %d\n",
		len(bc.Blocks), len(m.Transactions))
}
