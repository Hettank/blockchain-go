package blockchain

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

type Blockchain struct {
	Blocks  []*Block
	Mempool *Mempool
	Ledger  map[string]int
}

var Difficulty = 1

func NewBlockchain() *Blockchain {
	genesisBlock := NewBlock([]Transaction{}, []byte{})

	genesisBlock.MineBlock(Difficulty)

	return &Blockchain{
		Blocks:  []*Block{genesisBlock},
		Mempool: NewMempool(50),
		Ledger:  make(map[string]int),
	}
}

func (bc *Blockchain) GetLatestBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *Blockchain) GetBalance(address string) int {
	return bc.Ledger[address]
}

func (bc *Blockchain) ApplyTransaction(tx Transaction) {
	bc.Ledger[tx.From] -= tx.Amount
	bc.Ledger[tx.To] += tx.Amount
}

func (bc *Blockchain) HasSufficientBalance(tx Transaction) bool {
	return bc.GetBalance(tx.From) >= tx.Amount
}

func (bc *Blockchain) Print() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("  BLOCKCHAIN\n")
	fmt.Printf("  Difficulty: %d  |  Total Blocks: %d\n", Difficulty, len(bc.Blocks))
	fmt.Println(strings.Repeat("=", 60))

	for i, block := range bc.Blocks {
		fmt.Printf("\n  ┌── Block #%d ──\n", i)
		fmt.Printf("  │  Timestamp: %s\n", time.Unix(block.Timestamp, 0).Format("15:04:05"))

		if len(block.Transactions) == 0 {
			fmt.Printf("  │  Transactions: None (Genesis)\n")
		} else {
			fmt.Printf("  │  Transactions:\n")
			for j, tx := range block.Transactions {
				fmt.Printf("  │    %d. %s → %s: %d coins\n",
					j+1, tx.From, tx.To, tx.Amount)
			}
		}

		fmt.Printf("  │  Nonce: %d\n", block.Nonce)
		fmt.Printf("  │  Hash: %x\n", block.Hash[:12])
		if i > 0 {
			fmt.Printf("  │  Prev: %x\n", block.PrevHash[:12])
		}
		fmt.Printf("  └──\n")
	}
	fmt.Println(strings.Repeat("=", 60))
}

func (bc *Blockchain) Validate() bool {
	for i := 0; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]

		storedHash := currentBlock.Hash
		recalculatedHash := currentBlock.GenerateHash()

		if !bytes.Equal(storedHash, recalculatedHash) {
			return false
		}
	}

	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		previousBlock := bc.Blocks[i-1]

		if string(currentBlock.PrevHash) != string(previousBlock.Hash) {
			return false
		}
	}
	return true
}
