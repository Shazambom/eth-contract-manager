package contracts

import (
	"context"
	pb "contract-service/proto"
	"contract-service/storage"
	"encoding/hex"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"log"
	"math/big"
	"strconv"
	"strings"
)

type ContractManagerService struct {
	writer storage.RedisWriter
	repo storage.ContractRepository
	signer pb.SigningServiceClient
}

type ABIArg struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

func NewContractTransactionHandler(writer storage.RedisWriter, repo storage.ContractRepository, signer pb.SigningServiceClient) ContractTransactionHandler {
	return &ContractManagerService{
		writer: writer,
		repo:   repo,
		signer: signer,
	}
}

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

func (cms *ContractManagerService) BuildTransaction(ctx context.Context, msgSender, functionName string, numRequested int, arguments [][]byte, contract *storage.Contract) (*storage.Token, error) {
	log.Println("Unpacking ABI")
	funcDef, abiErr := abi.JSON(strings.NewReader(contract.ABI))
	if abiErr != nil {
		return nil, abiErr
	}
	log.Println("Packing arguments")
	args, byteArgs, argParseErr := cms.UnpackArgs(msgSender, arguments, funcDef.Methods[functionName], contract.Functions.Functions[functionName])
	if argParseErr != nil {
		return nil, argParseErr
	}

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

	log.Println("Token created")
	return storage.NewToken(contract.Address, msgSender, signature.Hash, contract.ABI, packed, numRequested), nil
}


func (cms *ContractManagerService) UnpackArgs(msgSender string, arguments [][]byte, method abi.Method, hashibleFunc storage.Function) ([]interface{}, [][]byte, error) {

	//All of this splitting logic is to nicely organize the arguments, names and types
	split := strings.Split(method.String(), "(")
	otherSplit := strings.Split(split[1], ")")
	argStrs := strings.Split(otherSplit[0], ",")
	abiArgs := []ABIArg{}
	for _, str := range argStrs {
		abiArg := strings.Split(strings.TrimLeft(str, " "), " ")
		abiArgs = append(abiArgs, ABIArg{Type: abiArg[0], Name: abiArg[1]})
	}

	if len(abiArgs) - 1 != len(arguments) {
		return nil, nil, errors.New("argument length mismatch")
	}

	argBytes := [][]byte{}
	hashArgMap := map[string]int{}
	for i, arg := range hashibleFunc.Arguments {

		argBytes = append(argBytes, []byte{})
		hashArgMap[arg.Name] = i
		if arg.Name == "msg.sender" {
			argBytes[i] = common.HexToAddress(msgSender).Bytes()
		}
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
			finalArg = arg
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

func (cms *ContractManagerService) StoreToken(ctx context.Context, token *storage.Token, contract *storage.Contract) error {
	err := cms.writer.MarkAddressAsUsed(ctx, token)
	if err != nil {
		return err
	}
	if token.NumRequested < 1 {
		return nil
	}
	return cms.writer.IncrementCounter(ctx, token.NumRequested, contract.MaxMintable, contract.Address)
}

func (cms *ContractManagerService) CheckIfValidRequest(ctx context.Context, msgSender string, numRequested int, contract *storage.Contract) error{
	if numRequested > contract.MaxIncrement {
		return errors.New("max increment exceeded with request")
	}
	invalid := cms.writer.VerifyValidAddress(ctx, msgSender, contract.Address)
	if invalid != nil {
		return invalid
	}
	if numRequested < 1 {
		return nil
	}
	return cms.writer.GetReservedCount(ctx, numRequested, contract.MaxMintable, contract.Address)
}