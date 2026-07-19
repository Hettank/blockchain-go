package main

import (
	"fmt"
	"time"

	"github.com/Hettank/blockchain-go/blockchain"
)

func main() {
	start := time.Now()

	fmt.Println("**Creating Blockchain...")
	bc := blockchain.NewBlockchain()

	// Create mempool with size 50
	mp := blockchain.NewMempool(50)

	// Set difficulty
	blockchain.Difficulty = 2

	fmt.Println("\nLoading Transactions form file...")
	txs, err := blockchain.LoadTransactionsFromFileSimple("transactions.json")

	if err != nil {
		fmt.Printf("Error loading transactions: %v\n", err)
		fmt.Println("Using fallback transactions...")

		// Fallback transactions
		txs = []blockchain.Transaction{
			{From: "Alice", To: "Bob", Amount: 5},
			{From: "Bob", To: "Charlie", Amount: 3},
			{From: "Charlie", To: "Dave", Amount: 2},
		}
	}

	// Add transactions to mempool
	fmt.Println("\nAdding transactions to mempool...")
	mp.AddTransactions(txs)

	// Mine transactions from mempool
	fmt.Println("\nMining block from mempool...")
	mp.MinePendingTransactions(bc)

	elapsed := time.Since(start)
	fmt.Printf("\nMining completed in: %s\n\n", elapsed)

	// Print blockchain
	fmt.Println("BLOCKCHAIN:")
	bc.Print()

	// Validate blockchain
	fmt.Println("\nVALIDATION:")
	if bc.Validate() {
		fmt.Println("Blockchain is valid!")
	} else {
		fmt.Println("Blockchain is INVALID!")
	}

	// Test tampering
	fmt.Println("\nTAMPER TEST:")
	if len(bc.Blocks) > 1 {
		fmt.Printf("Before: Block 1, Transaction 1 Amount = %d\n", bc.Blocks[2].Transactions[3].Amount)

		bc.Blocks[2].Transactions[3].Amount = 100 // Tamper

		fmt.Printf("After:  Block 1, Transaction 1 Amount = %d\n", bc.Blocks[2].Transactions[3].Amount)

		fmt.Println("\nVALIDATING AFTER TAMPERING:")
		if bc.Validate() {
			fmt.Println("Blockchain is valid!")
		} else {
			fmt.Println("Blockchain is INVALID! Tampering detected!")
		}
	} else {
		fmt.Println("Not enough blocks to test tampering")
	}
}
