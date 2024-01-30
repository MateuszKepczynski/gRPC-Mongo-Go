package server

import (
	"context"
	"fmt"
	"github.com/grpc-mongo-go/internal/db/models"
	"github.com/grpc-mongo-go/proto/blog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Server) CreateBlog(ctx context.Context, req *blog.Blog) (*blog.BlogId, error) {
	log.Println("Create blog server invoked")
	data := models.BlogItem{
		AuthorID: req.AuthorId,
		Title:    req.Tile,
		Content:  req.Content,
	}

	log.Println("Inserting data to DB")
	res, err := s.client.Database(blogdb).Collection(blogCollection).InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v\n", err),
		)
	}

	log.Println("Data inserted")
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil,
			status.Errorf(
				codes.Internal,
				fmt.Sprintf("Cannot convert insertedID to OID %v\n", err),
			)
	}

	return &blog.BlogId{Id: oid.Hex()}, nil
}

func (s *Server) ReadBlog(ctx context.Context, req *blog.BlogId) (*blog.Blog, error) {
	log.Println("Read blog invoked")

	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert insertedID to OID - %v\n", err),
		)
	}

	filter := bson.M{"_id": oid}
	res := s.client.Database(blogdb).Collection(blogCollection).FindOne(ctx, filter)

	data := &models.BlogItem{}
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot decode MongoDB response - %v\n", err),
		)
	}

	log.Println("Data successfully obtained")
	return data.DocumentToBlog(), nil
}

func (s *Server) UpdatedBlog(ctx context.Context, req *blog.Blog) (*emptypb.Empty, error) {
	log.Println("Server Update Blog Invoked")
	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert insertedID to OID - %v\n", err),
		)
	}

	data := &models.BlogItem{

		AuthorID: req.AuthorId,
		Title:    req.Tile,
		Content:  req.Content,
	}

	res, err := s.client.
		Database(blogdb).
		Collection(blogCollection).
		UpdateOne(
			ctx,
			bson.M{"_id": oid},
			bson.M{"$set": data},
		)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Could not update - %v\n", err),
		)
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find blog with ID",
		)
	}

	log.Println("Server successfully updated blog")

	return &emptypb.Empty{}, nil
}
