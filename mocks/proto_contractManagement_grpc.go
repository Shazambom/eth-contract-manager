// Code generated by MockGen. DO NOT EDIT.
// Source: ./proto/contractManagement_grpc.pb.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	pb "bitbucket.org/artie_inc/contract-service/proto"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockTransactionServiceClient is a mock of TransactionServiceClient interface.
type MockTransactionServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionServiceClientMockRecorder
}

// MockTransactionServiceClientMockRecorder is the mock recorder for MockTransactionServiceClient.
type MockTransactionServiceClientMockRecorder struct {
	mock *MockTransactionServiceClient
}

// NewMockTransactionServiceClient creates a new mock instance.
func NewMockTransactionServiceClient(ctrl *gomock.Controller) *MockTransactionServiceClient {
	mock := &MockTransactionServiceClient{ctrl: ctrl}
	mock.recorder = &MockTransactionServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionServiceClient) EXPECT() *MockTransactionServiceClientMockRecorder {
	return m.recorder
}

// CompleteTransaction mocks base method.
func (m *MockTransactionServiceClient) CompleteTransaction(ctx context.Context, in *pb.KeyTransactionRequest, opts ...grpc.CallOption) (*pb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CompleteTransaction", varargs...)
	ret0, _ := ret[0].(*pb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompleteTransaction indicates an expected call of CompleteTransaction.
func (mr *MockTransactionServiceClientMockRecorder) CompleteTransaction(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteTransaction", reflect.TypeOf((*MockTransactionServiceClient)(nil).CompleteTransaction), varargs...)
}

// ConstructTransaction mocks base method.
func (m *MockTransactionServiceClient) ConstructTransaction(ctx context.Context, in *pb.TransactionRequest, opts ...grpc.CallOption) (*pb.Transaction, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ConstructTransaction", varargs...)
	ret0, _ := ret[0].(*pb.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConstructTransaction indicates an expected call of ConstructTransaction.
func (mr *MockTransactionServiceClientMockRecorder) ConstructTransaction(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConstructTransaction", reflect.TypeOf((*MockTransactionServiceClient)(nil).ConstructTransaction), varargs...)
}

// DeleteTransaction mocks base method.
func (m *MockTransactionServiceClient) DeleteTransaction(ctx context.Context, in *pb.KeyTransactionRequest, opts ...grpc.CallOption) (*pb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteTransaction", varargs...)
	ret0, _ := ret[0].(*pb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTransaction indicates an expected call of DeleteTransaction.
func (mr *MockTransactionServiceClientMockRecorder) DeleteTransaction(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTransaction", reflect.TypeOf((*MockTransactionServiceClient)(nil).DeleteTransaction), varargs...)
}

// GetAllTransactions mocks base method.
func (m *MockTransactionServiceClient) GetAllTransactions(ctx context.Context, in *pb.Address, opts ...grpc.CallOption) (*pb.Transactions, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAllTransactions", varargs...)
	ret0, _ := ret[0].(*pb.Transactions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTransactions indicates an expected call of GetAllTransactions.
func (mr *MockTransactionServiceClientMockRecorder) GetAllTransactions(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTransactions", reflect.TypeOf((*MockTransactionServiceClient)(nil).GetAllTransactions), varargs...)
}

// GetContract mocks base method.
func (m *MockTransactionServiceClient) GetContract(ctx context.Context, in *pb.Address, opts ...grpc.CallOption) (*pb.Contract, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetContract", varargs...)
	ret0, _ := ret[0].(*pb.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContract indicates an expected call of GetContract.
func (mr *MockTransactionServiceClientMockRecorder) GetContract(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContract", reflect.TypeOf((*MockTransactionServiceClient)(nil).GetContract), varargs...)
}

// GetTransactions mocks base method.
func (m *MockTransactionServiceClient) GetTransactions(ctx context.Context, in *pb.Address, opts ...grpc.CallOption) (*pb.Transactions, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTransactions", varargs...)
	ret0, _ := ret[0].(*pb.Transactions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactions indicates an expected call of GetTransactions.
func (mr *MockTransactionServiceClientMockRecorder) GetTransactions(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactions", reflect.TypeOf((*MockTransactionServiceClient)(nil).GetTransactions), varargs...)
}

// MockTransactionServiceServer is a mock of TransactionServiceServer interface.
type MockTransactionServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionServiceServerMockRecorder
}

// MockTransactionServiceServerMockRecorder is the mock recorder for MockTransactionServiceServer.
type MockTransactionServiceServerMockRecorder struct {
	mock *MockTransactionServiceServer
}

// NewMockTransactionServiceServer creates a new mock instance.
func NewMockTransactionServiceServer(ctrl *gomock.Controller) *MockTransactionServiceServer {
	mock := &MockTransactionServiceServer{ctrl: ctrl}
	mock.recorder = &MockTransactionServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionServiceServer) EXPECT() *MockTransactionServiceServerMockRecorder {
	return m.recorder
}

// CompleteTransaction mocks base method.
func (m *MockTransactionServiceServer) CompleteTransaction(arg0 context.Context, arg1 *pb.KeyTransactionRequest) (*pb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompleteTransaction", arg0, arg1)
	ret0, _ := ret[0].(*pb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompleteTransaction indicates an expected call of CompleteTransaction.
func (mr *MockTransactionServiceServerMockRecorder) CompleteTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteTransaction", reflect.TypeOf((*MockTransactionServiceServer)(nil).CompleteTransaction), arg0, arg1)
}

// ConstructTransaction mocks base method.
func (m *MockTransactionServiceServer) ConstructTransaction(arg0 context.Context, arg1 *pb.TransactionRequest) (*pb.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConstructTransaction", arg0, arg1)
	ret0, _ := ret[0].(*pb.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConstructTransaction indicates an expected call of ConstructTransaction.
func (mr *MockTransactionServiceServerMockRecorder) ConstructTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConstructTransaction", reflect.TypeOf((*MockTransactionServiceServer)(nil).ConstructTransaction), arg0, arg1)
}

// DeleteTransaction mocks base method.
func (m *MockTransactionServiceServer) DeleteTransaction(arg0 context.Context, arg1 *pb.KeyTransactionRequest) (*pb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTransaction", arg0, arg1)
	ret0, _ := ret[0].(*pb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTransaction indicates an expected call of DeleteTransaction.
func (mr *MockTransactionServiceServerMockRecorder) DeleteTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTransaction", reflect.TypeOf((*MockTransactionServiceServer)(nil).DeleteTransaction), arg0, arg1)
}

// GetAllTransactions mocks base method.
func (m *MockTransactionServiceServer) GetAllTransactions(arg0 context.Context, arg1 *pb.Address) (*pb.Transactions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTransactions", arg0, arg1)
	ret0, _ := ret[0].(*pb.Transactions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTransactions indicates an expected call of GetAllTransactions.
func (mr *MockTransactionServiceServerMockRecorder) GetAllTransactions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTransactions", reflect.TypeOf((*MockTransactionServiceServer)(nil).GetAllTransactions), arg0, arg1)
}

// GetContract mocks base method.
func (m *MockTransactionServiceServer) GetContract(arg0 context.Context, arg1 *pb.Address) (*pb.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContract", arg0, arg1)
	ret0, _ := ret[0].(*pb.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContract indicates an expected call of GetContract.
func (mr *MockTransactionServiceServerMockRecorder) GetContract(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContract", reflect.TypeOf((*MockTransactionServiceServer)(nil).GetContract), arg0, arg1)
}

// GetTransactions mocks base method.
func (m *MockTransactionServiceServer) GetTransactions(arg0 context.Context, arg1 *pb.Address) (*pb.Transactions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactions", arg0, arg1)
	ret0, _ := ret[0].(*pb.Transactions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactions indicates an expected call of GetTransactions.
func (mr *MockTransactionServiceServerMockRecorder) GetTransactions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactions", reflect.TypeOf((*MockTransactionServiceServer)(nil).GetTransactions), arg0, arg1)
}

// mustEmbedUnimplementedTransactionServiceServer mocks base method.
func (m *MockTransactionServiceServer) mustEmbedUnimplementedTransactionServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedTransactionServiceServer")
}

// mustEmbedUnimplementedTransactionServiceServer indicates an expected call of mustEmbedUnimplementedTransactionServiceServer.
func (mr *MockTransactionServiceServerMockRecorder) mustEmbedUnimplementedTransactionServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedTransactionServiceServer", reflect.TypeOf((*MockTransactionServiceServer)(nil).mustEmbedUnimplementedTransactionServiceServer))
}

// MockUnsafeTransactionServiceServer is a mock of UnsafeTransactionServiceServer interface.
type MockUnsafeTransactionServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeTransactionServiceServerMockRecorder
}

// MockUnsafeTransactionServiceServerMockRecorder is the mock recorder for MockUnsafeTransactionServiceServer.
type MockUnsafeTransactionServiceServerMockRecorder struct {
	mock *MockUnsafeTransactionServiceServer
}

// NewMockUnsafeTransactionServiceServer creates a new mock instance.
func NewMockUnsafeTransactionServiceServer(ctrl *gomock.Controller) *MockUnsafeTransactionServiceServer {
	mock := &MockUnsafeTransactionServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeTransactionServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeTransactionServiceServer) EXPECT() *MockUnsafeTransactionServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedTransactionServiceServer mocks base method.
func (m *MockUnsafeTransactionServiceServer) mustEmbedUnimplementedTransactionServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedTransactionServiceServer")
}

// mustEmbedUnimplementedTransactionServiceServer indicates an expected call of mustEmbedUnimplementedTransactionServiceServer.
func (mr *MockUnsafeTransactionServiceServerMockRecorder) mustEmbedUnimplementedTransactionServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedTransactionServiceServer", reflect.TypeOf((*MockUnsafeTransactionServiceServer)(nil).mustEmbedUnimplementedTransactionServiceServer))
}

// MockContractManagementClient is a mock of ContractManagementClient interface.
type MockContractManagementClient struct {
	ctrl     *gomock.Controller
	recorder *MockContractManagementClientMockRecorder
}

// MockContractManagementClientMockRecorder is the mock recorder for MockContractManagementClient.
type MockContractManagementClientMockRecorder struct {
	mock *MockContractManagementClient
}

// NewMockContractManagementClient creates a new mock instance.
func NewMockContractManagementClient(ctrl *gomock.Controller) *MockContractManagementClient {
	mock := &MockContractManagementClient{ctrl: ctrl}
	mock.recorder = &MockContractManagementClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContractManagementClient) EXPECT() *MockContractManagementClientMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockContractManagementClient) Delete(ctx context.Context, in *pb.AddressOwner, opts ...grpc.CallOption) (*pb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*pb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockContractManagementClientMockRecorder) Delete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockContractManagementClient)(nil).Delete), varargs...)
}

// Get mocks base method.
func (m *MockContractManagementClient) Get(ctx context.Context, in *pb.Address, opts ...grpc.CallOption) (*pb.Contract, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*pb.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockContractManagementClientMockRecorder) Get(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockContractManagementClient)(nil).Get), varargs...)
}

// List mocks base method.
func (m *MockContractManagementClient) List(ctx context.Context, in *pb.Owner, opts ...grpc.CallOption) (*pb.Contracts, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*pb.Contracts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockContractManagementClientMockRecorder) List(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockContractManagementClient)(nil).List), varargs...)
}

// Store mocks base method.
func (m *MockContractManagementClient) Store(ctx context.Context, in *pb.Contract, opts ...grpc.CallOption) (*pb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Store", varargs...)
	ret0, _ := ret[0].(*pb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Store indicates an expected call of Store.
func (mr *MockContractManagementClientMockRecorder) Store(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockContractManagementClient)(nil).Store), varargs...)
}

// MockContractManagementServer is a mock of ContractManagementServer interface.
type MockContractManagementServer struct {
	ctrl     *gomock.Controller
	recorder *MockContractManagementServerMockRecorder
}

// MockContractManagementServerMockRecorder is the mock recorder for MockContractManagementServer.
type MockContractManagementServerMockRecorder struct {
	mock *MockContractManagementServer
}

// NewMockContractManagementServer creates a new mock instance.
func NewMockContractManagementServer(ctrl *gomock.Controller) *MockContractManagementServer {
	mock := &MockContractManagementServer{ctrl: ctrl}
	mock.recorder = &MockContractManagementServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContractManagementServer) EXPECT() *MockContractManagementServerMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockContractManagementServer) Delete(arg0 context.Context, arg1 *pb.AddressOwner) (*pb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*pb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockContractManagementServerMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockContractManagementServer)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockContractManagementServer) Get(arg0 context.Context, arg1 *pb.Address) (*pb.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*pb.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockContractManagementServerMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockContractManagementServer)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockContractManagementServer) List(arg0 context.Context, arg1 *pb.Owner) (*pb.Contracts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*pb.Contracts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockContractManagementServerMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockContractManagementServer)(nil).List), arg0, arg1)
}

