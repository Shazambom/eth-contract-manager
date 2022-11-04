//+build wireinject

package listener

import (
	"contract-service/storage"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/wire"
)

func InitializeListenerService(cfg *aws.Config, bucket string) (EventHandlerService, error) {
	wire.Build(NewEventHandlerService, storage.NewS3)
	return &Service{}, nil
}