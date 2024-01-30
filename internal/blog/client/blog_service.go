package client

import (
	"context"
	"github.com/grpc-mongo-go/gen/proto"
	"log"
	"time"
)

func CreateBlog(c proto.BlogServiceClient) (string, error) {
	log.Println("Create blog invoked")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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

func ReadBlog(c proto.BlogServiceClient, id string) (*proto.Blog, error) {
	log.Println("Read client blog invoked")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &proto.BlogId{Id: id}

	res, err := c.ReadBlog(ctx, req)
	if err != nil {
		return nil, err
	}

	log.Printf("Received blog - %v\n", res)

	return res, nil
}

func UpdateBlog(c proto.BlogServiceClient, b *proto.Blog) error {
	log.Println("Client update blog invoked")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := c.UpdatedBlog(ctx, b); err != nil {
		return err
	}

	log.Print("Client successfully updated blog")
	return nil
}
