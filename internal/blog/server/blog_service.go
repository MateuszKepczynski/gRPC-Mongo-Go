package server

import (
	"context"
	"fmt"
	"github.com/grpc-mongo-go/gen/proto"
	"github.com/grpc-mongo-go/internal/blog/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

// CreateBlog is a gRPC server method that creates a new blog based on the provided proto.Blog.
// It inserts the blog data into the server's MongoDB collection and returns the ID of the
// newly created blog along with any error encountered during the process.
//
// Parameters:
//   - ctx: The context for the request.
//   - req: The proto.Blog containing the blog data.
//
// Returns:
//   - *proto.BlogId: The ID of the newly created blog.
//   - error: An error, if any, encountered during the blog creation process.
func (s *Server) CreateBlog(ctx context.Context, req *proto.Blog) (*proto.BlogId, error) {
	log.Println("Create blog server invoked")
	data := db.BlogItem{
		AuthorID: req.AuthorId,
		Title:    req.Tile,
		Content:  req.Content,
	}

	log.Println("Inserting data to DB")
	res, err := s.collection.InsertOne(ctx, data)
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

	return &proto.BlogId{Id: oid.Hex()}, nil
}

// ReadBlog is a gRPC server method that retrieves a blog with the specified ID.
// It queries the server's MongoDB collection and returns the corresponding blog
// along with any error encountered during the process.
//
// Parameters:
//   - ctx: The context for the request.
//   - req: The proto.BlogId containing the ID of the blog to be retrieved.
//
// Returns:
//   - *proto.Blog: The retrieved blog.
//   - error: An error, if any, encountered during the blog retrieval process.
func (s *Server) ReadBlog(ctx context.Context, req *proto.BlogId) (*proto.Blog, error) {
	log.Println("Read blog invoked")

	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot convert insertedID to OID - %v\n", err),
		)
	}

	filter := bson.M{"_id": oid}
	res := s.collection.FindOne(ctx, filter)

	data := &db.BlogItem{}
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot decode MongoDB response - %v\n", err),
		)
	}

	log.Println("Data successfully obtained")
	return data.DocumentToBlog(), nil
}

// UpdateBlog is a gRPC server method that updates an existing blog with the provided data.
// It modifies the corresponding blog in the server's MongoDB collection and returns an empty
// response along with any error encountered during the process.
//
// Parameters:
//   - ctx: The context for the request.
//   - req: The proto.Blog containing the updated blog data.
//
// Returns:
//   - *emptypb.Empty: An empty response.
//   - error: An error, if any, encountered during the blog update process.
func (s *Server) UpdateBlog(ctx context.Context, req *proto.Blog) (*emptypb.Empty, error) {
	log.Println("Server Update Blog Invoked")
	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot convert insertedID to OID - %v\n", err),
		)
	}

	data := &db.BlogItem{

		AuthorID: req.AuthorId,
		Title:    req.Tile,
		Content:  req.Content,
	}

	res, err := s.collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": data})
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

// ListBlogs is a gRPC server method that retrieves a list of blogs from the server's MongoDB collection.
// It streams the retrieved blogs to the client and returns an error if the list operation encounters any issues.
//
// Parameters:
//   - _ : The empty request parameter.
//   - stream: The gRPC stream to send the list of blogs.
//
// Returns:
//   - error: An error, if any, encountered during the blog list operation.
func (s *Server) ListBlogs(_ *emptypb.Empty, stream proto.BlogService_ListBlogsServer) error {
	ctx := context.Background()
	res, err := s.collection.Find(ctx, primitive.D{{}})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Could not find blogs - %v\n", err),
		)
	}
	defer res.Close(ctx)

	for res.Next(ctx) {
		data := &db.BlogItem{}
		if err := res.Decode(data); err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Cannot decode response to blog structure - %v\n", err),
			)
		}

		if err := stream.Send(data.DocumentToBlog()); err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Cannot send data - %v\n", err),
			)
		}
	}

	if err := res.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknow error - %v\n", err),
		)
	}

	return nil
}

// DeleteBlog is a gRPC server method that deletes a blog with the specified ID from the server's MongoDB collection.
// It returns an empty response along with any error encountered during the deletion process.
//
// Parameters:
//   - ctx: The context for the request.
//   - req: The proto.BlogId containing the ID of the blog to be deleted.
//
// Returns:
//   - *emptypb.Empty: An empty response.
//   - error: An error, if any, encountered during the blog deletion process.
func (s *Server) DeleteBlog(ctx context.Context, req *proto.BlogId) (*emptypb.Empty, error) {
	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot convert insertedID to OID - %v\n", err),
		)
	}

	res, err := s.collection.DeleteOne(ctx, primitive.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete object under ID '%s' with error: %v", req.Id, err),
		)
	}

	if res.DeletedCount < 1 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find any objects with given ID '%s'", req.Id),
		)
	}

	return &emptypb.Empty{}, nil
}
