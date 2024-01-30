package client

import (
	"context"
	"github.com/grpc-mongo-go/gen/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
)

// CreateBlog creates a new blog using the provided BlogServiceClient.
// It sends a request to the server to create a blog with the specified
// AuthorId, Title, and Content. The function returns the ID of the
// created blog and an error, if any.
//
// Parameters:
//   - c: The BlogServiceClient used to communicate with the server.
//
// Returns:
//   - string: The ID of the created blog.
//   - error: An error, if any, encountered during the blog creation process.
func CreateBlog(ctx context.Context, c proto.BlogServiceClient) (string, error) {
	log.Println("Create blog invoked")

	req := &proto.Blog{
		AuthorId: "Matthew",
		Tile:     "Hello",
		Content:  "World",
	}

	res, err := c.CreateBlog(ctx, req)
	if err != nil {
		return "", err
	}

	log.Printf("Created blog with ID %s\n", res.Id)

	return res.Id, nil
}

// ReadBlog retrieves a blog with the specified ID using the provided BlogServiceClient.
// It sends a request to the server to read a blog based on the given ID. The function
// returns the retrieved blog and an error, if any.
//
// Parameters:
//   - c: The BlogServiceClient used to communicate with the server.
//   - id: The ID of the blog to be retrieved.
//
// Returns:
//   - *proto.Blog: The retrieved blog.
//   - error: An error, if any, encountered during the blog retrieval process.
func ReadBlog(ctx context.Context, c proto.BlogServiceClient, id string) (*proto.Blog, error) {
	log.Println("Read client blog invoked")

	req := &proto.BlogId{Id: id}

	res, err := c.ReadBlog(ctx, req)
	if err != nil {
		return nil, err
	}

	log.Printf("Received blog - %v\n", res)

	return res, nil
}

// UpdateBlog updates an existing blog using the provided BlogServiceClient.
// It sends a request to the server to update the specified blog. The function
// returns an error if the update process encounters any issues.
//
// Parameters:
//   - c: The BlogServiceClient used to communicate with the server.
//   - b: The blog with updated information to be sent for updating.
//
// Returns:
//   - error: An error, if any, encountered during the blog update process.
func UpdateBlog(ctx context.Context, c proto.BlogServiceClient, b *proto.Blog) error {
	log.Println("Client update blog invoked")

	if _, err := c.UpdateBlog(ctx, b); err != nil {
		return err
	}

	log.Print("Client successfully updated blog")
	return nil
}

// ListBlogs retrieves a list of blogs using the provided BlogServiceClient.
// It sends a request to the server to list all available blogs. The function
// logs each received blog and returns an error if the list operation encounters any issues.
//
// Parameters:
//   - c: The BlogServiceClient used to communicate with the server.
//
// Returns:
//   - error: An error, if any, encountered during the blog list operation.
func ListBlogs(ctx context.Context, c proto.BlogServiceClient) error {
	log.Println("Client list blogs invoked")

	stream, err := c.ListBlogs(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}

	log.Println("Starting to receive response from server")
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			log.Println("Stopping receiving data from server")
			return nil
		}
		if err != nil {
			return err
		}

		log.Printf("Received blog - %v\n", res)
	}
}

// DeleteBlog deletes a blog with the specified ID using the provided BlogServiceClient.
// It sends a request to the server to delete the blog identified by the given ID.
// The function returns an error if the delete operation encounters any issues.
//
// Parameters:
//   - c: The BlogServiceClient used to communicate with the server.
//   - id: The ID of the blog to be deleted.
//
// Returns:
//   - error: An error, if any, encountered during the blog delete operation.
func DeleteBlog(ctx context.Context, c proto.BlogServiceClient, id string) error {
	log.Println("Client delete blog invoked")

	_, err := c.DeleteBlog(ctx, &proto.BlogId{Id: id})
	if err != nil {
		return err
	}

	log.Println("Record deleted")

	return nil
}
