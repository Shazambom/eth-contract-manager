package signing

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type SignatureHandler struct {
}

func NewSigningService() SigningService {
	return &SignatureHandler{}
}

func (s *SignatureHandler) WrapHash(h common.Hash) common.Hash {
	hash := crypto.Keccak256Hash(
		[]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n32")),
		h.Bytes(),
	)
	return hash
}

func (s *SignatureHandler) WrapMessage(message string) common.Hash {
	return crypto.Keccak256Hash([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)))
}

func (s *SignatureHandler) SignTxn(signingKey string, args [][]byte) (string, string, error) {
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

	if signature[crypto.RecoveryIDOffset] == 0 || signature[crypto.RecoveryIDOffset] == 1 {
		signature[crypto.RecoveryIDOffset] += 27 // Transform yellow paper V from 0/1 to 27/28
	}

	return "0x"+hex.EncodeToString(hash.Bytes()), "0x"+hex.EncodeToString(signature), nil
}

func (s *SignatureHandler) VerifyPublicKey(signature []byte, hash common.Hash, key *ecdsa.PrivateKey) error {
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

func (s *SignatureHandler) GenerateKey() (privateKey, address string, err error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}
	privateKeyStr := hexutil.Encode(crypto.FromECDSA(key))[2:]
	address, addrErr := s.PrivateKeyToAddress(privateKeyStr)
	if addrErr != nil {
		return "", "", addrErr
	}
	return privateKeyStr, address, nil
}

func (s *SignatureHandler) PrivateKeyToAddress(privateKey string) (address string, err error) {
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}
	publicKeyECDSA, ok := key.Public().(*ecdsa.PublicKey)
	if !ok {
		return  "", errors.New("error building public key")
	}
	return crypto.PubkeyToAddress(*publicKeyECDSA).Hex(), nil
}

func (s *SignatureHandler) AddressToEIP55(address string) (string, error) {
	validAddress, err := common.NewMixedcaseAddressFromString(address)
	if err != nil {
		return "", err
	}
	return validAddress.Original(), nil
}

func (s *SignatureHandler) Verify(message, signature, address string) (err error) {
	sig := hexutil.MustDecode(signature)
	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}
	hash := s.WrapMessage(message)
	sigPublicKeyECDSA, verErr := crypto.SigToPub(hash.Bytes(), sig)
	if verErr != nil {
		fmt.Println(verErr)
		return verErr
	}

	sigAddress := crypto.PubkeyToAddress(*sigPublicKeyECDSA).Hex()
	//Make sure the addresses are valid
	validAddress, addressValidationErr := common.NewMixedcaseAddressFromString(sigAddress)
	if addressValidationErr != nil {
		fmt.Println("sigAddress: " + addressValidationErr.Error())
		fmt.Println(sigAddress)
		return addressValidationErr
	}
	eventValidAddress, eventAddressValidationErr := common.NewMixedcaseAddressFromString(address)
	if eventAddressValidationErr != nil {
		fmt.Println("givenAddress: " + eventAddressValidationErr.Error())
		fmt.Println(address)
		return eventAddressValidationErr
	}

	if validAddress.String() != eventValidAddress.String() {
		errMsg := fmt.Sprintf(
			"\nhash generated: %s\n address given: %s\naddress retrieved: %s\nmessage: %s\nsignature: %s\n",
			hash.Hex(), address, sigAddress, message, signature)
		fmt.Println(errMsg)
		sigErr := errors.New("invalid signature: " + errMsg)
		return sigErr
	}

	return nil
}