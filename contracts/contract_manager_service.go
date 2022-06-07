package contracts

import (
	"context"
	pb "contract-service/proto"
	"contract-service/storage"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/math"
	"log"
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

func (cms *ContractManagerService) BuildTransaction(ctx context.Context, msgSender, functionName string, numRequested int, arguments []string, contract *storage.Contract) (*storage.Token, error) {
	log.Println("Unpacking ABI")
	funcDef, abiErr := abi.JSON(strings.NewReader(contract.ABI))
	if abiErr != nil {
		return nil, abiErr
	}
	log.Println("Packing arguments")
	args, argParseErr := cms.UnpackArgs(arguments, functionName, funcDef)
	if argParseErr != nil {
		return nil, argParseErr
	}

	packed, packingErr := funcDef.Pack(functionName, args...)
	if packingErr != nil {
		return nil, packingErr
	}
	log.Println("Sending Signature Request")
	signature, signingErr := cms.signer.SignTxn(ctx, &pb.SignatureRequest{ContractAddress: contract.Address, Args: [][]byte{packed}})
	if signingErr != nil {
		return nil, signingErr
	}
	log.Println("Appending Signature to arguments and packing")
	args = append(args, signature.Signature)

	repacked, repackingErr := funcDef.Pack(functionName, args...)
	if repackingErr != nil {
		return nil, repackingErr
	}

	log.Println("Token created")
	return storage.NewToken(contract.Address, msgSender, signature.Hash, contract.ABI, repacked, numRequested), nil
}

func (cms *ContractManagerService) UnpackArgs(arguments []string, functionName string, funcDef abi.ABI) ([]interface{}, error) {
	method := funcDef.Methods[functionName]

	//All of this splitting logic is to nicely organize the arguments, names and types
	split := strings.Split(method.String(), "(")
	otherSplit := strings.Split(split[1], ")")
	argStrs := strings.Split(otherSplit[0], ",")
	abiArgs := []ABIArg{}
	for _, str := range argStrs {
		abiArg := strings.Split(strings.TrimLeft(str, " "), " ")
		abiArgs = append(abiArgs, ABIArg{Type: abiArg[0], Name: abiArg[1]})
	}

	if len(abiArgs) != len(arguments) {
		return nil, errors.New("argument length mismatch")
	}


	args := []interface{}{}
	for i, arg := range arguments {
		var finalArg interface {}
		switch abiArgs[i].Type {
		case "uint":
			bigInt, ok := math.ParseBig256(arg)
			if !ok {
				return nil, errors.New("Unable to parse uint256")
			}
			finalArg = bigInt
		case "uint8" :
			intVar, intConvErr := strconv.Atoi(arg)
			if intConvErr != nil {
				return nil, intConvErr
			}
			finalArg = uint8(intVar)

		case "uint16" :
			intVar, intConvErr := strconv.Atoi(arg)
			if intConvErr != nil {
				return nil, intConvErr
			}
			finalArg =  uint16(intVar)
		case "uint32" :
			intVar, intConvErr := strconv.Atoi(arg)
			if intConvErr != nil {
				return nil, intConvErr
			}
			finalArg = uint32(intVar)
		case "uint64" :
			intVar, intConvErr := strconv.Atoi(arg)
			if intConvErr != nil {
				return nil, intConvErr
			}
			finalArg = uint64(intVar)
		case "uint256" :
			bigInt, ok := math.ParseBig256(arg)
			if !ok {
				return nil, errors.New("Unable to parse uint256")
			}
			finalArg = bigInt
		case "int256" :
			bigInt, ok := math.ParseBig256(arg)
			if !ok {
				return nil, errors.New("Unable to parse uint256")
			}
			finalArg = bigInt
		case "int8" :
			intVar, intConvErr := strconv.Atoi(arg)
			if intConvErr != nil {
				return nil, intConvErr
			}
			finalArg = int8(intVar)
		case "int16" :
			intVar, intConvErr := strconv.Atoi(arg)
			if intConvErr != nil {
				return nil, intConvErr
			}
			finalArg = int16(intVar)
		case "int32" :
			intVar, intConvErr := strconv.Atoi(arg)
			if intConvErr != nil {
				return nil, intConvErr
			}
			finalArg = int32(intVar)
		case "int64" :
			intVar, intConvErr := strconv.Atoi(arg)
			if intConvErr != nil {
				return nil, intConvErr
			}
			finalArg = int64(intVar)
		case "address" :
			var address [160]byte
			copy(address[:], arg)
			finalArg = address
		case "bool" :
			finalArg = arg == "true"
		case "bytes":
			finalArg = []byte(arg)
		case "bytes8":
			var data [8]byte
			copy(data[:], arg)
			finalArg = data
		case "bytes16":
			var data [16]byte
			copy(data[:], arg)
			finalArg = data
		case "bytes24":
			var data [24]byte
			copy(data[:], arg)
			finalArg = data
		case "bytes4":
			var data [4]byte
			copy(data[:], arg)
			finalArg = data
		case "bytes32":
			var data [32]byte
			copy(data[:], arg)
			finalArg = data
		default :
			finalArg = []byte(arg)
		}
		args = append(args, finalArg)
	}
	return args, nil
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