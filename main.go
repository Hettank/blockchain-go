package main

import (
	"fmt"
	"time"

	"github.com/Hettank/blockchain-go/blockchain"
)

func main() {
	start := time.Now()

	fmt.Println("========================================")
	fmt.Println("CREATING BLOCKCHAIN")
	fmt.Println("========================================")

	// Create blockchain with mempool
	bc := blockchain.NewBlockchain()
	blockchain.Difficulty = 2

	fmt.Println("\nCREATING WALLETS")
	fmt.Println("----------------------------------------")

	// Create wallets
	alice := blockchain.NewWallet()
	bob := blockchain.NewWallet()
	charlie := blockchain.NewWallet()
	dave := blockchain.NewWallet()
	eve := blockchain.NewWallet()

	fmt.Printf("Alice:   %s\n", alice.Address)
	fmt.Printf("Bob:     %s\n", bob.Address)
	fmt.Printf("Charlie: %s\n", charlie.Address)
	fmt.Printf("Dave:    %s\n", dave.Address)
	fmt.Printf("Eve:     %s\n", eve.Address)

	fmt.Println("\nCREATING & SIGNING TRANSACTIONS")
	fmt.Println("----------------------------------------")

	// Create transactions
	transactions := []struct {
		Sender   *blockchain.Wallet
		Receiver *blockchain.Wallet
		Amount   int
		Desc     string
	}{
		{alice, bob, 10, "Alice → Bob"},
		{bob, charlie, 5, "Bob → Charlie"},
		{charlie, dave, 3, "Charlie → Dave"},
		{dave, eve, 2, "Dave → Eve"},
		{eve, alice, 1, "Eve → Alice"},
		{alice, charlie, 4, "Alice → Charlie"},
		{bob, dave, 6, "Bob → Dave"},
	}

	var signedTransactions []blockchain.Transaction

	for i, txData := range transactions {
		fmt.Printf("\n%d. %s: %d coins\n", i+1, txData.Desc, txData.Amount)

		// Create transaction
		tx := blockchain.Transaction{
			From:   txData.Sender.Address,
			To:     txData.Receiver.Address,
			Amount: txData.Amount,
		}

		// Sign the transaction
		txData.Sender.Sign(&tx)
		fmt.Printf("Signed: %x...\n", tx.Signature[:8])

		// Verify the transaction
		if tx.Verify() {
			fmt.Printf("Verified successfully!\n")
			signedTransactions = append(signedTransactions, tx)
		} else {
			fmt.Printf("Verification FAILED! Transaction rejected.\n")
		}
	}

	fmt.Println("\nADDING TRANSACTIONS TO MEMPOOL")
	fmt.Println("----------------------------------------")

	// Add verified transactions to mempool
	addedCount := 0
	for _, tx := range signedTransactions {
		if bc.Mempool.AddTransaction(tx) {
			addedCount++
		}
	}

	fmt.Printf("Added %d/%d transactions to mempool\n", addedCount, len(signedTransactions))
	fmt.Printf("Mempool: %d/%d transactions\n", bc.Mempool.Count(), bc.Mempool.MaxSize)

	fmt.Println("\nMINING BLOCKS FROM MEMPOOL")
	fmt.Println("----------------------------------------")

	// Mine transactions from mempool
	bc.Mempool.MinePendingTransactions(bc)

	elapsed := time.Since(start)
	fmt.Printf("\nTotal time: %s\n", elapsed)

	// Print blockchain
	fmt.Println("\nBLOCKCHAIN:")
	bc.Print()

	// Validate blockchain
	fmt.Println("\n🔍 VALIDATION:")
	if bc.Validate() {
		fmt.Println("Blockchain is valid!")
	} else {
		fmt.Println("Blockchain is INVALID!")
	}

	// Test tampering
	fmt.Println("\nTAMPER TEST:")
	if len(bc.Blocks) > 1 {
		// Find a transaction to tamper with
		tampered := false
		for i := 1; i < len(bc.Blocks); i++ {
			if len(bc.Blocks[i].Transactions) > 0 {
				fmt.Printf("Before: Block %d, Transaction 1 Amount = %d\n",
					i, bc.Blocks[i].Transactions[0].Amount)

				bc.Blocks[i].Transactions[0].Amount = 999 // Tamper

				fmt.Printf("After:  Block %d, Transaction 1 Amount = %d\n",
					i, bc.Blocks[i].Transactions[0].Amount)
				tampered = true
				break
			}
		}

		if tampered {
			fmt.Println("\nVALIDATING AFTER TAMPERING:")
			if bc.Validate() {
				fmt.Println("Blockchain is valid!")
			} else {
				fmt.Println("Blockchain is INVALID! Tampering detected!")
			}
		} else {
			fmt.Println("No transactions found to tamper with")
		}
	} else {
		fmt.Println("Not enough blocks to test tampering")
	}

	fmt.Println("\n========================================")
	fmt.Println("DONE")
	fmt.Println("========================================")
}
