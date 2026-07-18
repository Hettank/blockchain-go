package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	Timestamp    int64
	Transactions []Transaction
	PrevHash     []byte
	Hash         []byte
	Nonce        int64
}

func NewBlock(transactions []Transaction, prevHash []byte) *Block {

	b := &Block{
		Timestamp:    time.Now().Unix(),
		Transactions: transactions,
		PrevHash:     prevHash,
	}

	return b
}

func (b *Block) GenerateHash() []byte {
	serialized, err := json.Marshal(b.Transactions)

	if err != nil {
		panic(err)
	}

	info := fmt.Sprintf("%d%s%x%d", b.Timestamp, serialized, b.PrevHash, b.Nonce)
	hash := sha256.Sum256([]byte(info))

	return hash[:]
}
