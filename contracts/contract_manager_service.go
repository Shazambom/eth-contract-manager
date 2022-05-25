package contracts

import "contract-service/storage"

type ContractManagerService struct {
	writer storage.RedisWriter
}



func (cms *ContractManagerService) GetContract() (*storage.Contract, error) {
	return nil, nil
}

func (cms *ContractManagerService) BuildTransaction(contract *storage.Contract) (*storage.Token, error) {
	return nil, nil
}

func (cms *ContractManagerService) StoreToken(token *storage.Token) error {
	return nil
}