package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type SignedRequest struct {
	Event Event `json:"event"`
	Signature string `json:"signature"`
	Hash string `json:"hash"`
	ContractAddress string `json:"contract_address"`
	ABI []string `json:"abi"`
	QueueNumber int64 `json:"queue_number"`
}

func NewSignedRequest(event Event, signature , hash , contractAddress string, ABI []string, queueNumber int64) SignedRequest{
	return SignedRequest{
		Event: event,
		Signature: signature,
		Hash: hash,
		ContractAddress: contractAddress,
		ABI: ABI,
		QueueNumber: queueNumber,
	}
}

func (sr *SignedRequest) Gzip() (string, error) {
	raw, respMarshalErr := json.Marshal(sr)
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
