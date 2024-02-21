package main

import (
	"fmt"
	"log"
	"net"

	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	desc "github.com/Tim-Sa/chat-server/pkg/chat_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedChatV1Server
}

func main() {
	fmt.Println(color.GreenString("Chat service start"))

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", listen.Addr())

	if err = s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
