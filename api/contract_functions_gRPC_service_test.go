package api

import (
	"bitbucket.org/artie_inc/contract-service/contracts"
	"bitbucket.org/artie_inc/contract-service/mocks"
	pb "bitbucket.org/artie_inc/contract-service/proto"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
)

func newContractFunctionsServer(t *testing.T) (*mocks.MockTransactionServiceClient, *ContractIntegrationRPC, context.Context) {
	ctrl := gomock.NewController(t)
	mockTransactionServiceClient := mocks.NewMockTransactionServiceClient(ctrl)
	transactionClient := &contracts.TransactionClient{
		Connection: nil,
		Client:     mockTransactionServiceClient,
	}
	contractIntegrationRPC, newServerErr := NewContractIntegrationRPCService(getTestPort(), []grpc.ServerOption{grpc.EmptyServerOption{}}, transactionClient)
	assert.Nil(t, newServerErr)
	return mockTransactionServiceClient, contractIntegrationRPC, context.Background()
}


func TestNewContractIntegrationRPCService(t *testing.T) {
	_, contractIntegrationRPC, _ := newContractFunctionsServer(t)
	defer contractIntegrationRPC.Server.GracefulStop()
	assert.IsType(t, &ContractIntegrationRPC{}, contractIntegrationRPC)
}

//TODO Improve the mocking of these tests, implement a custom gomock.Matcher, example here: https://github.com/golang/mock/issues/43

func TestContractIntegrationRPC_BuildClaimTransaction(t *testing.T) {
	mockTransactionClient, contractIntegrationRPC, ctx := newContractFunctionsServer(t)
	defer contractIntegrationRPC.Server.GracefulStop()

	contractAddress := "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
	msgSender := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	req := &pb.ClaimRequest{
		MessageSender:   msgSender,
		TokenId:         100,
		ContractAddress: contractAddress,
	}


	mockTransactionClient.EXPECT().ConstructTransaction(ctx, gomock.AssignableToTypeOf(&pb.TransactionRequest{})).Return(&pb.Transaction{}, nil)

	resp, err := contractIntegrationRPC.BuildClaimTransaction(ctx, req)
	assert.Nil(t, err)
	assert.Equal(t, &pb.MintResponse{Status: pb.Code_CODE_SUCCESS}, resp)
}


func TestContractIntegrationRPC_BuildClaimTransaction_ErrBuildingTxn(t *testing.T) {
	mockTransactionClient, contractIntegrationRPC, ctx := newContractFunctionsServer(t)
	defer contractIntegrationRPC.Server.GracefulStop()

	contractAddress := "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
	msgSender := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	req := &pb.ClaimRequest{
		MessageSender:   msgSender,
		TokenId:         100,
		ContractAddress: contractAddress,
	}

	transactionServiceErr := errors.New("some txn service err")

	mockTransactionClient.EXPECT().ConstructTransaction(ctx, gomock.AssignableToTypeOf(&pb.TransactionRequest{})).Return(nil, transactionServiceErr)

	resp, err := contractIntegrationRPC.BuildClaimTransaction(ctx, req)
	assert.Nil(t, err)
	assert.Equal(t, &pb.MintResponse{Status: pb.Code_CODE_INTERNAL_SERVER_ERROR, Message: transactionServiceErr.Error()}, resp)
}


func TestContractIntegrationRPC_BuildMintTransaction(t *testing.T) {
	mockTransactionClient, contractIntegrationRPC, ctx := newContractFunctionsServer(t)
	defer contractIntegrationRPC.Server.GracefulStop()

	contractAddress := "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
	msgSender := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	req := &pb.MintRequest{
		MessageSender:   msgSender,
		ContractAddress: contractAddress,
		TransactionNumber: 500,
		NumberOfTokens: 3,
		Value: "450000000000000000",
	}


	mockTransactionClient.EXPECT().ConstructTransaction(ctx, gomock.AssignableToTypeOf(&pb.TransactionRequest{})).Return(&pb.Transaction{}, nil)

	resp, err := contractIntegrationRPC.BuildMintTransaction(ctx, req)
	assert.Nil(t, err)
	assert.Equal(t, &pb.MintResponse{Status: pb.Code_CODE_SUCCESS}, resp)
}


func TestContractIntegrationRPC_BuildMintTransaction_ErrBuildingTxn(t *testing.T) {
	mockTransactionClient, contractIntegrationRPC, ctx := newContractFunctionsServer(t)
	defer contractIntegrationRPC.Server.GracefulStop()

	contractAddress := "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
	msgSender := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	req := &pb.MintRequest{
		MessageSender:   msgSender,
		ContractAddress: contractAddress,
		TransactionNumber: 500,
		NumberOfTokens: 3,
		Value: "450000000000000000",
	}

	transactionServiceErr := errors.New("some txn service err")

	mockTransactionClient.EXPECT().ConstructTransaction(ctx, gomock.AssignableToTypeOf(&pb.TransactionRequest{})).Return(nil, transactionServiceErr)

	resp, err := contractIntegrationRPC.BuildMintTransaction(ctx, req)
	assert.Nil(t, err)
	assert.Equal(t, &pb.MintResponse{Status: pb.Code_CODE_INTERNAL_SERVER_ERROR, Message: transactionServiceErr.Error()}, resp)
}
