package main

import (
	"github.com/hown3d/learning-grpc/proto/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	var connection *grpc.ClientConn
	serverAddr := os.Getenv("SERVER_ADDR")
	serverPort := os.Getenv("SERVER_PORT")

	connection, err := grpc.Dial(serverAddr+":"+serverPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldnt connect to grpc server on port :9000 %v", err)
	}

	hello := helloworld.NewGreeterClient(connection)

	response, err := hello.SayHello(context.Background(), &helloworld.HelloRequest{Name: "Lukas"})

	_, err = hello.SayGoodbye(context.Background(), &helloworld.HelloRequest{
		Name: "Goodbye",
	})
	if err != nil {
		return
	}

	log.Printf("Got response %v", response)

}
