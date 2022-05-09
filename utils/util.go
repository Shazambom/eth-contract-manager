package utils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"math/big"
	"os"
	"strconv"
	"strings"
)



type ErrorResponse struct {
	Error string `json:"error"`
}

func GetEnvVars() (rdsEndpoint, rdsPwd, secretKey, gURL, siteKey, projectID, contractAddress, signingKey string, abi []string, maxCount , maxIncr int, validOrigins []string, envErr error) {
	if enabled := os.Getenv("ENABLED"); enabled != "true" {
		return "", "", "", "", "", "", "", "", nil, 0, 0, nil, errors.New("Lambda not Enabled")
	}
	if abiErr := json.Unmarshal([]byte(os.Getenv("ABI")), &abi); abiErr != nil {
		return "", "", "", "", "", "", "", "", nil, 0, 0, nil, abiErr
	}

	maxCount, envErr = strconv.Atoi(os.Getenv("MAX_COUNT"))
	if envErr != nil {
		return "", "", "", "", "", "", "", "", nil, 0, 0, nil, envErr
	}
	maxIncr, envErr = strconv.Atoi(os.Getenv("MAX_INCR"))


	return os.Getenv("REDIS_ENDPOINT"),
		os.Getenv("REDIS_PWD"),
		os.Getenv("SECRET_KEY"),
		os.Getenv("GOOGLE_URL"),
		os.Getenv("SITE_KEY"),
		os.Getenv("PROJECT_ID"),
		os.Getenv("CONTRACT_ADDRESS"),
		os.Getenv("SIGNER_PRIVATE_KEY"),
		abi, maxCount, maxIncr,
		strings.Split(os.Getenv("VALID_ORIGINS"), "~"),
		envErr
}

func HashEvent(event main.Event) common.Hash {
	//Nonce, Address, NumAvatars, TransactionNumber, IP
	nonce, _ := hex.DecodeString(event.Nonce[2:])
	hash := crypto.Keccak256Hash(
		nonce,
		common.HexToAddress(event.Address).Bytes(),
		common.LeftPadBytes(big.NewInt(int64(event.NumAvatars)).Bytes(), 32),
		common.LeftPadBytes(big.NewInt(int64(event.TransactionNumber)).Bytes(), 32),
	)
	return hash
}

func WrapHash(h common.Hash) common.Hash {
	hash := crypto.Keccak256Hash(
		[]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n32")),
		h.Bytes(),
	)
	return hash
}

func SignTxn(event main.Event, signingKey string) (string, string, error) {
	hash := WrapHash(HashEvent(event))

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

func VerifyPublicKey(signature []byte, hash common.Hash, key *ecdsa.PrivateKey) error {
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

func GetNonce() (string, error) {
	UUID := uuid.New()
	UUIDBytes, uuidErr  := UUID.MarshalBinary()
	if uuidErr != nil {
		return "", uuidErr
	}
	return "0x"+hex.EncodeToString(UUIDBytes), nil
}

func ConstructErrorResponse(response *events.APIGatewayProxyResponse, err error) error {
	errBody, marshErr := json.Marshal(ErrorResponse{err.Error()})
	if marshErr != nil {
		return marshErr
	}
	response.Body = string(errBody)
	return nil
}

func StrInStrList(str string, strList []string) bool {
	for _, val := range strList {
		if val == str {
			return true
		}
	}
	return false
}
