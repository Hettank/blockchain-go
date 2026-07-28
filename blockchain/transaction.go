package blockchain

import (
	"crypto/sha256"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

type Transaction struct {
	From      string
	To        string
	Amount    int
	PublicKey []byte
	Signature []byte
}

func (tx *Transaction) Hash() []byte {
	info := fmt.Sprintf("%s%s%d", tx.From, tx.To, tx.Amount)
	hash := sha256.Sum256([]byte(info))

	return hash[:]
}

func (tx *Transaction) Verify() bool {
	hash := tx.Hash()

	return crypto.VerifySignature(
		tx.PublicKey,
		hash,
		tx.Signature[:64],
	)
}
