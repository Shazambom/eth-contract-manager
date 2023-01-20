package signing

import (
	"context"
	"contract-service/mocks"
	pb "contract-service/proto"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
)

func newTestingVerifierServer(t *testing.T) (*mocks.MockSigningService, *VerifierRPCService, context.Context) {
	ctrl := gomock.NewController(t)
	mockSigningService := mocks.NewMockSigningService(ctrl)
	verifierServer, newServerErr := NewVerifierServer(getTestPort(), []grpc.ServerOption{grpc.EmptyServerOption{}}, mockSigningService)
	assert.Nil(t, newServerErr)
	return mockSigningService, verifierServer, context.Background()
}

func TestNewVerifierServer(t *testing.T) {
	_, verifierService, _ := newTestingVerifierServer(t)
	defer verifierService.Server.GracefulStop()
	assert.IsType(t, &VerifierRPCService{}, verifierService)
}

func TestVerifierRPCService_Verify(t *testing.T) {
	mockSigningService, verifierService, ctx := newTestingVerifierServer(t)
	defer verifierService.Server.GracefulStop()

	req := &pb.SignatureVerificationRequest{
		Message:   "a message",
		Signature: "a signature",
		Address:   "an address",
	}

	mockSigningService.EXPECT().Verify(req.Message, req.Signature, req.Address).Return(nil)

	resp, err := verifierService.Verify(ctx, req)
	assert.Nil(t, err)
	assert.True(t, resp.Success)
}

func TestVerifierRPCService_Verify_InvalidSignature(t *testing.T) {
	mockSigningService, verifierService, ctx := newTestingVerifierServer(t)
	defer verifierService.Server.GracefulStop()

	req := &pb.SignatureVerificationRequest{
		Message:   "a message",
		Signature: "a signature",
		Address:   "an address",
	}

	mockSigningService.EXPECT().Verify(req.Message, req.Signature, req.Address).Return(errors.New("some error"))

	resp, err := verifierService.Verify(ctx, req)
	assert.Nil(t, err)
	assert.False(t, resp.Success)
}
