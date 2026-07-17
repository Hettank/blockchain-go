package main

import (
	"fmt"

	"github.com/Hettank/blockchain-go/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()

	for i := range 4 {
		bc.AddBlock(fmt.Sprintf("This is block %d", i+1))
	}

	// Validate
	if bc.Validate() {
		fmt.Println("✅ Blockchain is valid!")
	} else {
		fmt.Println("❌ Blockchain is INVALID!")
	}

	// Tamper
	// bc.Blocks[2].Data = "Hello World"
	bc.Blocks[2].Data = "Hello World"

	bc.Blocks[2].Hash = bc.Blocks[2].GenerateHash()

	// Validate again
	if bc.Validate() {
		fmt.Println("✅ Blockchain is valid!")
	} else {
		fmt.Println("❌ Blockchain is INVALID!")
	}
}
