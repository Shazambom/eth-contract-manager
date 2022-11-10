package signing

import (
	"github.com/ethereum/go-ethereum/common"
)

type SigningService interface {
	SignTxn(signingKey string, args [][]byte) (string, string, error)
	SignMessage(signingKey, message string) (string, string, error)
	VerifyFromHash(signature []byte, hash common.Hash, address string) error
	GenerateKey() (privateKey, address string, err error)
	PrivateKeyToAddress(privateKey string) (address string, err error)
	AddressToEIP55(address string) (string, error)
	Verify(message, signature, address string) (err error)
}
