package storage

import (
	"bytes"
	"compress/gzip"
	pb "contract-service/proto"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Token struct {
	ContractAddress string `json:"contract_address"`
	ABIPackedTxn []byte `json:"abi_packed_txn"`
	ABI string `json:"abi"`
	UserAddress string `json:"user_address"`
	Hash string `json:"hash"`
	IsComplete bool `json:"is_complete"`
}

func NewToken(contractAddress, userAddress, hash string, abi string, txn []byte) *Token {
	return &Token{
		ContractAddress: contractAddress,
		ABIPackedTxn: txn,
		ABI: abi,
		UserAddress: userAddress,
		Hash: hash,
		IsComplete: false,
	}
}

func (token *Token) ToRPC() *pb.Transaction {
	return &pb.Transaction{
		Abi:        token.ABI,
		PackedArgs: token.ABIPackedTxn,
		Hash:       token.Hash,
		ContractAddress: token.ContractAddress,
		UserAddress: token.UserAddress,
		IsComplete: token.IsComplete,
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
	token.ABI = tok.ABI
	token.Hash = tok.Hash
	token.ABIPackedTxn = tok.ABIPackedTxn
	return nil
}