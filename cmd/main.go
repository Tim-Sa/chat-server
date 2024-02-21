package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math"
	"math/big"
	"net"

	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/Tim-Sa/chat-server/pkg/chat_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedChatV1Server
}

func randIndex() int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		return 0
	}
	return nBig.Int64()
}

func (s *server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	usernames := req.GetUsernames()

	log.Printf("chat with %v created.", usernames)

	return &desc.CreateResponse{
		Id: randIndex(),
	}, nil
}

func (s *server) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	id := req.GetId()

	log.Printf("chat %d deleted.", id)

	empty := emptypb.Empty{}
	return &empty, nil
}

func (s *server) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	from := req.GetFrom()
	text := req.GetText()
	when := req.GetTimestamp()

	log.Printf("message from %s:\n\n\t %s \n\n at %v\n", from, text, when)

	empty := emptypb.Empty{}
	return &empty, nil
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
