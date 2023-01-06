package storage

import (
	pb "contract-service/proto"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"testing"
)


func TestNewContractRepository(t *testing.T) {
	assert.IsType(t, &ContractRepo{}, cr)
	inter := reflect.TypeOf((*ContractRepository)(nil)).Elem()
	assert.True(t, reflect.TypeOf(cr).Implements(inter))
}

func TestContractRepo_UpsertContract(t *testing.T) {
	contract := &Contract{
		Address:      "0xUpsertContract",
		ABI:          "[\n\t{\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"address payable\",\n\t\t\t\t\"name\": \"artieCharAddress\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"internalType\": \"address payable\",\n\t\t\t\t\"name\": \"withdrawAddress\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"signer\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"constructor\"\n\t},\n\t{\n\t\t\"anonymous\": false,\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"indexed\": true,\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"previousOwner\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"indexed\": true,\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"newOwner\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"OwnershipTransferred\",\n\t\t\"type\": \"event\"\n\t},\n\t{\n\t\t\"anonymous\": false,\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"indexed\": false,\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"to\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"indexed\": false,\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"amount\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"indexed\": false,\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"current\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"Season01Mint\",\n\t\t\"type\": \"event\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"MAX_TOKEN\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"PURCHASE_LIMIT\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"artie\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"contract Artie\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"current\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"bytes16\",\n\t\t\t\t\"name\": \"nonce\",\n\t\t\t\t\"type\": \"bytes16\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"numberOfTokens\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"transactionNumber\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"internalType\": \"bytes\",\n\t\t\t\t\"name\": \"signature\",\n\t\t\t\t\"type\": \"bytes\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"mint\",\n\t\t\"outputs\": [],\n\t\t\"stateMutability\": \"payable\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"owner\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"price\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"renounceOwnership\",\n\t\t\"outputs\": [],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"saleStarted\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"bool\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"bool\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"signer\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"setSignerAddress\",\n\t\t\"outputs\": [],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"address payable\",\n\t\t\t\t\"name\": \"givenWithdrawalAddress\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"setWithdrawalAddress\",\n\t\t\"outputs\": [],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"signingAddress\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"startSale\",\n\t\t\"outputs\": [],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"stopSale\",\n\t\t\"outputs\": [],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"newOwner\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"transferOwnership\",\n\t\t\"outputs\": [],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"bytes16\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"bytes16\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"usedNonces\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"bool\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"bool\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"withdrawEth\",\n\t\t\"outputs\": [],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"withdrawalAddress\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"address payable\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t}\n]",
		Functions:    map[string]Function{"mint": {Arguments: []Argument{
				{Name: "nonce", Type: "bytes16"},
				{Name: "msg.sender", Type: "address"},
				{Name: "numberOfTokens", Type: "uint256"},
				{Name: "transactionNumber", Type: "uint256"},
			}}},
		ContractOwner:        "Tester",
	}
	upsertErr := cr.UpsertContract(ctx, contract)
	assert.Nil(t, upsertErr)
}

func TestContractRepo_GetContract(t *testing.T) {

	contract := &Contract{
		Address:      "0xGetContract",
		ABI:          "abi :)",
		Functions:    map[string]Function{"mint": {Arguments: []Argument{
				{Name: "nonce", Type: "bytes16"},
				{Name: "msg.sender", Type: "address"},
				{Name: "numberOfTokens", Type: "uint256"},
				{Name: "transactionNumber", Type: "uint256"},
			}}},
		ContractOwner:        "Tester",
	}
	upsertErr := cr.UpsertContract(ctx, contract)
	assert.Nil(t, upsertErr)

	retrievedContract, getErr := cr.GetContract(ctx, contract.Address)
	assert.Nil(t, getErr)
	assert.Equal(t, contract, retrievedContract)
}

