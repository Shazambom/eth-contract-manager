package contracts

import (
	pb "bitbucket.org/artie_inc/contract-service/proto"
	"bitbucket.org/artie_inc/contract-service/storage"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"log"
	"math/big"
	"strconv"
	"strings"
)

type ContractManagerService struct {
	repo storage.ContractRepository
	signer pb.SigningServiceClient
	txnRepo storage.TransactionRepository
}

type ABIArg struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

func NewContractTransactionHandler(repo storage.ContractRepository, signer pb.SigningServiceClient, txnRepo storage.TransactionRepository) ContractTransactionHandler {
	return &ContractManagerService{
		repo:   repo,
		signer: signer,
		txnRepo: txnRepo,
	}
}

//TODO Add verifier service client to the ContractManagerHandler so it can verify ownership of contracts

func NewContractManagerHandler(repo storage.ContractRepository) ContractManagerHandler {
	return &ContractManagerService{repo: repo}
}

func (cms *ContractManagerService) GetContract(ctx context.Context, address string) (*storage.Contract, error) {
	return cms.repo.GetContract(ctx, address)
}

func (cms *ContractManagerService) StoreContract(ctx context.Context, contract *storage.Contract) error {
	return cms.repo.UpsertContract(ctx, contract)
}

func (cms *ContractManagerService) DeleteContract(ctx context.Context, address, owner string) error {
	return cms.repo.DeleteContract(ctx, address, owner)
}

func (cms *ContractManagerService) ListContracts(ctx context.Context, owner string) ([]*storage.Contract, error) {
	return cms.repo.GetContractsByOwner(ctx, owner)
}

func (cms *ContractManagerService) BuildTransaction(ctx context.Context, senderInHash bool, msgSender, functionName string, arguments [][]byte, value string, contract *storage.Contract) (*storage.Transaction, error) {
	log.Println("Unpacking ABI")
	funcDef, abiErr := abi.JSON(strings.NewReader(contract.ABI))
	if abiErr != nil {
		return nil, abiErr
	}
	log.Println("Packing arguments")
	function, funcOk := contract.Functions[functionName]
	if !funcOk {
		log.Println("Function not a hashible function")
		return nil, errors.New("function selected is not hashible")
	}
	args, byteArgs, argParseErr := cms.UnpackArgs(arguments, funcDef.Methods[functionName], function)
	if argParseErr != nil {
		return nil, argParseErr
	}

	//Pre-pending value into arguments to validate the txn actually pays the contract what it's owed at runtime
	valueInt, valueOk := math.ParseBig256(value)
	if !valueOk {
		return nil, errors.New("invalid value, value is of type int256 and represents the amount of eth in wei")
	}
	byteArgs = append([][]byte{common.LeftPadBytes(valueInt.Bytes(), 32)}, byteArgs...)
	//Pre-pending sender into arguments to validate the sender of the txn is who should be sending it
	if senderInHash {
		byteArgs = append([][]byte{common.HexToAddress(msgSender).Bytes()}, byteArgs...)
	}
	//Argument priority order: msg.sender, msg.value, args...
	//So we prepend the value first, then prepend the sender so the sender goes before the value

	log.Println("Sending Signature Request")
	signature, signingErr := cms.signer.SignTxn(ctx, &pb.SignatureRequest{ContractAddress: contract.Address, Args: byteArgs})
	if signingErr != nil {
		return nil, signingErr
	}
	log.Println("Appending Signature to arguments and packing")
	sigBytes, decodeErr := hex.DecodeString(signature.Signature[2:])
	if decodeErr != nil {
		return nil, decodeErr
	}
	args = append(args, sigBytes)

	packed, packingErr := funcDef.Pack(functionName, args...)
	if packingErr != nil {
		return nil, packingErr
	}

	log.Println("Transaction created")
	return storage.NewTransaction(contract.Address, msgSender, signature.Hash, packed, value)
}

