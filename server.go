package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"

	filters "github.com/Romero027/echo-grpc-adn/filters"
	echo "github.com/Romero027/echo-grpc-adn/pb"
	grpc "github.com/Romero027/grpc-go"
)

type server struct {
	echo.UnimplementedEchoServiceServer
}

func (s *server) Echo(ctx context.Context, x *echo.Msg) (*echo.Msg, error) {
	log.Printf("got: [%s]", x.GetBody())
	return x, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// s := grpc.NewServer(grpc.UnaryInterceptor(grpc.ChainUnaryServer(UnaryServerInterceptor, UnaryServerInterceptor2)))
	s := grpc.NewServer(grpc.UnaryInterceptor(grpc.ChainUnaryServer(filters.ContentBasedACL)))
	fmt.Printf("Starting server at port 9000\n")

	echo.RegisterEchoServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
