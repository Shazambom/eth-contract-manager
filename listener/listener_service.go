package listener

import (
	"contract-service/storage"
	"encoding/json"
	"fmt"
)

type Service struct {
	s3 *storage.S3
}

func NewEventHandlerService(s3 *storage.S3) EventHandlerService {
	return &Service{s3: s3}
}

func (ls *Service) InitService() error {
	return ls.s3.InitBucket()
}

func (ls *Service) Handle(key, val string, err error) error {
	if err != nil {
		fmt.Printf("Error with redis stream: %s\n", err)
		return err
	}
	fmt.Printf("key: %s\nval:%s\n", key, val)
	token := storage.Token{}
	parseErr := json.Unmarshal([]byte(val), &token)
	if parseErr != nil {
		fmt.Printf("Error with redis format: %s\n", parseErr.Error())
		return nil
	}
	storeErr := ls.s3.StoreToken(&token)
	if storeErr != nil {
		fmt.Printf("Error storing in s3: %s\n", storeErr.Error())
	}
	return nil
}
