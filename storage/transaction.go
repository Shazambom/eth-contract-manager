package storage

import (
	"bytes"
	"compress/gzip"
	pb "contract-service/proto"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/math"
	"io/ioutil"
)

type Transaction struct {
	ContractAddress string `json:"contract_address"`
	ABIPackedTxn []byte `json:"abi_packed_txn"`
	UserAddress string `json:"user_address"`
	Hash string `json:"hash"`
	IsComplete bool `json:"is_complete"`
	Value	string `json:"value"`
}

func NewTransaction(contractAddress, userAddress, hash string, txn []byte, value string) (*Transaction, error){
	if _, ok := math.ParseBig256(value); !ok {
		return nil, errors.New("Error parsing value from string " + value + " is an invalid amount of wei")
	}
	return &Transaction{
		ContractAddress: contractAddress,
		ABIPackedTxn: txn,
		UserAddress: userAddress,
		Hash: hash,
		IsComplete: false,
		Value: value,
	}, nil
}

func (token *Transaction) FromRPC(txn *pb.Transaction) error {
	if _, ok := math.ParseBig256(txn.Value); !ok {
		return errors.New("Error parsing value from string " + txn.Value + " is an invalid amount of wei")
	}
	token.Hash = txn.Hash
	token.ABIPackedTxn = txn.PackedArgs
	token.ContractAddress = txn.ContractAddress
	token.UserAddress = txn.UserAddress
	token.IsComplete = txn.IsComplete
	token.Value = txn.Value
	return nil
}

func (token *Transaction) ToRPC() *pb.Transaction {
	return &pb.Transaction{
		PackedArgs: token.ABIPackedTxn,
		Hash:       token.Hash,
		ContractAddress: token.ContractAddress,
		UserAddress: token.UserAddress,
		IsComplete: token.IsComplete,
		Value: token.Value,
	}
}


func (token *Transaction) ToString() (string, error) {
	byteArr, err := json.Marshal(token)
	return string(byteArr), err
}

func (token *Transaction) Gzip() (string, error) {
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

func (token *Transaction) UnZip(payload []byte) error {
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
	var tok Transaction
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