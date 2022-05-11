package signing

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Signer struct {
}

func NewSigningService() SigningService {
	return &Signer{}
}

func (s *Signer) WrapHash(h common.Hash) common.Hash {
	hash := crypto.Keccak256Hash(
		[]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n32")),
		h.Bytes(),
	)
	return hash
}

func (s *Signer) SignTxn(signingKey string, args [][]byte) (string, string, error) {
	hash := s.WrapHash(crypto.Keccak256Hash(args...))

	key, keyParseErr := crypto.HexToECDSA(signingKey)
	if keyParseErr != nil {
		fmt.Println(keyParseErr.Error())
		return "", "", keyParseErr
	}

	signature, signErr := crypto.Sign(hash.Bytes(), key)
	if signErr != nil {
		fmt.Println(signErr.Error())
		return "", "", signErr
	}

	return "0x"+hex.EncodeToString(hash.Bytes()), "0x"+hex.EncodeToString(signature), nil
}

func (s *Signer) VerifyPublicKey(signature []byte, hash common.Hash, key *ecdsa.PrivateKey) error {
	sigPublicKeyECDSA, verErr := crypto.SigToPub(hash.Bytes(), signature)
	if verErr != nil {
		return verErr
	}
	publicKey := key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("error getting public key from private key")
	}

	pubKeyStr := crypto.PubkeyToAddress(*publicKeyECDSA).String()
	if pubKeyStr != crypto.PubkeyToAddress(*sigPublicKeyECDSA).String() {
		return errors.New("signature public key verification failure")
	}
	return nil
}