package contracts

import (
	"bitbucket.org/artie_inc/contract-service/mocks"
	pb "bitbucket.org/artie_inc/contract-service/proto"
	"bitbucket.org/artie_inc/contract-service/storage"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
)

func newTestingContractServer(t *testing.T) (*mocks.MockContractManagerHandler, *ContractRPCService, context.Context) {
	ctrl := gomock.NewController(t)
	mockContractManager := mocks.NewMockContractManagerHandler(ctrl)
	contractServer, newServerErr := NewContractServer(getTestPort(), []grpc.ServerOption{grpc.EmptyServerOption{}}, mockContractManager)
	assert.Nil(t, newServerErr)
	return mockContractManager, contractServer, context.Background()
}

func TestNewContractServer(t *testing.T) {
	_, contractServer, _ := newTestingContractServer(t)
	defer contractServer.Server.GracefulStop()
	assert.IsType(t, &ContractRPCService{}, contractServer)
}

func TestContractRPCService_Get(t *testing.T) {
	mockContractManager, contractServer, ctx := newTestingContractServer(t)
	defer contractServer.Server.GracefulStop()
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

	mockContractManager.EXPECT().GetContract(ctx, address).Return(contract, nil)

	returnedContract, err := contractServer.Get(ctx, &pb.Address{Address: address})
	assert.Nil(t, err)
	assert.Equal(t, contract.ToRPC(), returnedContract)
}

func TestContractRPCService_Get_ReturnErr(t *testing.T) {
	mockContractManager, contractServer, ctx := newTestingContractServer(t)
	defer contractServer.Server.GracefulStop()
	address := "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
	contractManagerErr := errors.New("some contract management error")

	mockContractManager.EXPECT().GetContract(ctx, address).Return(nil, contractManagerErr)

	returnedContract, err := contractServer.Get(ctx, &pb.Address{Address: address})
	assert.Nil(t, returnedContract)
	assert.Equal(t, contractManagerErr, err)
}

func TestContractRPCService_Store(t *testing.T) {
	mockContractManager, contractServer, ctx := newTestingContractServer(t)
	defer contractServer.Server.GracefulStop()
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

	mockContractManager.EXPECT().StoreContract(ctx, contract).Return(nil)

	empty, err := contractServer.Store(ctx, contract.ToRPC())
	assert.Nil(t, err)
	assert.NotNil(t, empty)
}

func TestContractRPCService_Store_ReturnErr(t *testing.T) {
	mockContractManager, contractServer, ctx := newTestingContractServer(t)
	defer contractServer.Server.GracefulStop()
	address := "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
	contractManagerErr := errors.New("some contract management error")
	contract := &storage.Contract{
		Address: address,
		ABI:     claimAbi_Flattened,
		Functions: map[string]storage.Function{"mintArtie": {Arguments: []storage.Argument{
			{Name: "nonce", Type: "bytes16"},
			{Name: "tokenId", Type: "uint256"},
		}}},
		ContractOwner: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
	}
	mockContractManager.EXPECT().StoreContract(ctx, contract).Return(contractManagerErr)

	empty, err := contractServer.Store(ctx, contract.ToRPC())
	assert.NotNil(t, empty)
	assert.Equal(t, contractManagerErr, err)
}

func TestContractRPCService_Delete(t *testing.T) {
	mockContractManager, contractServer, ctx := newTestingContractServer(t)
	defer contractServer.Server.GracefulStop()
	address := "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
	contractOwner := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"

	mockContractManager.EXPECT().DeleteContract(ctx, address, contractOwner).Return(nil)

	empty, err := contractServer.Delete(ctx, &pb.AddressOwner{
		Address: address,
		Owner:   contractOwner,
	})
	assert.Nil(t, err)
	assert.NotNil(t, empty)
}

func TestContractRPCService_Delete_ReturnErr(t *testing.T) {
	mockContractManager, contractServer, ctx := newTestingContractServer(t)
	defer contractServer.Server.GracefulStop()
	address := "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
	contractOwner := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	contractManagerErr := errors.New("some contract management error")

	mockContractManager.EXPECT().DeleteContract(ctx, address, contractOwner).Return(contractManagerErr)

	empty, err := contractServer.Delete(ctx, &pb.AddressOwner{
		Address: address,
		Owner:   contractOwner,
	})
	assert.Equal(t, contractManagerErr, err)
	assert.NotNil(t, empty)
}

func TestContractRPCService_List(t *testing.T) {
	mockContractManager, contractServer, ctx := newTestingContractServer(t)
	defer contractServer.Server.GracefulStop()
	owner := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	contractSeasonSale := &storage.Contract{
		Address: "0xEA917326e8A95299c02655fe947962C43a11487f",
		ABI:     testAbi,
		Functions: map[string]storage.Function{"mint": {Arguments: []storage.Argument{
			{Name: "nonce", Type: "bytes16"},
			{Name: "numberOfTokens", Type: "uint256"},
			{Name: "transactionNumber", Type: "uint256"},
		}}},
		ContractOwner: owner,
	}
	contractClaim := &storage.Contract{
		Address: "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		ABI:     claimAbi_Flattened,
		Functions: map[string]storage.Function{"mintArtie": {Arguments: []storage.Argument{
			{Name: "nonce", Type: "bytes16"},
			{Name: "tokenId", Type: "uint256"},
		}}},
		ContractOwner: owner,
	}

	mockContractManager.EXPECT().ListContracts(ctx, owner).Return([]*storage.Contract{contractClaim, contractSeasonSale}, nil)

	contracts, err := contractServer.List(ctx, &pb.Owner{Owner: owner})
	assert.Nil(t, err)
	assert.Equal(t, &pb.Contracts{Contracts: []*pb.Contract{contractClaim.ToRPC(), contractSeasonSale.ToRPC()}}, contracts)
}

func TestContractRPCService_list_ReturnErr(t *testing.T) {
	mockContractManager, contractServer, ctx := newTestingContractServer(t)
	defer contractServer.Server.GracefulStop()
	contractManagerErr := errors.New("some contract management error")
	owner := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	mockContractManager.EXPECT().ListContracts(ctx, owner).Return(nil, contractManagerErr)

	contracts, err := contractServer.List(ctx, &pb.Owner{Owner: owner})
	assert.Nil(t, contracts)
	assert.Equal(t, contractManagerErr, err)
}
