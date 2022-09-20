package gClient

import (
	"context"
	"github.com/razvanIncrys/simplegrpc/pb"
	"google.golang.org/grpc"
)

type Client struct {
	VariableServiceClient pb.VariableServiceClient
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{pb.NewVariableServiceClient(conn)}
}

func (c *Client) GetVariable(ctx context.Context, variableName string) (*pb.GetVariableResponse, error) {
	return &pb.GetVariableResponse{VariableValue: variableName}, nil
}
