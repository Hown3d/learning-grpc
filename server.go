package main

import (
	"github.com/hown3d/learning-grpc/proto/helloworld"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	serverPort := os.Getenv("SERVER_PORT")
	lis, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatalf("Failed to listen on port 9000 %v", err)
	}

	s := helloworld.Server{}

	grpcServer := grpc.NewServer()
	helloworld.RegisterGreeterServer(grpcServer, &s)

	log.Println("Starting grpc server...")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start grpc Server %v", err)
	}
}