func TestContractRepo_DeleteContract(t *testing.T) {

	contract := &Contract{
		Address:      "0xDeleteContract",
		ABI:          "abi :)",
		Functions:    map[string]Function{"mint": {Arguments: []Argument{
				{Name: "nonce", Type: "bytes16"},
				{Name: "msg.sender", Type: "address"},
				{Name: "numberOfTokens", Type: "uint256"},
				{Name: "transactionNumber", Type: "uint256"},
			}}},
		ContractOwner:        "DeleteTester",
	}
	upsertErr := cr.UpsertContract(ctx, contract)
	assert.Nil(t, upsertErr)

	retrievedContract, getErr := cr.GetContract(ctx, contract.Address)
	assert.Nil(t, getErr)
	assert.Equal(t, contract, retrievedContract)

	assert.Nil(t, cr.DeleteContract(ctx, contract.Address, contract.ContractOwner))
	_, failedGetErr := cr.GetContract(ctx, contract.Address)
	assert.Error(t, failedGetErr)
}

func TestContractRepo_DeleteContract_WrongOwner(t *testing.T) {

	contract := &Contract{
		Address:      "0xDeleteContract_WrongOwner",
		ABI:          "abi :)",
		Functions:    map[string]Function{"mint": {Arguments: []Argument{
				{Name: "nonce", Type: "bytes16"},
				{Name: "msg.sender", Type: "address"},
				{Name: "numberOfTokens", Type: "uint256"},
				{Name: "transactionNumber", Type: "uint256"},
			}}},
		ContractOwner:        "DeleteTester_WrongOwner",
	}
	upsertErr := cr.UpsertContract(ctx, contract)
	assert.Nil(t, upsertErr)

	retrievedContract, getErr := cr.GetContract(ctx, contract.Address)
	assert.Nil(t, getErr)
	assert.Equal(t, contract, retrievedContract)

	assert.Error(t, cr.DeleteContract(ctx, contract.Address, "WrongOwner"))
	_, secondGetErr := cr.GetContract(ctx, contract.Address)
	assert.Nil(t, secondGetErr)
}

func TestContractRepo_GetContractsByOwner(t *testing.T) {


	owner := "TesterGetByOwnerContract"
	contracts := []*Contract{}
	for i := 0; i<10; i++ {
		contracts = append(contracts,
			&Contract{
				Address:      fmt.Sprintf("0xGetByOwnerContract_%d", i) ,
				ABI:          "abi :)",
				Functions:    map[string]Function{"mint": {Arguments: []Argument{
						{Name: "nonce", Type: "bytes16"},
						{Name: "msg.sender", Type: "address"},
						{Name: "numberOfTokens", Type: "uint256"},
						{Name: "transactionNumber", Type: "uint256"},
					}}},
				ContractOwner:        owner,
			})
	}
	for _, contract := range contracts {
		upsertErr := cr.UpsertContract(ctx, contract)
		assert.Nil(t, upsertErr)
	}

	retrievedContracts, getErr := cr.GetContractsByOwner(ctx, owner)
	assert.Nil(t, getErr)
	for i := 0; i < 10; i ++ {
		indexStr := retrievedContracts[i].Address[len(retrievedContracts[i].Address)-1:]
		index, convErr := strconv.Atoi(indexStr)
		assert.Nil(t, convErr)
		assert.Equal(t, contracts[index], retrievedContracts[i])
	}
}

