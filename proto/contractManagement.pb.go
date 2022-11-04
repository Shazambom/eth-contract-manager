// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/contractManagement.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contractManagement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contractManagement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_proto_contractManagement_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type Owner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *Owner) Reset() {
	*x = Owner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contractManagement_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Owner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Owner) ProtoMessage() {}

func (x *Owner) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contractManagement_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Owner.ProtoReflect.Descriptor instead.
func (*Owner) Descriptor() ([]byte, []int) {
	return file_proto_contractManagement_proto_rawDescGZIP(), []int{1}
}

func (x *Owner) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

type AddressOwner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Owner   string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *AddressOwner) Reset() {
	*x = AddressOwner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contractManagement_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressOwner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressOwner) ProtoMessage() {}

func (x *AddressOwner) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contractManagement_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressOwner.ProtoReflect.Descriptor instead.
func (*AddressOwner) Descriptor() ([]byte, []int) {
	return file_proto_contractManagement_proto_rawDescGZIP(), []int{2}
}

func (x *AddressOwner) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AddressOwner) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

type Argument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Argument) Reset() {
	*x = Argument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contractManagement_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Argument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Argument) ProtoMessage() {}

func (x *Argument) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contractManagement_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Argument.ProtoReflect.Descriptor instead.
func (*Argument) Descriptor() ([]byte, []int) {
	return file_proto_contractManagement_proto_rawDescGZIP(), []int{3}
}

func (x *Argument) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Argument) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Function struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Arguments []*Argument `protobuf:"bytes,1,rep,name=arguments,proto3" json:"arguments,omitempty"`
}

func (x *Function) Reset() {
	*x = Function{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contractManagement_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Function) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Function) ProtoMessage() {}

func (x *Function) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contractManagement_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Function.ProtoReflect.Descriptor instead.
func (*Function) Descriptor() ([]byte, []int) {
	return file_proto_contractManagement_proto_rawDescGZIP(), []int{4}
}

func (x *Function) GetArguments() []*Argument {
	if x != nil {
		return x.Arguments
	}
	return nil
}

type Functions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Functions map[string]*Function `protobuf:"bytes,1,rep,name=functions,proto3" json:"functions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Functions) Reset() {
	*x = Functions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contractManagement_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Functions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Functions) ProtoMessage() {}

func (x *Functions) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contractManagement_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Functions.ProtoReflect.Descriptor instead.
func (*Functions) Descriptor() ([]byte, []int) {
	return file_proto_contractManagement_proto_rawDescGZIP(), []int{5}
}

func (x *Functions) GetFunctions() map[string]*Function {
	if x != nil {
		return x.Functions
	}
	return nil
}

type Contract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address           string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Abi               string     `protobuf:"bytes,2,opt,name=abi,proto3" json:"abi,omitempty"`
	HashableFunctions *Functions `protobuf:"bytes,3,opt,name=hashable_functions,json=hashableFunctions,proto3" json:"hashable_functions,omitempty"`
	MaxMintable       int64      `protobuf:"varint,4,opt,name=max_mintable,json=maxMintable,proto3" json:"max_mintable,omitempty"`
	MaxIncrement      int64      `protobuf:"varint,5,opt,name=max_increment,json=maxIncrement,proto3" json:"max_increment,omitempty"`
	Owner             string     `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *Contract) Reset() {
	*x = Contract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contractManagement_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contract) ProtoMessage() {}

func (x *Contract) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contractManagement_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contract.ProtoReflect.Descriptor instead.
func (*Contract) Descriptor() ([]byte, []int) {
	return file_proto_contractManagement_proto_rawDescGZIP(), []int{6}
}

func (x *Contract) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Contract) GetAbi() string {
	if x != nil {
		return x.Abi
	}
	return ""
}

func (x *Contract) GetHashableFunctions() *Functions {
	if x != nil {
		return x.HashableFunctions
	}
	return nil
}

func (x *Contract) GetMaxMintable() int64 {
	if x != nil {
		return x.MaxMintable
	}
	return 0
}

func (x *Contract) GetMaxIncrement() int64 {
	if x != nil {
		return x.MaxIncrement
	}
	return 0
}

func (x *Contract) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

type TransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageSender string    `protobuf:"bytes,1,opt,name=message_sender,json=messageSender,proto3" json:"message_sender,omitempty"`
	FunctionName  string    `protobuf:"bytes,2,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	NumRequested  int64     `protobuf:"varint,3,opt,name=num_requested,json=numRequested,proto3" json:"num_requested,omitempty"`
	Args          [][]byte  `protobuf:"bytes,4,rep,name=args,proto3" json:"args,omitempty"`
	Contract      *Contract `protobuf:"bytes,5,opt,name=contract,proto3" json:"contract,omitempty"`
}

func (x *TransactionRequest) Reset() {
	*x = TransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contractManagement_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionRequest) ProtoMessage() {}

func (x *TransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contractManagement_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionRequest.ProtoReflect.Descriptor instead.
func (*TransactionRequest) Descriptor() ([]byte, []int) {
	return file_proto_contractManagement_proto_rawDescGZIP(), []int{7}
}

func (x *TransactionRequest) GetMessageSender() string {
	if x != nil {
		return x.MessageSender
	}
	return ""
}

func (x *TransactionRequest) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

func (x *TransactionRequest) GetNumRequested() int64 {
	if x != nil {
		return x.NumRequested
	}
	return 0
}

func (x *TransactionRequest) GetArgs() [][]byte {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *TransactionRequest) GetContract() *Contract {
	if x != nil {
		return x.Contract
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err string `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contractManagement_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contractManagement_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_proto_contractManagement_proto_rawDescGZIP(), []int{8}
}

