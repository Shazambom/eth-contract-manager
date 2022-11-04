//+build wireinject

package signing

import (
	"contract-service/storage"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

func InitializeSigningServer(host int, opts []grpc.ServerOption, tableName string, cfg ...*aws.Config) (*SignerRPCService, error) {
	wire.Build(NewSignerServer, NewSigningService, storage.NewPrivateKeyRepository)
	return &SignerRPCService{}, nil
}
