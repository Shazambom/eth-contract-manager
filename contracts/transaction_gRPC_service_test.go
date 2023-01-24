package contracts

import (
	"bitbucket.org/artie_inc/contract-service/mocks"
	pb "bitbucket.org/artie_inc/contract-service/proto"
	"bitbucket.org/artie_inc/contract-service/storage"
	"bitbucket.org/artie_inc/contract-service/utils"
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
)

func newTestingTransactionServer(t *testing.T) (*mocks.MockContractTransactionHandler, *TransactionRPCService, context.Context) {
	ctrl := gomock.NewController(t)
	mockTransactionHandler := mocks.NewMockContractTransactionHandler(ctrl)
	transactionServer, newServerErr := NewTransactionServer(getTestPort(), []grpc.ServerOption{grpc.EmptyServerOption{}}, mockTransactionHandler)
	assert.Nil(t, newServerErr)
	return mockTransactionHandler, transactionServer, context.Background()
}

func TestNewTransactionServer(t *testing.T) {
	_, transactionServer, _ := newTestingTransactionServer(t)
	defer transactionServer.Server.GracefulStop()
	assert.IsType(t, &TransactionRPCService{}, transactionServer)
}

func TestTransactionRPCService_GetContract(t *testing.T) {
	mockTransactionHandler, transactionServer, ctx := newTestingTransactionServer(t)
	defer transactionServer.Server.GracefulStop()
	address := "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
	contract := &storage.Contract{
		Address: address,
		ABI:     claimAbi_Flattened,
		Functions: map[string]storage.Function{"mintArtie": {Arguments: []storage.Argument{
			{Name: "nonce", Type: "bytes16"},
			{Name: "tokenId", Type: "uint256"},
		}}},
		ContractOwner: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
	}

	mockTransactionHandler.EXPECT().GetContract(ctx, address).Return(contract, nil)

	returnedContract, err := transactionServer.GetContract(ctx, &pb.Address{Address: address})
	assert.Nil(t, err)
	assert.Equal(t, contract.ToRPC(), returnedContract)
}

func TestTransactionRPCService_GetContract_Err(t *testing.T) {
	mockTransactionHandler, transactionServer, ctx := newTestingTransactionServer(t)
	defer transactionServer.Server.GracefulStop()
	address := "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
	contractManagerErr := errors.New("some contract management error")

	mockTransactionHandler.EXPECT().GetContract(ctx, address).Return(nil, contractManagerErr)

	returnedContract, err := transactionServer.GetContract(ctx, &pb.Address{Address: address})
	assert.Nil(t, returnedContract)
	assert.Equal(t, contractManagerErr, err)
}

func TestTransactionRPCService_ConstructTransaction(t *testing.T) {
	mockTransactionHandler, transactionServer, ctx := newTestingTransactionServer(t)
	defer transactionServer.Server.GracefulStop()
	nonceBytes, nonceErr := utils.GetNonceBytes()
	assert.Nil(t, nonceErr)
	contractAddress := "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
	contract := &storage.Contract{
		Address: contractAddress,
		ABI:     claimAbi_Flattened,
		Functions: map[string]storage.Function{"mintArtie": {Arguments: []storage.Argument{
			{Name: "nonce", Type: "bytes16"},
			{Name: "tokenId", Type: "uint256"},
		}}},
		ContractOwner: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
	}
	txnReq := &pb.TransactionRequest{
		SenderInHash:    true,
		MessageSender:   "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
		FunctionName:    "mintArtie",
		Args:            [][]byte{nonceBytes, []byte(fmt.Sprintf("%d", 10))},
		ContractAddress: contractAddress,
		Value:           "0",
	}
	transaction, transactionErr := storage.NewTransaction(
		contractAddress,
		txnReq.MessageSender,
		"0xce11e286abab09c3ad05f1f9fff4daaf4f5139214a1f4746661018f0b855f075",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"0")
	assert.Nil(t, transactionErr)

	mockTransactionHandler.EXPECT().GetContract(ctx, contractAddress).Return(contract, nil)
	mockTransactionHandler.EXPECT().BuildTransaction(ctx, txnReq.SenderInHash, txnReq.MessageSender, txnReq.FunctionName, txnReq.Args, txnReq.Value, contract).Return(transaction, nil)
	mockTransactionHandler.EXPECT().StoreTransaction(ctx, transaction).Return(nil)

	txn, err := transactionServer.ConstructTransaction(ctx, txnReq)
	assert.Nil(t, err)
	assert.Equal(t, transaction.ToRPC(), txn)
}

