package blockchain

import (
	"bytes"
	"fmt"
)

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	genesisBlock := NewBlock("Genesis Block", []byte{})

	return &Blockchain{
		Blocks: []*Block{genesisBlock},
	}
}

func (bc *Blockchain) GetLatestBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *Blockchain) AddBlock(data string) {
	latestBlock := bc.GetLatestBlock()

	newBlock := NewBlock(data, latestBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) Print() {
	fmt.Println("=========================")

	for i, block := range bc.Blocks {
		fmt.Printf("\n Block %d:\n", i)
		fmt.Printf("   Timestamp: %d\n", block.Timestamp)
		fmt.Printf("   Data: %s\n", block.Data)
		fmt.Printf("   Hash: %x\n", block.Hash)

		if i > 0 {
			fmt.Printf("   Previous Hash: %x\n", block.PrevHash)
		}
	}
	fmt.Println("=========================")
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
