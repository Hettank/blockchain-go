package blockchain

import "fmt"

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

func (bc *Blockchain) PrintBlockchain() {
	fmt.Println("=========================")

	for i, block := range bc.Blocks {
		fmt.Printf("\n Block %d:\n", i)
		fmt.Printf("   Address: %p\n", block)
		fmt.Printf("   Data: %s\n", block.Data)
		fmt.Printf("   Hash: %x\n", block.Hash)

		if i > 0 {
			fmt.Printf("   Previous Hash: %x\n", block.PrevHash)
		}
	}
	fmt.Println("=========================")
}
