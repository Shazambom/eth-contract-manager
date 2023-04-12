package web

import (
	"bitbucket.org/artie_inc/contract-service/contracts"
	"bitbucket.org/artie_inc/contract-service/mocks"
	pb "bitbucket.org/artie_inc/contract-service/proto"
	"bitbucket.org/artie_inc/contract-service/storage"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func newTestTransactionAPI(t *testing.T) (*mocks.MockTransactionServiceClient, *TransactionAPI) {
	ctrl := gomock.NewController(t)
	mockTransactionServiceClient := mocks.NewMockTransactionServiceClient(ctrl)
	transactionClient := &contracts.TransactionClient{
		Connection: nil,
		Client:     mockTransactionServiceClient,
	}
	transactionAPI := NewTransactionAPI(transactionClient)
	return mockTransactionServiceClient, transactionAPI
}

func TestNewTransactionAPI(t *testing.T) {
	_, transactionAPI := newTestTransactionAPI(t)
	assert.IsType(t, &TransactionAPI{}, transactionAPI)
	assert.Implements(t, new(Servable), transactionAPI)
}

func callGetTransactionsFromAddress(port int, address string) (*http.Response, error) {
	return http.Get(fmt.Sprintf("http://localhost:%d/transactions/%s", port, address))
}
func TestTransactionAPI_GetTransactionsFromAddress(t *testing.T) {
	mockTransactionService, transactionAPI := newTestTransactionAPI(t)
	serverErr := make(chan string)
	port := getTestPort()
	transactionAPI.Serve(port, serverErr)
	userAddress := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	contractAddress := "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
	transaction, transactionErr := storage.NewTransaction(
		contractAddress,
		userAddress,
		"0xce11e286abab09c3ad05f1f9fff4daaf4f5139214a1f4746661018f0b855f075",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"0")
	assert.Nil(t, transactionErr)

	mockTransactionService.EXPECT().GetTransactions(gomock.Any(), &pb.Address{Address: userAddress}).Return(&pb.Transactions{Transactions: []*pb.Transaction{transaction.ToRPC()}}, nil)

	resp, reqErr := callGetTransactionsFromAddress(port, userAddress)
	assert.Nil(t, reqErr)

	respTxn := []*storage.Transaction{}
	body, readErr := io.ReadAll(resp.Body)
	assert.Nil(t, readErr)
	assert.Nil(t, json.Unmarshal(body, &respTxn))
	assert.Equal(t, transaction, respTxn[0])
}

func TestTransactionAPI_GetTransactionsFromAddress_ErrGettingTxn(t *testing.T) {
	mockTransactionService, transactionAPI := newTestTransactionAPI(t)
	serverErr := make(chan string)
	port := getTestPort()
	transactionAPI.Serve(port, serverErr)
	userAddress := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	someTransactionErr := errors.New("some transaction error")

	mockTransactionService.EXPECT().GetTransactions(gomock.Any(), &pb.Address{Address: userAddress}).Return(nil, someTransactionErr)

	resp, reqErr := callGetTransactionsFromAddress(port, userAddress)
	assert.Nil(t, reqErr)
	assert.Equal(t, "500 Internal Server Error", resp.Status)

}

func TestTransactionAPI_GetTransactionsFromAddress_InvalidTransaction(t *testing.T) {
	mockTransactionService, transactionAPI := newTestTransactionAPI(t)
	serverErr := make(chan string)
	port := getTestPort()
	transactionAPI.Serve(port, serverErr)
	userAddress := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	contractAddress := "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
	transaction := &storage.Transaction{
		ContractAddress: contractAddress,
		ABIPackedTxn:    []byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		UserAddress:     userAddress,
		Hash:            "0xce11e286abab09c3ad05f1f9fff4daaf4f5139214a1f4746661018f0b855f075",
		IsComplete:      false,
		Value:           "This isn't a valid value :)",
	}

	mockTransactionService.EXPECT().GetTransactions(gomock.Any(), &pb.Address{Address: userAddress}).Return(&pb.Transactions{Transactions: []*pb.Transaction{transaction.ToRPC()}}, nil)

	resp, reqErr := callGetTransactionsFromAddress(port, userAddress)
	assert.Nil(t, reqErr)
	assert.Equal(t, "500 Internal Server Error", resp.Status)
}
