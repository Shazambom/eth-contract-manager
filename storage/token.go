package storage

import (
	"bytes"
	"compress/gzip"
	pb "contract-service/proto"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Token struct {
	ContractAddress string `json:"contract_address"`
	ABIPackedTxn []byte `json:"abi_packed_txn"`
	UserAddress string `json:"user_address"`
	Hash string `json:"hash"`
	IsComplete bool `json:"is_complete"`
	//TODO Maybe refactor this to just be a string in wei? This would remove the need for the DynamoTransaction struct and reduce complexity
	// Use BigInt and string and just represent it in wei. Check contract_manager_service.go
	Value	float32 `json:"value"`
}

func NewToken(contractAddress, userAddress, hash string, txn []byte, value float32) *Token {
	return &Token{
		ContractAddress: contractAddress,
		ABIPackedTxn: txn,
		UserAddress: userAddress,
		Hash: hash,
		IsComplete: false,
		Value: value,
	}
}

type DynamoTransaction struct {
	ContractAddress string `json:"contract_address"`
	ABIPackedTxn []byte `json:"abi_packed_txn"`
	UserAddress string `json:"user_address"`
	Hash string `json:"hash"`
	IsComplete bool `json:"is_complete"`
	Value	string `json:"value"`
}

func (dt *DynamoTransaction) ToToken() (*Token, error) {
	value, err := strconv.ParseFloat(dt.Value, 32)
	if err != nil {
		return nil, err
	}
	return &Token{
		ContractAddress: dt.ContractAddress,
		ABIPackedTxn:    dt.ABIPackedTxn,
		UserAddress:     dt.UserAddress,
		Hash:            dt.Hash,
		IsComplete:      dt.IsComplete,
		Value:           float32(value),
	}, nil
}

func (token *Token) ToDynamo() *DynamoTransaction {
	return &DynamoTransaction{
		ContractAddress: token.ContractAddress,
		ABIPackedTxn:    token.ABIPackedTxn,
		UserAddress:     token.UserAddress,
		Hash:            token.Hash,
		IsComplete:      token.IsComplete,
		Value:           fmt.Sprintf("%f", token.Value),
	}
}

func (token *Token) FromRPC(txn *pb.Transaction) {
	token.Hash = txn.Hash
	token.ABIPackedTxn = txn.PackedArgs
	token.ContractAddress = txn.ContractAddress
	token.UserAddress = txn.UserAddress
	token.IsComplete = txn.IsComplete
	token.Value = txn.Value
}

func (token *Token) ToRPC() *pb.Transaction {
	return &pb.Transaction{
		PackedArgs: token.ABIPackedTxn,
		Hash:       token.Hash,
		ContractAddress: token.ContractAddress,
		UserAddress: token.UserAddress,
		IsComplete: token.IsComplete,
		Value: token.Value,
	}
}


func (token *Token) ToString() (string, error) {
	byteArr, err := json.Marshal(token)
	return string(byteArr), err
}

func (token *Token) Gzip() (string, error) {
	raw, respMarshalErr := json.Marshal(token)
	if respMarshalErr != nil {
		return "", respMarshalErr
	}
	buff := bytes.Buffer{}
	gz := gzip.NewWriter(&buff)
	if _, gzErr := gz.Write(raw); gzErr != nil {
		fmt.Println(gzErr.Error())
		return "", gzErr
	}
	if flshErr := gz.Flush(); flshErr != nil {
		fmt.Println(flshErr.Error())
		return "", flshErr
	}
	if closeErr := gz.Close(); closeErr != nil {
		fmt.Println(closeErr.Error())
		return "", closeErr
	}
	return base64.StdEncoding.EncodeToString(buff.Bytes()), nil
}

func (token *Token) UnZip(payload []byte) error {
	data, decodeErr := base64.StdEncoding.DecodeString(string(payload))
	if decodeErr != nil {
		return decodeErr
	}
	gz, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return err
	}
	unzipped, unzipErr := ioutil.ReadAll(gz)
	if unzipErr != nil {
		return unzipErr
	}
	var tok Token
	marshalErr := json.Unmarshal(unzipped, &tok)
	if marshalErr != nil {
		return marshalErr
	}
	token.ContractAddress = tok.ContractAddress
	token.UserAddress = tok.UserAddress
	token.Value = tok.Value
	token.Hash = tok.Hash
	token.ABIPackedTxn = tok.ABIPackedTxn
	return nil
}