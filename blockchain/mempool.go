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
	const MaxTransactionsPerBlock = 5

	// Check if there are transactions
	if len(m.Transactions) == 0 {
		fmt.Println("No transactions in mempool to mine")
		return
	}

	for len(m.Transactions) > 0 {
		count := MaxTransactionsPerBlock
		if len(m.Transactions) < count {
			count = len(m.Transactions)
		}

		// Take transactions from the front
		transactions := m.Transactions[:count]

		// Remove them from mempool
		m.Transactions = m.Transactions[count:]

		// Get previous block hash
		latestBlock := bc.GetLatestBlock()

		// Create a new block
		newBlock := NewBlock(transactions, latestBlock.Hash)

		// Mine the block
		newBlock.MineBlock(Difficulty)

		// Append block to blockchain
		bc.Blocks = append(bc.Blocks, newBlock)
	}
}
