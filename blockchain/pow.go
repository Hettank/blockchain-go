package blockchain

import (
	"fmt"
	"strings"
)

func (b *Block) MineBlock(difficulty int) {
	// Build target string based on difficulty
	var target strings.Builder

	for i := 0; i < difficulty; i++ {
		target.WriteString("0")
	}

	targetString := target.String()

	for {
		hash := b.GenerateHash()
		hashString := fmt.Sprintf("%x", hash)

		if len(hashString) >= difficulty && hashString[:difficulty] == targetString {
			b.Hash = hash
			fmt.Printf("✅ Block mined! Nonce: %d, Hash: %x\n", b.Nonce, hash[:8])
			return
		}

		b.Nonce++
	}
}
