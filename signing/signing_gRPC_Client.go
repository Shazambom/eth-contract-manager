package signing

import (
	pb "contract-service/proto"
	"google.golang.org/grpc"
)

type Client struct {
	Connection *grpc.ClientConn
	SigningClient pb.SigningServiceClient
}

func NewClient(host string, opts []grpc.DialOption) (*Client, error) {
	conn, err := grpc.Dial(host, opts...)
	if err != nil {
		return nil, err
	}
	return &Client{Connection: conn, SigningClient: pb.NewSigningServiceClient(conn)}, nil
}


func (c *Client) DisconnectGracefully() error {
	closeConnErr := c.Connection.Close()
	if closeConnErr != nil {
		return closeConnErr
	}
	return nil
}

