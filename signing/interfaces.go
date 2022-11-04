package signing

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
)

type SigningService interface {
	SignTxn(signingKey string, args [][]byte) (string, string, error)
	VerifyPublicKey(signature []byte, hash common.Hash, key *ecdsa.PrivateKey) error
	GenerateKey() (privateKey, address string, err error)
	PrivateKeyToAddress(privateKey string) (address string, err error)
	AddressToEIP55(address string) (string, error)
	Verify(message, signature, address string) (err error)
}
