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

	fmt.Println("\n**Creating Transactions...")
	txs := []blockchain.Transaction{
		{
			From:   "Alice",
			To:     "Bob",
			Amount: 5,
		},
		{
			From:   "Bob",
			To:     "Charlie",
			Amount: 2,
		},
		{
			From:   "Charlie",
			To:     "Het",
			Amount: 1,
		},
	}

	fmt.Println("**Mining Block...")
	bc.AddBlock(txs)

	elapsed := time.Since(start)
	fmt.Printf("\n**Mining completed in: %s\n\n", elapsed)

	// Print blockchain
	fmt.Println("**BLOCKCHAIN:")
	bc.Print()

	// Validate
	fmt.Println("\n**VALIDATION:")
	if bc.Validate() {
		fmt.Println("Blockchain is valid!")
	} else {
		fmt.Println("Blockchain is INVALID!")
	}

	// Test tampering
	fmt.Println("\n**TAMPER TEST:")
	fmt.Printf("Before: Block 1, Transaction 1 Amount = %d\n", bc.Blocks[1].Transactions[0].Amount)
	bc.Blocks[1].Transactions[0].Amount = 100 // Tamper!
	fmt.Printf("After:  Block 1, Transaction 1 Amount = %d\n", bc.Blocks[1].Transactions[0].Amount)

	fmt.Println("\n**VALIDATING AFTER TAMPERING:")
	if bc.Validate() {
		fmt.Println("Blockchain is valid!")
	} else {
		fmt.Println("Blockchain is INVALID! Tampering detected!")
	}
}
