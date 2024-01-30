package client

import (
	"context"
	"github.com/grpc-mongo-go/proto/blog"
	"log"
	"time"
)

func CreateBlog(c blog.BlogServiceClient) (string, error) {
	log.Println("Create blog invoked")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &blog.Blog{
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

func ReadBlog(c blog.BlogServiceClient, id string) (*blog.Blog, error) {
	log.Println("Read client blog invoked")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &blog.BlogId{Id: id}

	res, err := c.ReadBlog(ctx, req)
	if err != nil {
		return nil, err
	}

	log.Printf("Received blog - %v\n", res)

	return res, nil
}
