package main

import (
	"context"
	"github.com/grpc-mongo-go/gen/proto"
	"github.com/grpc-mongo-go/internal/blog/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	log.Printf("Starting server with address - %s\n", address)

	ctx := context.Background()

	blogServer := server.NewBlogServer(ctx, blogDB, blogCollection)
	defer blogServer.CloseDBConn(ctx)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	s := grpc.NewServer()

	proto.RegisterBlogServiceServer(s, blogServer)

	if err := s.Serve(lis); err != nil {
		log.Fatal(lis)
	}
}
