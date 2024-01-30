package server

import (
	"context"
	"github.com/grpc-mongo-go/gen/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Server struct {
	proto.BlogServiceServer
	client *mongo.Client
}

func (s *Server) CloseDBConn(ctx context.Context) error {
	if err := s.client.Disconnect(ctx); err != nil {
		return err
	}

	return nil
}

func NewBlogServer(ctx context.Context) *Server {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	return &Server{client: client}
}
