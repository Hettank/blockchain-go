package blockchain

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
	Address    string
}

func NewWallet() *Wallet {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("cannot assert type")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return &Wallet{
		PrivateKey: privateKey,
		PublicKey:  publicKeyECDSA,
		Address:    address,
	}
}

func (w *Wallet) Sign(tx *Transaction) {
	hash := tx.Hash()

	signature, err := crypto.Sign(hash, w.PrivateKey)
	if err != nil {
		panic(err)
	}

	tx.Signature = signature
	tx.PublicKey = crypto.FromECDSAPub(w.PublicKey)
}
