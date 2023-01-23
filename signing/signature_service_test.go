package signing

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
)

func TestNewSigningService(t *testing.T) {
	signatureService := NewSigningService()
	assert.IsType(t, &SignatureHandler{}, signatureService)
	inter := reflect.TypeOf((*SigningService)(nil)).Elem()
	assert.True(t, reflect.TypeOf(signatureService).Implements(inter))
}

func TestSignatureHandler_WrapHash(t *testing.T) {
	hash := crypto.Keccak256Hash([]byte("Hello World"))
	assert.Equal(t, s.WrapHash(hash), crypto.Keccak256Hash([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n32")), hash.Bytes()))
}

func TestSignatureHandler_WrapMessage(t *testing.T) {
	message := "Hello World"
	assert.Equal(t, s.WrapMessage(message), crypto.Keccak256Hash([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message))))
}

func TestSignatureHandler_GenerateKey(t *testing.T) {
	priv, addr, err := s.GenerateKey()
	assert.Nil(t, err)
	assert.Equal(t, len(priv), 64)
	assert.Equal(t, len(addr), 42)
	fmt.Printf("private key: %s\n", priv)
	fmt.Printf("address: %s\n", addr)
}

func TestSignatureHandler_SignTxn_and_VerifyFromHash(t *testing.T) {
	someHexArg, _ := hex.DecodeString("abdcefffffff")
	args := [][]byte{[]byte("I'm packed arguments ready to be signed"), someHexArg}
	hash, signature, err := s.SignTxn(privateKey, args)
	fmt.Printf("hash: %s\nsignature: %s\n", hash, signature)
	assert.Nil(t, err)
	assert.Equal(t, len(signature), 132)
	assert.Equal(t, len(hash), 66)
	hashBytes, hashDecodeErr := hex.DecodeString(hash[2:])
	assert.Nil(t, hashDecodeErr)
	sigBytes, decodeErr := hex.DecodeString(signature[2:])
	assert.Nil(t, decodeErr)
	assert.Nil(t, s.VerifyFromHash(sigBytes, common.BytesToHash(hashBytes), address))
}

func TestSignatureHandler_PrivateKeyToAddress(t *testing.T) {
	decodedAddress, err := s.PrivateKeyToAddress(privateKey)
	assert.Nil(t, err)
	assert.Equal(t, decodedAddress, address)
}

func TestSignatureHandler_AddressToEIP55(t *testing.T) {
	lowered := strings.ToLower(address)
	eip55Address, err := s.AddressToEIP55(lowered)
	assert.Nil(t, err)
	assert.Equal(t, eip55Address, address)
}

func TestSignatureHandler_SignMessage_and_Verify(t *testing.T) {
	message := "I'm this playerID and I can prove it!"
	_, signature, signingErr := s.SignMessage(privateKey, message)
	assert.Nil(t, signingErr)
	assert.Nil(t, s.Verify(message, signature, address))
	assert.Nil(t, s.Verify(message, signature, strings.ToLower(address)))
}

func TestSignatureHandler_SignMessage_and_Verify_InvalidSignature(t *testing.T) {
	message := "I'm this playerID and I can prove it!"
	_, signature, signingErr := s.SignMessage(privateKey, message)
	assert.Nil(t, signingErr)
	assert.Error(t, s.Verify(message, signature[1:], address))
}

func TestSignatureHandler_SignMessage_and_Verify_InvalidAddress(t *testing.T) {
	message := "I'm this playerID and I can prove it!"
	_, signature, signingErr := s.SignMessage(privateKey, message)
	assert.Nil(t, signingErr)
	assert.Error(t, s.Verify(message, signature, address[:len(address)-2]))
}

func TestSignatureHandler_SignMessage_and_Verify_ImpersonatingAddress(t *testing.T) {
	_, wrongAddr, _ := s.GenerateKey()
	message := "I'm this playerID and I can prove it!"
	_, signature, signingErr := s.SignMessage(privateKey, message)
	assert.Nil(t, signingErr)
	assert.Error(t, s.Verify(message, signature, wrongAddr))
}



