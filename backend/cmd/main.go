package main

import (
	"github.com/razvanIncrys/simplegrpc/backend/pkg/server"
	"github.com/razvanIncrys/simplegrpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

const (
	PORT = "9000"
)

// start server
func main() {

	os.Setenv("SIMPLE_MESSAGE", "This is a simple message")

	lis, err := net.Listen("tcp", "0.0.0.0:"+PORT)
	if err != nil {
		panic("failed to listen: " + err.Error())
	}

	gServer := grpc.NewServer()

	pb.RegisterVariableServiceServer(gServer, &server.Server{})
	log.Printf("Start server on port %s", PORT)

	reflection.Register(gServer)

	err = gServer.Serve(lis)
	if err != nil {
		panic("failed to start grpc server: " + err.Error())
	}
}
