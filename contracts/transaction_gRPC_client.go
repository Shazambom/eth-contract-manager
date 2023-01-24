package contracts

import (
	pb "bitbucket.org/artie_inc/contract-service/proto"
	"google.golang.org/grpc"
)

type TransactionClient struct {
	Connection *grpc.ClientConn
	Client pb.TransactionServiceClient
}

func NewTransactionClient(host string, opts []grpc.DialOption) (*TransactionClient, error) {
	conn, err := grpc.Dial(host, opts...)
	if err != nil {
		return nil, err
	}
	return &TransactionClient{Connection: conn, Client: pb.NewTransactionServiceClient(conn)}, nil
}


func (c *TransactionClient) DisconnectGracefully() error {
	closeConnErr := c.Connection.Close()
	if closeConnErr != nil {
		return closeConnErr
	}
	return nil
}
