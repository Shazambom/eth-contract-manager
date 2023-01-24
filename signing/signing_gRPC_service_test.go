package signing

import (
	"bitbucket.org/artie_inc/contract-service/mocks"
	pb "bitbucket.org/artie_inc/contract-service/proto"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
)

var mockContractAddress = "mockContractAddress"

func newTestingSignerServer(t *testing.T) (*mocks.MockSigningService, *mocks.MockPrivateKeyRepository, *SignerRPCService, context.Context) {
	ctrl := gomock.NewController(t)
	mockSigningService := mocks.NewMockSigningService(ctrl)
	mockPrivateKeyRepo := mocks.NewMockPrivateKeyRepository(ctrl)
	signerServer, newServerErr := NewSignerServer(getTestPort(), []grpc.ServerOption{grpc.EmptyServerOption{}}, mockSigningService, mockPrivateKeyRepo)
	assert.Nil(t, newServerErr)
	return mockSigningService, mockPrivateKeyRepo, signerServer, context.Background()
}

func TestNewSignerServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	signerServer, err := NewSignerServer(getTestPort(), []grpc.ServerOption{grpc.EmptyServerOption{}}, mocks.NewMockSigningService(ctrl), mocks.NewMockPrivateKeyRepository(ctrl))
	assert.Nil(t, err)
	defer signerServer.Server.GracefulStop()
	assert.IsType(t, &SignerRPCService{}, signerServer)
}

func TestSignerRPCService_GenerateNewKey(t *testing.T) {
	mockSigningService, mockPrivateKeyRepo, signerServer, ctx := newTestingSignerServer(t)
	defer signerServer.Server.GracefulStop()

	mockSigningService.EXPECT().GenerateKey().Return(privateKey, address, nil)
	mockPrivateKeyRepo.EXPECT().UpsertPrivateKey(ctx, mockContractAddress, privateKey).Return(nil)

	resp, genKeyErr := signerServer.GenerateNewKey(ctx, &pb.KeyManagementRequest{ContractAddress: mockContractAddress})
	assert.Nil(t, genKeyErr)
	assert.Equal(t, mockContractAddress, resp.ContractAddress)
	assert.Equal(t, address, resp.PublicKey)
}

func TestSignerRPCService_GenerateNewKey_InvalidKeyReturned(t *testing.T) {
	mockSigningService, _, signerServer, ctx := newTestingSignerServer(t)
	defer signerServer.Server.GracefulStop()

	err := errors.New("testing error for testing")
	mockSigningService.EXPECT().GenerateKey().Return("", "", err)

	resp, genKeyErr := signerServer.GenerateNewKey(ctx, &pb.KeyManagementRequest{ContractAddress: mockContractAddress})

	assert.Nil(t, resp)
	assert.Error(t, genKeyErr)
	assert.Equal(t, err, genKeyErr)
}
func TestSignerRPCService_GenerateNewKey_UpsertFailed(t *testing.T) {
	mockSigningService, mockPrivateKeyRepo, signerServer, ctx := newTestingSignerServer(t)
	defer signerServer.Server.GracefulStop()

	err := errors.New("testing error for testing")
	mockSigningService.EXPECT().GenerateKey().Return(privateKey, address, nil)
	mockPrivateKeyRepo.EXPECT().UpsertPrivateKey(ctx, mockContractAddress, privateKey).Return(err)

	resp, genKeyErr := signerServer.GenerateNewKey(ctx, &pb.KeyManagementRequest{ContractAddress: mockContractAddress})
	assert.Nil(t, resp)
	assert.Error(t, genKeyErr)
	assert.Equal(t, err, genKeyErr)
}

func TestSignerRPCService_DeleteKey(t *testing.T) {
	_, mockPrivateKeyRepo, signerServer, ctx := newTestingSignerServer(t)
	defer signerServer.Server.GracefulStop()

	mockPrivateKeyRepo.EXPECT().DeletePrivateKey(ctx, mockContractAddress).Return(nil)

	resp, deleteKeyErr := signerServer.DeleteKey(ctx, &pb.KeyManagementRequest{ContractAddress: mockContractAddress})
	assert.Nil(t, deleteKeyErr)
	assert.Equal(t, mockContractAddress, resp.ContractAddress)
}