func TestContract_FromRPC_ToRPC(t *testing.T) {
	contract := &Contract{
		Address:      "0xFromRPC_ToRPC",
		ABI:          "[\n\t{\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"address payable\",\n\t\t\t\t\"name\": \"artieCharAddress\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"internalType\": \"address payable\",\n\t\t\t\t\"name\": \"withdrawAddress\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"signer\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"constructor\"\n\t},\n\t{\n\t\t\"anonymous\": false,\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"indexed\": true,\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"previousOwner\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"indexed\": true,\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"newOwner\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"OwnershipTransferred\",\n\t\t\"type\": \"event\"\n\t},\n\t{\n\t\t\"anonymous\": false,\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"indexed\": false,\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"to\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"indexed\": false,\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"amount\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"indexed\": false,\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"current\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"Season01Mint\",\n\t\t\"type\": \"event\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"MAX_TOKEN\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"PURCHASE_LIMIT\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"artie\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"contract Artie\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"current\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"bytes16\",\n\t\t\t\t\"name\": \"nonce\",\n\t\t\t\t\"type\": \"bytes16\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"numberOfTokens\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"transactionNumber\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"internalType\": \"bytes\",\n\t\t\t\t\"name\": \"signature\",\n\t\t\t\t\"type\": \"bytes\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"mint\",\n\t\t\"outputs\": [],\n\t\t\"stateMutability\": \"payable\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"owner\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"price\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"renounceOwnership\",\n\t\t\"outputs\": [],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"saleStarted\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"bool\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"bool\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"signer\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"setSignerAddress\",\n\t\t\"outputs\": [],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"address payable\",\n\t\t\t\t\"name\": \"givenWithdrawalAddress\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"setWithdrawalAddress\",\n\t\t\"outputs\": [],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"signingAddress\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"startSale\",\n\t\t\"outputs\": [],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"stopSale\",\n\t\t\"outputs\": [],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"newOwner\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"transferOwnership\",\n\t\t\"outputs\": [],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"bytes16\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"bytes16\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"usedNonces\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"bool\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"bool\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"withdrawEth\",\n\t\t\"outputs\": [],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"withdrawalAddress\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"address payable\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t}\n]",
		Functions:    map[string]Function{"mint": {Arguments: []Argument{
				{Name: "nonce", Type: "bytes16"},
				{Name: "msg.sender", Type: "address"},
				{Name: "numberOfTokens", Type: "uint256"},
				{Name: "transactionNumber", Type: "uint256"},
			}}},
		ContractOwner:        "Tester",
	}

	rpc := contract.ToRPC()
	assert.Equal(t, &pb.Contract{
		Address:           "0xFromRPC_ToRPC",
		Abi:               "[\n\t{\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"address payable\",\n\t\t\t\t\"name\": \"artieCharAddress\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"internalType\": \"address payable\",\n\t\t\t\t\"name\": \"withdrawAddress\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"signer\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"constructor\"\n\t},\n\t{\n\t\t\"anonymous\": false,\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"indexed\": true,\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"previousOwner\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"indexed\": true,\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"newOwner\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"OwnershipTransferred\",\n\t\t\"type\": \"event\"\n\t},\n\t{\n\t\t\"anonymous\": false,\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"indexed\": false,\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"to\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"indexed\": false,\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"amount\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"indexed\": false,\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"current\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"Season01Mint\",\n\t\t\"type\": \"event\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"MAX_TOKEN\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"PURCHASE_LIMIT\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"artie\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"contract Artie\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"current\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"bytes16\",\n\t\t\t\t\"name\": \"nonce\",\n\t\t\t\t\"type\": \"bytes16\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"numberOfTokens\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"transactionNumber\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"internalType\": \"bytes\",\n\t\t\t\t\"name\": \"signature\",\n\t\t\t\t\"type\": \"bytes\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"mint\",\n\t\t\"outputs\": [],\n\t\t\"stateMutability\": \"payable\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"owner\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"price\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"renounceOwnership\",\n\t\t\"outputs\": [],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"saleStarted\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"bool\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"bool\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"signer\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"setSignerAddress\",\n\t\t\"outputs\": [],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"address payable\",\n\t\t\t\t\"name\": \"givenWithdrawalAddress\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"setWithdrawalAddress\",\n\t\t\"outputs\": [],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"signingAddress\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"startSale\",\n\t\t\"outputs\": [],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"stopSale\",\n\t\t\"outputs\": [],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"newOwner\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"transferOwnership\",\n\t\t\"outputs\": [],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"bytes16\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"bytes16\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"usedNonces\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"bool\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"bool\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"withdrawEth\",\n\t\t\"outputs\": [],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"withdrawalAddress\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"address payable\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t}\n]",
		Functions: map[string]*pb.Function{"mint": { Arguments: []*pb.Argument{
			{Name: "nonce", Type: "bytes16"},
			{Name: "msg.sender", Type: "address"},
			{Name: "numberOfTokens", Type: "uint256"},
			{Name: "transactionNumber", Type: "uint256"},
		}}},
		Owner:             "Tester",
	}, rpc)

	fromRPC := &Contract{}
	fromRPC.FromRPC(rpc)
	assert.Equal(t, contract, fromRPC)

	fmt.Println(fromRPC.Functions)
}