// Store mocks base method.
func (m *MockContractManagementServer) Store(arg0 context.Context, arg1 *pb.Contract) (*pb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0, arg1)
	ret0, _ := ret[0].(*pb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Store indicates an expected call of Store.
func (mr *MockContractManagementServerMockRecorder) Store(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockContractManagementServer)(nil).Store), arg0, arg1)
}

// mustEmbedUnimplementedContractManagementServer mocks base method.
func (m *MockContractManagementServer) mustEmbedUnimplementedContractManagementServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedContractManagementServer")
}

// mustEmbedUnimplementedContractManagementServer indicates an expected call of mustEmbedUnimplementedContractManagementServer.
func (mr *MockContractManagementServerMockRecorder) mustEmbedUnimplementedContractManagementServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedContractManagementServer", reflect.TypeOf((*MockContractManagementServer)(nil).mustEmbedUnimplementedContractManagementServer))
}

// MockUnsafeContractManagementServer is a mock of UnsafeContractManagementServer interface.
type MockUnsafeContractManagementServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeContractManagementServerMockRecorder
}

// MockUnsafeContractManagementServerMockRecorder is the mock recorder for MockUnsafeContractManagementServer.
type MockUnsafeContractManagementServerMockRecorder struct {
	mock *MockUnsafeContractManagementServer
}

// NewMockUnsafeContractManagementServer creates a new mock instance.
func NewMockUnsafeContractManagementServer(ctrl *gomock.Controller) *MockUnsafeContractManagementServer {
	mock := &MockUnsafeContractManagementServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeContractManagementServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeContractManagementServer) EXPECT() *MockUnsafeContractManagementServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedContractManagementServer mocks base method.
func (m *MockUnsafeContractManagementServer) mustEmbedUnimplementedContractManagementServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedContractManagementServer")
}

// mustEmbedUnimplementedContractManagementServer indicates an expected call of mustEmbedUnimplementedContractManagementServer.
func (mr *MockUnsafeContractManagementServerMockRecorder) mustEmbedUnimplementedContractManagementServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedContractManagementServer", reflect.TypeOf((*MockUnsafeContractManagementServer)(nil).mustEmbedUnimplementedContractManagementServer))
}
