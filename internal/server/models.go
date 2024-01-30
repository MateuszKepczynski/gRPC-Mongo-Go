package server

import (
	"context"
	"fmt"
	"github.com/grpc-mongo-go/internal/db/models"
	"github.com/grpc-mongo-go/proto/blog"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type Server struct {
	blog.BlogServiceServer
	client *mongo.Client
}

func (s *Server) CreateBlog(ctx context.Context, req *blog.Blog) (*blog.BlogId, error) {
	data := models.BlogItem{
		AuthorID: req.AuthorId,
		Title:    req.Tile,
		Content:  req.Content,
	}

	log.Println("Inserting data to DB.")
	res, err := s.client.Database(blogdb).Collection(blogCollection).InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v\n", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil,
			status.Errorf(
				codes.Internal,
				"Cannot convert insertedID to OID",
			)
	}

	return &blog.BlogId{Id: oid.Hex()}, nil
}

func (s *Server) CloseDBConn(ctx context.Context) error {
	if err := s.client.Disconnect(ctx); err != nil {
		return err
	}

	return nil
}

func NewBlogServer(ctx context.Context) *Server {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:root@mongo:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	return &Server{client: client}
}