func (x *Error) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type Contracts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contracts []*Contract `protobuf:"bytes,1,rep,name=contracts,proto3" json:"contracts,omitempty"`
}

func (x *Contracts) Reset() {
	*x = Contracts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contractManagement_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contracts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contracts) ProtoMessage() {}

func (x *Contracts) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contractManagement_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contracts.ProtoReflect.Descriptor instead.
func (*Contracts) Descriptor() ([]byte, []int) {
	return file_proto_contractManagement_proto_rawDescGZIP(), []int{9}
}

func (x *Contracts) GetContracts() []*Contract {
	if x != nil {
		return x.Contracts
	}
	return nil
}

var File_proto_contractManagement_proto protoreflect.FileDescriptor

var file_proto_contractManagement_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x23, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x1d, 0x0a, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x22, 0x3e, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x22, 0x32, 0x0a, 0x08, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x33, 0x0a, 0x08, 0x46, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x8d, 0x01,
	0x0a, 0x09, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x37, 0x0a, 0x09, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x47, 0x0a, 0x0e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xcf, 0x01,
	0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x62, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x61, 0x62, 0x69, 0x12, 0x39, 0x0a, 0x12, 0x68, 0x61, 0x73, 0x68, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x11,
	0x68, 0x61, 0x73, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f, 0x6d, 0x69, 0x6e, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x4d, 0x69, 0x6e, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x6e, 0x63, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6d, 0x61, 0x78,
	0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22,
	0xc0, 0x01, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x23, 0x0a,
	0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x75, 0x6d, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6e, 0x75, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x25, 0x0a, 0x08, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x22, 0x19, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x65,
	0x72, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x34, 0x0a,
	0x09, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x12, 0x27, 0x0a, 0x09, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x73, 0x32, 0x71, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x08, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x1a, 0x09, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x22, 0x00, 0x12,
	0x35, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x00, 0x32, 0x91, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x08, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x09,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x22, 0x00, 0x12, 0x1c, 0x0a, 0x05, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x12, 0x09, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x1a,
	0x06, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x00, 0x12, 0x21, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x1a, 0x06, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x00, 0x12, 0x1c, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x06, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x1a, 0x0a, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f,
	0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_contractManagement_proto_rawDescOnce sync.Once
	file_proto_contractManagement_proto_rawDescData = file_proto_contractManagement_proto_rawDesc
)

func file_proto_contractManagement_proto_rawDescGZIP() []byte {
	file_proto_contractManagement_proto_rawDescOnce.Do(func() {
		file_proto_contractManagement_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_contractManagement_proto_rawDescData)
	})
	return file_proto_contractManagement_proto_rawDescData
}

var file_proto_contractManagement_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_contractManagement_proto_goTypes = []interface{}{
	(*Address)(nil),            // 0: Address
	(*Owner)(nil),              // 1: Owner
	(*AddressOwner)(nil),       // 2: AddressOwner
	(*Argument)(nil),           // 3: Argument
	(*Function)(nil),           // 4: Function
	(*Functions)(nil),          // 5: Functions
	(*Contract)(nil),           // 6: Contract
	(*TransactionRequest)(nil), // 7: TransactionRequest
	(*Error)(nil),              // 8: Error
	(*Contracts)(nil),          // 9: Contracts
	nil,                        // 10: Functions.FunctionsEntry
}
var file_proto_contractManagement_proto_depIdxs = []int32{
	3,  // 0: Function.arguments:type_name -> Argument
	10, // 1: Functions.functions:type_name -> Functions.FunctionsEntry
	5,  // 2: Contract.hashable_functions:type_name -> Functions
	6,  // 3: TransactionRequest.contract:type_name -> Contract
	6,  // 4: Contracts.contracts:type_name -> Contract
	4,  // 5: Functions.FunctionsEntry.value:type_name -> Function
	0,  // 6: TransactionService.GetContract:input_type -> Address
	7,  // 7: TransactionService.ConstructTransaction:input_type -> TransactionRequest
	0,  // 8: ContractManagement.Get:input_type -> Address
	6,  // 9: ContractManagement.Store:input_type -> Contract
	2,  // 10: ContractManagement.Delete:input_type -> AddressOwner
	1,  // 11: ContractManagement.List:input_type -> Owner
	6,  // 12: TransactionService.GetContract:output_type -> Contract
	8,  // 13: TransactionService.ConstructTransaction:output_type -> Error
	6,  // 14: ContractManagement.Get:output_type -> Contract
	8,  // 15: ContractManagement.Store:output_type -> Error
	8,  // 16: ContractManagement.Delete:output_type -> Error
	9,  // 17: ContractManagement.List:output_type -> Contracts
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_contractManagement_proto_init() }
func file_proto_contractManagement_proto_init() {
	if File_proto_contractManagement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_contractManagement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_contractManagement_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Owner); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_contractManagement_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressOwner); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_contractManagement_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Argument); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_contractManagement_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Function); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_contractManagement_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Functions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_contractManagement_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contract); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_contractManagement_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_contractManagement_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_contractManagement_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contracts); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_contractManagement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_contractManagement_proto_goTypes,
		DependencyIndexes: file_proto_contractManagement_proto_depIdxs,
		MessageInfos:      file_proto_contractManagement_proto_msgTypes,
	}.Build()
	File_proto_contractManagement_proto = out.File
	file_proto_contractManagement_proto_rawDesc = nil
	file_proto_contractManagement_proto_goTypes = nil
	file_proto_contractManagement_proto_depIdxs = nil
}
