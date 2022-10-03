package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/razvanIncrys/simplegrpc/backend/pb"
	"github.com/razvanIncrys/simplegrpc/backend/pkg/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
)

const (
	PORT      = "9000"
	HTTP_PORT = "80"
)

// start server
func main() {

	os.Setenv("SIMPLE_MESSAGE", "This is a simple message")

	myServer := server.Server{}

	go func() {
		mux := runtime.NewServeMux()
		pb.RegisterVariableServiceHandlerServer(context.Background(), mux, &myServer)
		log.Printf("Start hgateway on port %s", HTTP_PORT)
		log.Fatalln("gateway error:", http.ListenAndServe(":"+HTTP_PORT, mux))
	}()

	lis, err := net.Listen("tcp", "0.0.0.0:"+PORT)
	if err != nil {
		panic("failed to listen: " + err.Error())
	}

	gServer := grpc.NewServer()

	pb.RegisterVariableServiceServer(gServer, &myServer)
	log.Printf("Start server on port %s", PORT)

	reflection.Register(gServer)

	err = gServer.Serve(lis)
	if err != nil {
		panic("failed to start grpc server: " + err.Error())
	}
}
