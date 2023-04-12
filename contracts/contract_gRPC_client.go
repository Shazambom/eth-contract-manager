package contracts

import (
	pb "bitbucket.org/artie_inc/contract-service/proto"
	"google.golang.org/grpc"
)

type ContractClient struct {
	Connection *grpc.ClientConn
	Client     pb.ContractManagementClient
}

func NewContractClient(host string, opts []grpc.DialOption) (*ContractClient, error) {
	conn, err := grpc.Dial(host, opts...)
	if err != nil {
		return nil, err
	}
	return &ContractClient{Connection: conn, Client: pb.NewContractManagementClient(conn)}, nil
}

func (c *ContractClient) DisconnectGracefully() error {
	closeConnErr := c.Connection.Close()
	if closeConnErr != nil {
		return closeConnErr
	}
	return nil
}
