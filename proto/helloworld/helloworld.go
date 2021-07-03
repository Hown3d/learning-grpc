package helloworld

import (
	"golang.org/x/net/context"

	log "github.com/sirupsen/logrus"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *HelloRequest) (*HelloReply, error) {
	log.Printf("recieved message from %v", message.Name)

	return &HelloReply{Message: "Hello " + message.Name}, nil

}

func (s *Server) SayGoodbye(ctx context.Context, request *HelloRequest) (*HelloReply, error) {
	log.Println("implement me")
	return nil, nil
}

func (s *Server) mustEmbedUnimplementedGreeterServer() {}
