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
	client     *mongo.Client
	collection *mongo.Collection
}

// CloseDBConn disconnects the MongoDB client associated with the server.
// It takes a context as a parameter and returns an error if the disconnection process fails.
//
// Parameters:
//   - ctx: The context for the disconnection request.
//
// Returns:
//   - error: An error, if any, encountered during the MongoDB client disconnection process.
func (s *Server) CloseDBConn(ctx context.Context) error {
	if err := s.client.Disconnect(ctx); err != nil {
		return err
	}

	return nil
}

// NewBlogServer creates and returns a new instance of the Blog gRPC server.
// It initializes a MongoDB client connection based on the provided database and collection names.
// The function takes a context, database name, and collection name as parameters and returns
// the initialized server instance.
//
// Parameters:
//   - ctx: The context for the MongoDB client connection.
//   - database: The name of the MongoDB database.
//   - collection: The name of the MongoDB collection within the specified database.
//
// Returns:
//   - *Server: A new instance of the Blog gRPC server.
func NewBlogServer(ctx context.Context, database, collection string) *Server {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	return &Server{client: client, collection: client.Database(database).Collection(collection)}
}
