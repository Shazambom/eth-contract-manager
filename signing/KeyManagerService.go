package signing

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type KeyManager struct {

}

func NewKeyManagerService() KeyManagerService {
	return &KeyManager{}
}

func (km *KeyManager) GenerateKey() (privateKey, address string, err error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}
	privateKeyBytes := crypto.FromECDSA(key)
	publicKeyECDSA, ok := key.Public().(*ecdsa.PublicKey)
	if !ok {
		return "", "", errors.New("error building public key")
	}
	addr := crypto.PubkeyToAddress(*publicKeyECDSA)
	return hexutil.Encode(privateKeyBytes)[2:], addr.Hex(), nil
}