func TestTransactionRPCService_ConstructTransaction_ErrGetContract(t *testing.T) {
	mockTransactionHandler, transactionServer, ctx := newTestingTransactionServer(t)
	defer transactionServer.Server.GracefulStop()
	nonceBytes, nonceErr := utils.GetNonceBytes()
	assert.Nil(t, nonceErr)
	transactionHandlerErr := errors.New("some transaction handler error")
	contractAddress := "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
	txnReq := &pb.TransactionRequest{
		SenderInHash:    true,
		MessageSender:   "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
		FunctionName:    "mintArtie",
		Args:            [][]byte{nonceBytes, []byte(fmt.Sprintf("%d", 10))},
		ContractAddress: contractAddress,
		Value:           "0",
	}

	mockTransactionHandler.EXPECT().GetContract(ctx, contractAddress).Return(nil, transactionHandlerErr)

	txn, err := transactionServer.ConstructTransaction(ctx, txnReq)
	assert.Nil(t, txn)
	assert.Equal(t, transactionHandlerErr, err)
}

func TestTransactionRPCService_ConstructTransaction_ErrBuildTransaction(t *testing.T) {
	mockTransactionHandler, transactionServer, ctx := newTestingTransactionServer(t)
	defer transactionServer.Server.GracefulStop()
	nonceBytes, nonceErr := utils.GetNonceBytes()
	assert.Nil(t, nonceErr)
	transactionHandlerErr := errors.New("some transaction handler error")
	contractAddress := "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
	contract := &storage.Contract{
		Address: contractAddress,
		ABI:     claimAbi_Flattened,
		Functions: map[string]storage.Function{"mintArtie": {Arguments: []storage.Argument{
			{Name: "nonce", Type: "bytes16"},
			{Name: "tokenId", Type: "uint256"},
		}}},
		ContractOwner: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
	}
	txnReq := &pb.TransactionRequest{
		SenderInHash:    true,
		MessageSender:   "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
		FunctionName:    "mintArtie",
		Args:            [][]byte{nonceBytes, []byte(fmt.Sprintf("%d", 10))},
		ContractAddress: contractAddress,
		Value:           "0",
	}

	mockTransactionHandler.EXPECT().GetContract(ctx, contractAddress).Return(contract, nil)
	mockTransactionHandler.EXPECT().BuildTransaction(ctx, txnReq.SenderInHash, txnReq.MessageSender, txnReq.FunctionName, txnReq.Args, txnReq.Value, contract).Return(nil, transactionHandlerErr)

	txn, err := transactionServer.ConstructTransaction(ctx, txnReq)
	assert.Nil(t, txn)
	assert.Equal(t, transactionHandlerErr, err)
}

