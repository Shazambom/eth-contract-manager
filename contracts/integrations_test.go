package contracts

import (
	"context"
	pb "contract-service/proto"
	"contract-service/signing"
	"contract-service/storage"
	"contract-service/utils"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"math/big"
	"strings"
	"testing"
)

func TestStore_And_TransactionFlow(t *testing.T) {
	ctx := context.Background()

	rds := storage.NewRedisWriter(storage.RedisConfig{
		Endpoint: "localhost:6379",
		Password: "pass",
		CountKey: "Count",
	})

	_, pingErr := rds.Ping()
	assert.Nil(t, pingErr)

	contractClient, contractConnErr := NewContractClient("localhost:8082", []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	assert.Nil(t, contractConnErr)
	defer contractClient.DisconnectGracefully()

	transactionClient, transactionConnErr := NewTransactionClient("localhost:8083", []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	assert.Nil(t, transactionConnErr)
	defer transactionClient.DisconnectGracefully()

	signerClient, signerConnErr := signing.NewClient("localhost:8081", []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	assert.Nil(t, signerConnErr)
	defer signerClient.DisconnectGracefully()

	contract := &pb.Contract{
		Address:      "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		Abi: testAbi,
		HashableFunctions:    &pb.Functions{Functions: map[string]*pb.Function{"mint": {Arguments: []*pb.Argument{
			{Name: "nonce", Type: "bytes16"},
			{Name: "msg.sender", Type: "address"},
			{Name: "numberOfTokens", Type: "uint256"},
			{Name: "transactionNumber", Type: "uint256"},
		}}}},
		MaxMintable:  1000,
		MaxIncrement: 3,
		Owner:        "Owner",
	}
	
	storeErr, storeErrr := contractClient.Client.Store(ctx, contract)
	if storeErrr != nil {
		fmt.Println(storeErrr)
	}
	assert.Nil(t, storeErrr)
	assert.NotNil(t, storeErr)
	assert.Equal(t, "", storeErr.Err)

	newKey, keyErr := signerClient.SigningClient.GenerateNewKey(ctx, &pb.KeyManagementRequest{ContractAddress: contract.Address})
	assert.Nil(t, keyErr)
	fmt.Println("New signing public address: " + newKey.PublicKey)

	nonce, nonceErr := utils.GetNonce()
	assert.Nil(t, nonceErr)
	nonceBytes, decodeErr := hex.DecodeString(nonce[2:])
	assert.Nil(t, decodeErr)

	msgSender := "0x0fA37C622C7E57A06ba12afF48c846F42241F7F0"

	transactionResponse, transactionErr := transactionClient.Client.ConstructTransaction(ctx, &pb.TransactionRequest{
		MessageSender: msgSender,
		FunctionName:  "mint",
		NumRequested:  3,
		Args:          [][]byte{nonceBytes, []byte("3"), []byte("1")},
		Contract:      contract,
	})
	if transactionErr != nil {
		fmt.Println(transactionErr)
	}
	assert.Nil(t, transactionErr)
	assert.NotNil(t, transactionResponse)
	assert.Equal(t, "", transactionResponse.Err)


	token, tokenErr := rds.Get(ctx, msgSender, contract.Address)
	assert.Nil(t, tokenErr)
	fmt.Printf("Token: %+v\n", token)
	args := [][]byte{nonceBytes, common.HexToAddress(msgSender).Bytes(), common.LeftPadBytes(big.NewInt(int64(3)).Bytes(),32), common.LeftPadBytes(big.NewInt(int64(1)).Bytes(),32)}

	fmt.Println(len(args))
	for _, arg := range args {
		fmt.Println(string(arg))
	}

	signer := signing.Signer{}

	builtHash := signer.WrapHash(crypto.Keccak256Hash(args...)).String()
	assert.Equal(t, builtHash, token.Hash)
	fmt.Println("Hashes:")
	fmt.Println(token.Hash)
	fmt.Println(builtHash)


	resp, signatureErr := signerClient.SigningClient.SignTxn(ctx, &pb.SignatureRequest{ContractAddress: contract.Address, Args: args})
	assert.Nil(t, signatureErr)
	fmt.Printf("%+v\n", resp)


	funcDef, abiErr := abi.JSON(strings.NewReader(testAbi))
	assert.Nil(t, abiErr)

	packedHex := hex.EncodeToString(token.ABIPackedTxn)
	fmt.Println(packedHex)

	decodedSig, sigDecodeErr := hex.DecodeString(packedHex[:8])
	assert.Nil(t, sigDecodeErr)
	fmt.Println(decodedSig)
	method, methodErr := funcDef.MethodById(decodedSig)
	assert.Nil(t, methodErr)

	decodedData, decodedErr := hex.DecodeString(packedHex[8:])
	assert.Nil(t, decodedErr)

	unpacked, unpackErr := method.Inputs.Unpack(decodedData)
	assert.Nil(t, unpackErr)
	fmt.Println(unpacked)

	unpackedSignature := hex.EncodeToString((unpacked[len(unpacked) - 1]).([]byte))
	fmt.Println(unpackedSignature)
	fmt.Println(resp.Signature[2:])
	assert.Equal(t, resp.Signature[2:], unpackedSignature)

}

