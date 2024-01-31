package main

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-mongo-go/gen/proto"
	"github.com/grpc-mongo-go/internal/blog/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"runtime/debug"
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

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
		),
		grpc.ChainStreamInterceptor(
			recovery.StreamServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
		),
	)

	proto.RegisterBlogServiceServer(s, blogServer)

	if err := s.Serve(lis); err != nil {
		log.Fatal(lis)
	}
}

func grpcPanicRecoveryHandler(p any) (err error) {
	log.Println("msg", "recovered from panic", "panic", p, "stack", string(debug.Stack()))
	return status.Errorf(codes.Internal, "%s", p)
}