func TestTransactionRPCService_ConstructTransaction_ErrStoringTxn(t *testing.T) {
	mockTransactionHandler, transactionServer, ctx := newTestingTransactionServer(t)
	defer transactionServer.Server.GracefulStop()
	nonceBytes, nonceErr := utils.GetNonceBytes()
	transactionHandlerErr := errors.New("some transaction handler error")
	assert.Nil(t, nonceErr)
	contractAddress := "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
	contract := &storage.Contract{
		Address: contractAddress,
		ABI:     claimAbi_Flattened,
		Functions: map[string]storage.Function{"mintArtie": {Arguments: []storage.Argument{
			{Name: "nonce", Type: "bytes16"},
			{Name: "tokenId", Type: "uint256"},
		}}},
		ContractOwner: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
	}
	txnReq := &pb.TransactionRequest{
		SenderInHash:    true,
		MessageSender:   "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
		FunctionName:    "mintArtie",
		Args:            [][]byte{nonceBytes, []byte(fmt.Sprintf("%d", 10))},
		ContractAddress: contractAddress,
		Value:           "0",
	}
	transaction, transactionErr := storage.NewTransaction(
		contractAddress,
		txnReq.MessageSender,
		"0xce11e286abab09c3ad05f1f9fff4daaf4f5139214a1f4746661018f0b855f075",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"0")
	assert.Nil(t, transactionErr)

	mockTransactionHandler.EXPECT().GetContract(ctx, contractAddress).Return(contract, nil)
	mockTransactionHandler.EXPECT().BuildTransaction(ctx, txnReq.SenderInHash, txnReq.MessageSender, txnReq.FunctionName, txnReq.Args, txnReq.Value, contract).Return(transaction, nil)
	mockTransactionHandler.EXPECT().StoreTransaction(ctx, transaction).Return(transactionHandlerErr)

	txn, err := transactionServer.ConstructTransaction(ctx, txnReq)
	assert.Nil(t, txn)
	assert.Equal(t, transactionHandlerErr, err)
}

func TestTransactionRPCService_GetTransactions(t *testing.T) {
	mockTransactionHandler, transactionServer, ctx := newTestingTransactionServer(t)
	defer transactionServer.Server.GracefulStop()

	contractAddress := "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
	msgSender := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	transaction, transactionErr := storage.NewTransaction(
		contractAddress,
		msgSender,
		"0xce11e286abab09c3ad05f1f9fff4daaf4f5139214a1f4746661018f0b855f075",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"0")
	assert.Nil(t, transactionErr)

	mockTransactionHandler.EXPECT().GetTransactions(ctx, msgSender).Return([]*storage.Transaction{transaction, transaction}, nil)

	txns, err := transactionServer.GetTransactions(ctx, &pb.Address{Address: msgSender})
	assert.Nil(t, err)
	assert.Equal(t, &pb.Transactions{Transactions: []*pb.Transaction{transaction.ToRPC(), transaction.ToRPC()}}, txns)
}

func TestTransactionRPCService_GetTransactions_ErrGettingTxns(t *testing.T) {
	mockTransactionHandler, transactionServer, ctx := newTestingTransactionServer(t)
	defer transactionServer.Server.GracefulStop()
	transactionHandlerErr := errors.New("some transaction handler error")

	msgSender := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"

	mockTransactionHandler.EXPECT().GetTransactions(ctx, msgSender).Return(nil, transactionHandlerErr)

	txns, err := transactionServer.GetTransactions(ctx, &pb.Address{Address: msgSender})
	assert.Nil(t, txns)
	assert.Equal(t, transactionHandlerErr, err)
}

func TestTransactionRPCService_GetAllTransactions(t *testing.T) {
	mockTransactionHandler, transactionServer, ctx := newTestingTransactionServer(t)
	defer transactionServer.Server.GracefulStop()

	contractAddress := "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
	msgSender := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	transaction, transactionErr := storage.NewTransaction(
		contractAddress,
		msgSender,
		"0xce11e286abab09c3ad05f1f9fff4daaf4f5139214a1f4746661018f0b855f075",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"0")
	assert.Nil(t, transactionErr)

	completedTxn, completedTxnErr := storage.NewTransaction(
		contractAddress,
		msgSender,
		"0xce11e286abab09c3ad05f1f9fff4daaf4f5139214a1f4746661018f0b855f075",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"0")
	assert.Nil(t, completedTxnErr)
	completedTxn.IsComplete = true

	mockTransactionHandler.EXPECT().GetAllTransactions(ctx, msgSender).Return([]*storage.Transaction{transaction, completedTxn}, nil)

	txns, err := transactionServer.GetAllTransactions(ctx, &pb.Address{Address: msgSender})
	assert.Nil(t, err)
	assert.Equal(t, &pb.Transactions{Transactions: []*pb.Transaction{transaction.ToRPC(), completedTxn.ToRPC()}}, txns)
}

