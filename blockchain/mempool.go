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
	// Check size limit
	if len(m.Transactions) >= m.MaxSize {
		fmt.Printf("Mempool is FULL! (Limit: %d)\n", m.MaxSize)
		fmt.Printf("   Transaction %s→%s rejected\n", tx.From, tx.To)
		return false
	}

	// Check amount
	if tx.Amount <= 0 {
		fmt.Println("Invalid transaction: amount must be positive")
		return false
	}

	// Check sender != receiver
	if tx.From == tx.To {
		fmt.Println("Invalid transaction: sender and receiver same")
		return false
	}

	// NEW: Verify transaction signature
	if !tx.Verify() {
		fmt.Printf("Invalid transaction: signature verification failed!\n")
		fmt.Printf("   From: %s, To: %s, Amount: %d\n", tx.From, tx.To, tx.Amount)
		return false
	}

	// Add to mempool
	m.Transactions = append(m.Transactions, tx)
	fmt.Printf("Transaction added to mempool: %s → %s: %d coins\n",
		tx.From, tx.To, tx.Amount)
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

	fmt.Printf("Mining %d transactions from mempool...\n", len(m.Transactions))

	blocksMined := 0
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
		blocksMined++

		fmt.Printf("Mined block #%d with %d transactions (Remaining: %d)\n",
			blocksMined, count, len(m.Transactions))
	}

	fmt.Printf("Mining complete! Mined %d blocks.\n", blocksMined)
}

func (m *Mempool) Count() int {
	return len(m.Transactions)
}
