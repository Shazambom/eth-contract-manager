package signing

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
)

type SigningService interface {
	SignTxn(signingKey string, args [][]byte) (string, string, error)
	VerifyPublicKey(signature []byte, hash common.Hash, key *ecdsa.PrivateKey) error
}

type KeyManagerService interface {
	GenerateKey() (privateKey, address string, err error)
}