func TestTransactionRPCService_GetAllTransactions_ErrGettingTxns(t *testing.T) {
	mockTransactionHandler, transactionServer, ctx := newTestingTransactionServer(t)
	defer transactionServer.Server.GracefulStop()
	transactionHandlerErr := errors.New("some transaction handler error")

	msgSender := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"

	mockTransactionHandler.EXPECT().GetAllTransactions(ctx, msgSender).Return(nil, transactionHandlerErr)

	txns, err := transactionServer.GetAllTransactions(ctx, &pb.Address{Address: msgSender})
	assert.Nil(t, txns)
	assert.Equal(t, transactionHandlerErr, err)
}

func TestTransactionRPCService_CompleteTransaction(t *testing.T) {
	mockTransactionHandler, transactionServer, ctx := newTestingTransactionServer(t)
	defer transactionServer.Server.GracefulStop()

	msgSender := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	req := &pb.KeyTransactionRequest{
		Address: msgSender,
		Hash:    "0xce11e286abab09c3ad05f1f9fff4daaf4f5139214a1f4746661018f0b855f075",
	}

	mockTransactionHandler.EXPECT().CompleteTransaction(ctx, msgSender, req.Hash).Return(nil)

	empty, err := transactionServer.CompleteTransaction(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, empty)
}

func TestTransactionRPCService_CompleteTransaction_Err(t *testing.T) {
	mockTransactionHandler, transactionServer, ctx := newTestingTransactionServer(t)
	defer transactionServer.Server.GracefulStop()

	transactionHandlerErr := errors.New("some transaction handler error")
	msgSender := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	req := &pb.KeyTransactionRequest{
		Address: msgSender,
		Hash:    "0xce11e286abab09c3ad05f1f9fff4daaf4f5139214a1f4746661018f0b855f075",
	}

	mockTransactionHandler.EXPECT().CompleteTransaction(ctx, msgSender, req.Hash).Return(transactionHandlerErr)

	empty, err := transactionServer.CompleteTransaction(ctx, req)
	assert.NotNil(t, empty)
	assert.Equal(t, transactionHandlerErr, err)
}

func TestTransactionRPCService_DeleteTransaction(t *testing.T) {
	mockTransactionHandler, transactionServer, ctx := newTestingTransactionServer(t)
	defer transactionServer.Server.GracefulStop()

	msgSender := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	req := &pb.KeyTransactionRequest{
		Address: msgSender,
		Hash:    "0xce11e286abab09c3ad05f1f9fff4daaf4f5139214a1f4746661018f0b855f075",
	}

	mockTransactionHandler.EXPECT().DeleteTransaction(ctx, msgSender, req.Hash).Return(nil)

	empty, err := transactionServer.DeleteTransaction(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, empty)
}

func TestTransactionRPCService_DeleteTransaction_Err(t *testing.T) {
	mockTransactionHandler, transactionServer, ctx := newTestingTransactionServer(t)
	defer transactionServer.Server.GracefulStop()

	transactionHandlerErr := errors.New("some transaction handler error")
	msgSender := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	req := &pb.KeyTransactionRequest{
		Address: msgSender,
		Hash:    "0xce11e286abab09c3ad05f1f9fff4daaf4f5139214a1f4746661018f0b855f075",
	}

	mockTransactionHandler.EXPECT().DeleteTransaction(ctx, msgSender, req.Hash).Return(transactionHandlerErr)

	empty, err := transactionServer.DeleteTransaction(ctx, req)
	assert.NotNil(t, empty)
	assert.Equal(t, transactionHandlerErr, err)
}
