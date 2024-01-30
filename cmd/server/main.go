package main

import (
	"context"
	"github.com/grpc-mongo-go/internal/server"
	"github.com/grpc-mongo-go/proto/blog"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func main() {
	log.Printf("Starting server with address - %s", address)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	blogServer := server.NewBlogServer(ctx)
	defer blogServer.CloseDBConn(ctx)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	s := grpc.NewServer()

	blog.RegisterBlogServiceServer(s, blogServer)

	if err := s.Serve(lis); err != nil {
		log.Fatal(lis)
	}
}
