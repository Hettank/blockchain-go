package main

import (
	"fmt"

	"github.com/Hettank/blockchain-go/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()

	for i := range 10 {
		bc.AddBlock(fmt.Sprintf("This is block %d", i+1))
	}

	bc.PrintBlockchain()
}
