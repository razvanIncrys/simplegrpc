package gClient

import (
	"context"
	"github.com/razvanIncrys/simplegrpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Client struct {
	VariableServiceClient pb.VariableServiceClient
	connection            *grpc.ClientConn
}

func NewClient(port string) *Client {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	variableServiceClient := pb.NewVariableServiceClient(conn)
	return &Client{VariableServiceClient: variableServiceClient, connection: conn}
}

func (c *Client) GetMyVariable(ctx context.Context, in *pb.GetMyVariableRequest) (*pb.GetMyVariableResponse, error) {
	return c.VariableServiceClient.GetMyVariable(ctx, in)
}
func (c *Client) ListAllMyVariables(ctx context.Context, in *pb.ListMyAllVariableRequest, opts ...grpc.CallOption) (*pb.ListAllMyVariableResponse, error) {
	return c.VariableServiceClient.ListAllMyVariables(ctx, in, grpc.WaitForReady(true))
}
func (c *Client) DeleteAllMyVariables(ctx context.Context, in *pb.DeleteAllMyVariableRequest, opts ...grpc.CallOption) (*pb.DeleteAllMyVariableResponse, error) {
	return c.VariableServiceClient.DeleteAllMyVariables(ctx, in, opts...)
}

func (c *Client) Close() error {
	return c.connection.Close()
}
