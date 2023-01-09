package contracts

import (
	"context"
	pb "contract-service/proto"
	"contract-service/signing"
	"contract-service/storage"
	"contract-service/utils"
	"encoding/hex"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
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
	//Initializing all services needed for creating a contract, building a transaction, and signing it along with services to directly check that everything was stored properly
	ctx := context.Background()

	txnDB, txnDBErr := storage.NewTransactionRepo(storage.TransactionConfig{
		TableName: "Transactions",
		CFG:       []*aws.Config{{
			Endpoint:         aws.String("localhost:8000"),
			Region:           aws.String("us-east-1"),
			Credentials:      credentials.NewStaticCredentials("xxx","yyy", ""),
			DisableSSL:       aws.Bool(true),
		}},
	})
	assert.Nil(t, txnDBErr)

	//GRPC Clients for the contract, transaction, and signer services
	contractClient, contractConnErr := NewContractClient("localhost:8082", []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	assert.Nil(t, contractConnErr)
	defer contractClient.DisconnectGracefully()

	transactionClient, transactionConnErr := NewTransactionClient("localhost:8083", []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	assert.Nil(t, transactionConnErr)
	defer transactionClient.DisconnectGracefully()

	signerClient, signerConnErr := signing.NewClient("localhost:8081", []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	assert.Nil(t, signerConnErr)
	defer signerClient.DisconnectGracefully()

	//Contract using the abi for the Season01 Artie Sale contract: https://etherscan.io/address/0x8c539b123424dbb7949b9f683ac466fbadfb0699
	contract := &pb.Contract{
		Address:      "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		Abi: testAbi,
		Functions:   map[string]*pb.Function{"mint": {Arguments: []*pb.Argument{
			{Name: "nonce", Type: "bytes16"},
			{Name: "numberOfTokens", Type: "uint256"},
			{Name: "transactionNumber", Type: "uint256"},
		}}},
		Owner:        "Owner",
	}
	fmt.Printf("%+v\n", contract)
	//Storing the contract and registering it with the contract service
	_, storeErrr := contractClient.Client.Store(ctx, contract)
	if storeErrr != nil {
		fmt.Println(storeErrr)
	}
	assert.Nil(t, storeErrr)


	//Generating a signing key for the above contract
	newKey, keyErr := signerClient.SigningClient.GenerateNewKey(ctx, &pb.KeyManagementRequest{ContractAddress: contract.Address})
	assert.Nil(t, keyErr)
	fmt.Println("New signing public address: " + newKey.PublicKey)

	nonce, nonceErr := utils.GetNonce()
	assert.Nil(t, nonceErr)
	nonceBytes, decodeErr := hex.DecodeString(nonce[2:])
	assert.Nil(t, decodeErr)

	msgSender := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"

	//Building a transaction for the "mint" function by passing in the nonce, the num requested, and the transaction number
	_, transactionErr := transactionClient.Client.ConstructTransaction(ctx, &pb.TransactionRequest{
		SenderInHash: true,
		MessageSender: msgSender,
		FunctionName:  "mint",
		Args:          [][]byte{nonceBytes, []byte("3"), []byte("1")},
		ContractAddress:      contract.Address,
		Value: "450000000000000000",
	})
	if transactionErr != nil {
		fmt.Println(transactionErr)
	}
	assert.Nil(t, transactionErr)
	fmt.Printf("%+v\n", [][]byte{nonceBytes, []byte("3"), []byte("1")})


	//Checking that the token was processed correctly, the transaction was signed, and the token was placed in redis
	tokens, tokenErr := txnDB.GetTransactions(ctx, msgSender)
	assert.Nil(t, tokenErr)
	fmt.Printf("Token: %+v\n", tokens[0])
	args := [][]byte{nonceBytes, common.LeftPadBytes(big.NewInt(int64(3)).Bytes(),32), common.LeftPadBytes(big.NewInt(int64(1)).Bytes(),32)}

	fmt.Println(len(args))
	for _, arg := range args {
		fmt.Println(string(arg))
	}

	signer := signing.SignatureHandler{}

	//Checking the hashes to ensure the signer hashes the transaction properly
	senderAddrBytes, senderErr := hex.DecodeString(msgSender[2:])
	assert.Nil(t, senderErr)
	args = append([][]byte{senderAddrBytes}, args...)
	builtHash := signer.WrapHash(crypto.Keccak256Hash(args...)).String()
	assert.Equal(t, builtHash, tokens[0].Hash)
	fmt.Println("Hashes:")
	fmt.Println(tokens[0].Hash)
	fmt.Println(builtHash)


	//Manually signing the transaction with the signer service to ensure that the signature is the same in the live gRPC service and the signature service
	resp, signatureErr := signerClient.SigningClient.SignTxn(ctx, &pb.SignatureRequest{ContractAddress: contract.Address, Args: args})
	assert.Nil(t, signatureErr)
	fmt.Printf("%+v\n", resp)


	//Decoding the ABIPackedTxn data to ensure every field is packed as expected
	funcDef, abiErr := abi.JSON(strings.NewReader(testAbi))
	assert.Nil(t, abiErr)

	packedHex := hex.EncodeToString(tokens[0].ABIPackedTxn)
	fmt.Println("Packed Hex of Txn: " + packedHex)

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

	var uNonce [16]byte
	uNonce = (unpacked[0]).([16]byte)

	unpackedNonce := hex.EncodeToString(uNonce[:])

	fmt.Println(unpackedNonce)
	fmt.Println(nonce[2:])
	assert.Equal(t, nonce[2:], unpackedNonce)

	assert.Equal(t, big.NewInt(3), (unpacked[1]).(*big.Int))
	assert.Equal(t, big.NewInt(1), (unpacked[2]).(*big.Int))
}

