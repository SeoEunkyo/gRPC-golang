//server/main.go
package main

import (
	"context"
	bookpd "github.com/SeoEunkyo/gRPC-golang/protos/v1"
	"github.com/SeoEunkyo/gRPC-golang/server/data"
	"google.golang.org/grpc"
	"log"
	"net"
)

const portNumber = "9000"

type bookServer struct {
	bookpd.BookServer
}

func (s *bookServer) GetBook(ctx context.Context, req *bookpd.GetBookRequest) (*bookpd.GetBookResponse, error) {
	bookId := req.Id

	var bookMessage *bookpd.BookMessage

	bookMessage = data.BookData[bookId]
	return &bookpd.GetBookResponse{
		BookMessage: bookMessage,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	bookpd.RegisterBookServer(grpcServer, &bookServer{})

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