func (cms *ContractManagerService) UnpackArgs(arguments [][]byte, method abi.Method, hashibleFunc storage.Function) ([]interface{}, [][]byte, error) {
	//All of this splitting logic is to nicely organize the arguments, names and types
	if method.String() == "" {
		return nil, nil, errors.New("method could not be found")
	}
	split := strings.Split(method.String(), "(")
	otherSplit := strings.Split(split[1], ")")
	argStrs := strings.Split(otherSplit[0], ",")
	abiArgs := []ABIArg{}
	for _, str := range argStrs {
		abiArg := strings.Split(strings.TrimLeft(str, " "), " ")
		abiArgs = append(abiArgs, ABIArg{Type: abiArg[0], Name: abiArg[1]})
	}

	//We subtract 1 from abiArgs because there is an implicit signature value that is added by the service
	//TODO Decide the structure or argument structure to allow the service to pack txns without a signature (if that is needed)
	if len(abiArgs) - 1 != len(arguments) {
		return nil, nil, errors.New(fmt.Sprintf("argument length mismatch abi: %d argument length recieved: %d", len(abiArgs) - 1, len(arguments)))
	}

	argBytes := [][]byte{}
	hashArgMap := map[string]int{}
	for i, arg := range hashibleFunc.Arguments {

		argBytes = append(argBytes, []byte{})
		hashArgMap[arg.Name] = i
	}


	args := []interface{}{}
	for i, arg := range arguments {
		var finalArg interface {}
		switch abiArgs[i].Type {
		case "uint":
			bigInt, ok := math.ParseBig256(string(arg))
			if !ok {
				return nil, nil, errors.New("Unable to parse uint256")
			}
			finalArg = bigInt
			if hashInd, in := hashArgMap[abiArgs[i].Name]; in {
				argBytes[hashInd] = common.LeftPadBytes(bigInt.Bytes(), 32)
			}
		case "uint8" :
			intVar, intConvErr := strconv.Atoi(string(arg))
			if intConvErr != nil {
				return nil, nil, intConvErr
			}
			finalArg = uint8(intVar)
			if hashInd, in := hashArgMap[abiArgs[i].Name]; in {
				argBytes[hashInd] = common.LeftPadBytes(big.NewInt(int64(intVar)).Bytes(), 32)
			}
		case "uint16" :
			intVar, intConvErr := strconv.Atoi(string(arg))
			if intConvErr != nil {
				return nil, nil, intConvErr
			}
			finalArg =  uint16(intVar)
			if hashInd, in := hashArgMap[abiArgs[i].Name]; in {
				argBytes[hashInd] = common.LeftPadBytes(big.NewInt(int64(intVar)).Bytes(), 32)
			}
		case "uint32" :
			intVar, intConvErr := strconv.Atoi(string(arg))
			if intConvErr != nil {
				return nil, nil, intConvErr
			}
			finalArg = uint32(intVar)
			if hashInd, in := hashArgMap[abiArgs[i].Name]; in {
				argBytes[hashInd] = common.LeftPadBytes(big.NewInt(int64(intVar)).Bytes(), 32)
			}
		case "uint64" :
			intVar, intConvErr := strconv.Atoi(string(arg))
			if intConvErr != nil {
				return nil, nil, intConvErr
			}
			finalArg = uint64(intVar)
			if hashInd, in := hashArgMap[abiArgs[i].Name]; in {
				argBytes[hashInd] = common.LeftPadBytes(big.NewInt(int64(intVar)).Bytes(), 32)
			}
		case "uint256" :
			bigInt, ok := math.ParseBig256(string(arg))
			if !ok {
				return nil, nil, errors.New("Unable to parse uint256")
			}
			finalArg = bigInt
			if hashInd, in := hashArgMap[abiArgs[i].Name]; in {
				argBytes[hashInd] = common.LeftPadBytes(bigInt.Bytes(), 32)
			}
		case "int256" :
			bigInt, ok := math.ParseBig256(string(arg))
			if !ok {
				return nil, nil, errors.New("Unable to parse uint256")
			}
			finalArg = bigInt
			if hashInd, in := hashArgMap[abiArgs[i].Name]; in {
				argBytes[hashInd] = common.LeftPadBytes(bigInt.Bytes(), 32)
			}
		case "int8" :
			intVar, intConvErr := strconv.Atoi(string(arg))
			if intConvErr != nil {
				return nil, nil, intConvErr
			}
			finalArg = int8(intVar)
			if hashInd, in := hashArgMap[abiArgs[i].Name]; in {
				argBytes[hashInd] = common.LeftPadBytes(big.NewInt(int64(intVar)).Bytes(), 32)
			}
		case "int16" :
			intVar, intConvErr := strconv.Atoi(string(arg))
			if intConvErr != nil {
				return nil, nil, intConvErr
			}
			finalArg = int16(intVar)
			if hashInd, in := hashArgMap[abiArgs[i].Name]; in {
				argBytes[hashInd] = common.LeftPadBytes(big.NewInt(int64(intVar)).Bytes(), 32)
			}
		case "int32" :
			intVar, intConvErr := strconv.Atoi(string(arg))
			if intConvErr != nil {
				return nil, nil, intConvErr
			}
			finalArg = int32(intVar)
			if hashInd, in := hashArgMap[abiArgs[i].Name]; in {
				argBytes[hashInd] = common.LeftPadBytes(big.NewInt(int64(intVar)).Bytes(), 32)
			}
		case "int64" :
			intVar, intConvErr := strconv.Atoi(string(arg))
			if intConvErr != nil {
				return nil, nil, intConvErr
			}
			finalArg = int64(intVar)
			if hashInd, in := hashArgMap[abiArgs[i].Name]; in {
				argBytes[hashInd] = common.LeftPadBytes(big.NewInt(int64(intVar)).Bytes(), 32)
			}
		case "address" :
			var data [160]byte
			copy(data[:], arg)
			finalArg = data
			if hashInd, in := hashArgMap[abiArgs[i].Name]; in {
				argBytes[hashInd] = arg
			}
		case "bool" :
			//TODO this is likely wrong, pls fix
			finalArg = string(arg) == "true"
			if hashInd, in := hashArgMap[abiArgs[i].Name]; in {
				if string(arg) == "true" {
					byteVal, decodeErr := hex.DecodeString("01")
					if decodeErr != nil {
						return nil, nil, decodeErr
					}
					argBytes[hashInd] = byteVal
				} else {
					byteVal, decodeErr := hex.DecodeString("00")
					if decodeErr != nil {
						return nil, nil, decodeErr
					}
					argBytes[hashInd] = byteVal
				}
			}
		case "bytes":
			finalArg = arg
			if hashInd, in := hashArgMap[abiArgs[i].Name]; in {
				argBytes[hashInd] = arg
			}
		case "bytes8":
			var data [8]byte
			copy(data[:], common.LeftPadBytes(arg, 8))
			finalArg = data
			if hashInd, in := hashArgMap[abiArgs[i].Name]; in {
				argBytes[hashInd] = arg
			}
		case "bytes16":
			var data [16]byte
			copy(data[:], common.LeftPadBytes(arg, 16))
			finalArg = data
			if hashInd, in := hashArgMap[abiArgs[i].Name]; in {
				argBytes[hashInd] = arg
			}
		case "bytes24":
			var data [24]byte
			copy(data[:], common.LeftPadBytes(arg, 24))
			finalArg = data
			if hashInd, in := hashArgMap[abiArgs[i].Name]; in {
				argBytes[hashInd] = arg
			}
		case "bytes4":
			var data [4]byte
			copy(data[:], common.LeftPadBytes(arg, 4))
			finalArg = data
			if hashInd, in := hashArgMap[abiArgs[i].Name]; in {
				argBytes[hashInd] = arg
			}
		case "bytes32":
			var data [32]byte
			copy(data[:], common.LeftPadBytes(arg, 32))
			finalArg = data
			if hashInd, in := hashArgMap[abiArgs[i].Name]; in {
				argBytes[hashInd] = arg
			}
		default :
			finalArg = arg
			if hashInd, in := hashArgMap[abiArgs[i].Name]; in {
				argBytes[hashInd] = arg
			}
		}
		args = append(args, finalArg)
	}
	return args, argBytes, nil
}

func (cms *ContractManagerService) StoreTransaction(ctx context.Context, token *storage.Transaction) error {
	return cms.txnRepo.StoreTransaction(ctx, *token)
}

func (cms *ContractManagerService) GetTransactions(ctx context.Context, address string) ([]*storage.Transaction, error) {
	return cms.txnRepo.GetTransactions(ctx, address)
}

func (cms *ContractManagerService) GetAllTransactions(ctx context.Context, address string) ([]*storage.Transaction, error) {
	return cms.txnRepo.GetAllTransactions(ctx, address)
}

func (cms *ContractManagerService) DeleteTransaction(ctx context.Context, address, hash string) error {
	return cms.txnRepo.DeleteTransaction(ctx, address, hash)
}

func (cms *ContractManagerService) CompleteTransaction(ctx context.Context, address, hash string) error {
	return cms.txnRepo.CompleteTransaction(ctx, address, hash)
}