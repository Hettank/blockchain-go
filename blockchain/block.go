package blockchain

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	Timestamp int64
	Data      string
	PrevHash  []byte
	Hash      []byte
}

func NewBlock(data string, prevHash []byte) *Block {

	b := &Block{
		Timestamp: time.Now().Unix(),
		Data:      data,
		PrevHash:  prevHash,
	}

	b.Hash = b.GenerateHash()
	return b
}

func (b *Block) GenerateHash() []byte {
	info := fmt.Sprintf("%d%s%s", b.Timestamp, b.Data, b.PrevHash)
	hash := sha256.Sum256([]byte(info))

	return hash[:]
}