func TestSignerRPCService_DeleteKey_DeleteErr(t *testing.T) {
	_, mockPrivateKeyRepo, signerServer, ctx := newTestingSignerServer(t)
	defer signerServer.Server.GracefulStop()

	err := errors.New("testing error for testing")
	mockPrivateKeyRepo.EXPECT().DeletePrivateKey(ctx, mockContractAddress).Return(err)

	resp, deleteKeyErr := signerServer.DeleteKey(ctx, &pb.KeyManagementRequest{ContractAddress: mockContractAddress})
	assert.Equal(t, err, deleteKeyErr)
	assert.Nil(t, resp)
}

func TestSignerRPCService_SignTxn(t *testing.T) {
	mockSigningService, mockPrivateKeyRepo, signerServer, ctx := newTestingSignerServer(t)
	defer signerServer.Server.GracefulStop()

	args := [][]byte{[]byte("some arguments or something")}
	hash := "0xI'mAHash"
	signature := "0xI'mASignature"
	mockPrivateKeyRepo.EXPECT().GetPrivateKey(ctx, mockContractAddress).Return(privateKey, nil)
	mockSigningService.EXPECT().SignTxn(privateKey, args).Return(hash, signature, nil)

	resp, signTxnErr := signerServer.SignTxn(ctx, &pb.SignatureRequest{ContractAddress: mockContractAddress, Args: args})
	assert.Nil(t, signTxnErr)
	assert.Equal(t, hash, resp.Hash)
	assert.Equal(t, signature, resp.Signature)
}

func TestSignerRPCService_SignTxn_PrivateKeyErr(t *testing.T) {
	_, mockPrivateKeyRepo, signerServer, ctx := newTestingSignerServer(t)
	defer signerServer.Server.GracefulStop()

	err := errors.New("testing error for testing")
	args := [][]byte{[]byte("some arguments or something")}
	mockPrivateKeyRepo.EXPECT().GetPrivateKey(ctx, mockContractAddress).Return("", err)

	resp, signTxnErr := signerServer.SignTxn(ctx, &pb.SignatureRequest{ContractAddress: mockContractAddress, Args: args})
	assert.Equal(t, err, signTxnErr)
	assert.Nil(t, resp)
}

func TestSignerRPCService_SignTxn_ErrSigning(t *testing.T) {
	mockSigningService, mockPrivateKeyRepo, signerServer, ctx := newTestingSignerServer(t)
	defer signerServer.Server.GracefulStop()

	err := errors.New("testing error for testing")
	args := [][]byte{[]byte("some arguments or something")}
	mockPrivateKeyRepo.EXPECT().GetPrivateKey(ctx, mockContractAddress).Return(privateKey, nil)
	mockSigningService.EXPECT().SignTxn(privateKey, args).Return("", "", err)

	resp, signTxnErr := signerServer.SignTxn(ctx, &pb.SignatureRequest{ContractAddress: mockContractAddress, Args: args})
	assert.Equal(t, err, signTxnErr)
	assert.Nil(t, resp)
}

func TestSignerRPCService_GetKey(t *testing.T) {
	mockSigningService, mockPrivateKeyRepo, signerServer, ctx := newTestingSignerServer(t)
	defer signerServer.Server.GracefulStop()

	mockPrivateKeyRepo.EXPECT().GetPrivateKey(ctx, mockContractAddress).Return(privateKey, nil)
	mockSigningService.EXPECT().PrivateKeyToAddress(privateKey).Return(address, nil)

	resp, signTxnErr := signerServer.GetKey(ctx, &pb.KeyManagementRequest{ContractAddress: mockContractAddress})
	assert.Nil(t, signTxnErr)
	assert.Equal(t, mockContractAddress, resp.ContractAddress)
	assert.Equal(t, address, resp.PublicKey)
}

func TestSignerRPCService_GetKey_ErrGetKey(t *testing.T) {
	_, mockPrivateKeyRepo, signerServer, ctx := newTestingSignerServer(t)
	defer signerServer.Server.GracefulStop()

	err := errors.New("testing error for testing")
	mockPrivateKeyRepo.EXPECT().GetPrivateKey(ctx, mockContractAddress).Return("", err)

	resp, signTxnErr := signerServer.GetKey(ctx, &pb.KeyManagementRequest{ContractAddress: mockContractAddress})
	assert.Equal(t, err, signTxnErr)
	assert.Nil(t, resp)
}

