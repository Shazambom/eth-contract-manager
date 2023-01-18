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

func (s *SignatureHandler) SignMessage(signingKey, message string) (string, string, error) {
	key, keyParseErr := crypto.HexToECDSA(signingKey)
	if keyParseErr != nil {
		fmt.Println(keyParseErr.Error())
		return "", "", keyParseErr
	}
	hash := s.WrapMessage(message)
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

func (s *SignatureHandler) VerifyFromHash(signature []byte, hash common.Hash, address string) error {
	if signature[crypto.RecoveryIDOffset] == 27 || signature[crypto.RecoveryIDOffset] == 28 {
		signature[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}
	sigPublicKeyECDSA, verErr := crypto.SigToPub(hash.Bytes(), signature)
	if verErr != nil {
		return verErr
	}
	recoveredAddress := crypto.PubkeyToAddress(*sigPublicKeyECDSA).String()
	if address != recoveredAddress {
		fmt.Printf("address input: %s\n, address recovered: %s\n", address, recoveredAddress)
		return errors.New("signature public key verification failure")
	}
	return nil
}

func (s *SignatureHandler) GenerateKey() (privateKey, address string, err error) {
	key, genErr := crypto.GenerateKey()
	if genErr != nil {
		return "", "", genErr
	}
	privateKeyStr := hexutil.Encode(crypto.FromECDSA(key))[2:]
	addr, addrErr := s.PrivateKeyToAddress(privateKeyStr)
	if addrErr != nil {
		return "", "", addrErr
	}
	return privateKeyStr, addr, nil
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
	return validAddress.Address().String(), nil
}

func (s *SignatureHandler) Verify(message, signature, address string) (err error) {
	sig, decodeErr := hexutil.Decode(signature)
	if decodeErr != nil {
		fmt.Println("error decoding signature: " + signature)
		return decodeErr
	}
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
	validAddress, addressValidationErr := s.AddressToEIP55(sigAddress)
	if addressValidationErr != nil {
		fmt.Println("sigAddress: " + addressValidationErr.Error())
		fmt.Println(sigAddress)
		return addressValidationErr
	}
	eventValidAddress, eventAddressValidationErr := s.AddressToEIP55(address)
	if eventAddressValidationErr != nil {
		fmt.Println("givenAddress: " + eventAddressValidationErr.Error())
		fmt.Println(address)
		return eventAddressValidationErr
	}

	if validAddress != eventValidAddress {
		errMsg := fmt.Sprintf(
			"\nhash generated: %s\n address given: %s\naddress retrieved: %s\nmessage: %s\nsignature: %s\n",
			hash.Hex(), address, sigAddress, message, signature)
		fmt.Println(errMsg)
		sigErr := errors.New("invalid signature: " + errMsg)
		return sigErr
	}

	return nil
}