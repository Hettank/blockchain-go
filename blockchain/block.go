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
	MinerAddress string
	Reward       int
}

func NewBlock(transactions []Transaction, prevHash []byte, minerAddress string) *Block {

	b := &Block{
		Timestamp:    time.Now().Unix(),
		Transactions: transactions,
		PrevHash:     prevHash,
		MinerAddress: minerAddress,
		Reward:       MiningReward,
	}

	return b
}

func (b *Block) GenerateHash() []byte {
	serialized, err := json.Marshal(b.Transactions)

	if err != nil {
		panic(err)
	}

	info := fmt.Sprintf(
		"%d%s%x%d%s%d",
		b.Timestamp,
		serialized,
		b.PrevHash,
		b.Nonce,
		b.MinerAddress,
		b.Reward,
	)

	hash := sha256.Sum256([]byte(info))

	return hash[:]
}