func TestSignerRPCService_GetKey_ErrPrivateKeyConversion(t *testing.T) {
	mockSigningService, mockPrivateKeyRepo, signerServer, ctx := newTestingSignerServer(t)
	defer signerServer.Server.GracefulStop()

	err := errors.New("testing error for testing")
	mockPrivateKeyRepo.EXPECT().GetPrivateKey(ctx, mockContractAddress).Return(privateKey, nil)
	mockSigningService.EXPECT().PrivateKeyToAddress(privateKey).Return("", err)

	resp, signTxnErr := signerServer.GetKey(ctx, &pb.KeyManagementRequest{ContractAddress: mockContractAddress})
	assert.Equal(t, err, signTxnErr)
	assert.Nil(t, resp)
}

func TestSignerRPCService_BatchSignTxn(t *testing.T) {
	mockSigningService, mockPrivateKeyRepo, signerServer, ctx := newTestingSignerServer(t)
	defer signerServer.Server.GracefulStop()

	args := [][]byte{[]byte("some arguments or something")}
	requests := []*pb.SignatureRequest{}
	numRequests := 10
	for i := 0; i < numRequests; i++ {
		requests = append(requests, &pb.SignatureRequest{
			Args:            args,
			ContractAddress: mockContractAddress,
		})
	}
	hash := "0xI'mAHash"
	signature := "0xI'mASignature"
	for i := 0; i < numRequests; i++ {
		mockPrivateKeyRepo.EXPECT().GetPrivateKey(ctx, mockContractAddress).Return(privateKey, nil)
		mockSigningService.EXPECT().SignTxn(privateKey, args).Return(hash, signature, nil)
	}
	batchResp, signTxnErr := signerServer.BatchSignTxn(ctx, &pb.BatchSignatureRequest{SignatureRequests: requests})
	assert.Nil(t, signTxnErr)
	assert.Equal(t, numRequests, len(batchResp.SignatureResponses))
	for _, resp := range batchResp.SignatureResponses {
		assert.Equal(t, hash, resp.Hash)
		assert.Equal(t, signature, resp.Signature)
	}
}

func TestSignerRPCService_BatchSignTxn_ErrorWithOneRequest(t *testing.T) {
	mockSigningService, mockPrivateKeyRepo, signerServer, ctx := newTestingSignerServer(t)
	defer signerServer.Server.GracefulStop()

	args := [][]byte{[]byte("some arguments or something")}
	requests := []*pb.SignatureRequest{}
	numRequests := 10
	for i := 0; i < numRequests; i++ {
		requests = append(requests, &pb.SignatureRequest{
			Args:            args,
			ContractAddress: mockContractAddress,
		})
	}
	err := errors.New("testing error for testing")
	hash := "0xI'mAHash"
	signature := "0xI'mASignature"
	mockPrivateKeyRepo.EXPECT().GetPrivateKey(ctx, mockContractAddress).Return(privateKey, nil)
	mockSigningService.EXPECT().SignTxn(privateKey, args).Return("", "", err)
	for i := 0; i < numRequests-1; i++ {
		mockPrivateKeyRepo.EXPECT().GetPrivateKey(ctx, mockContractAddress).Return(privateKey, nil)
		mockSigningService.EXPECT().SignTxn(privateKey, args).Return(hash, signature, nil)
	}
	batchResp, signTxnErr := signerServer.BatchSignTxn(ctx, &pb.BatchSignatureRequest{SignatureRequests: requests})
	assert.Nil(t, signTxnErr)
	assert.Equal(t, numRequests-1, len(batchResp.SignatureResponses))
	for _, resp := range batchResp.SignatureResponses {
		assert.Equal(t, hash, resp.Hash)
		assert.Equal(t, signature, resp.Signature)
	}
}

func TestSignerRPCService_BatchSignTxn_ErrWithAllTxns(t *testing.T) {
	_, mockPrivateKeyRepo, signerServer, ctx := newTestingSignerServer(t)
	defer signerServer.Server.GracefulStop()

	args := [][]byte{[]byte("some arguments or something")}
	requests := []*pb.SignatureRequest{}
	numRequests := 10
	for i := 0; i < numRequests; i++ {
		requests = append(requests, &pb.SignatureRequest{
			Args:            args,
			ContractAddress: mockContractAddress,
		})
	}
	err := errors.New("testing error for testing")
	for i := 0; i < numRequests; i++ {
		mockPrivateKeyRepo.EXPECT().GetPrivateKey(ctx, mockContractAddress).Return("", err)
	}
	batchResp, signTxnErr := signerServer.BatchSignTxn(ctx, &pb.BatchSignatureRequest{SignatureRequests: requests})
	assert.Equal(t, errors.New("none of the signing requests could be fulfilled"), signTxnErr)
	assert.Nil(t, batchResp)
}
