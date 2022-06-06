package contracts

import (
	"context"
	pb "contract-service/proto"
	"contract-service/storage"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"log"
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
	Indexed bool `json:"indexed"`
}

type ABIArgs struct {
	Inputs []ABIArg `json:"inputs"`
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
	//Not sure if this is the correct way to do this. An alternative would be to just convert the args array into a []string and then just convert it
	//to a [][]byte, then pass it into the signing request. We might have to go this route and find a different function for packing.

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
	//method := funcDef.Methods[functionName]
	//argTypes := method.Inputs

	args := []interface{}{}
	for _, arg := range arguments {
		args = append(args, arg)
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