package main

import (
	"context"
	"github.com/grpc-mongo-go/gen/proto"
	"github.com/grpc-mongo-go/internal/blog/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := proto.NewBlogServiceClient(conn)

	id, err := client.CreateBlog(ctx, c)
	if err != nil {
		log.Fatal(err)
	}

	blogData, err := client.ReadBlog(ctx, c, id)
	if err != nil {
		log.Fatal(err)
	}

	blogData.Content = "Modified " + time.Now().String()
	if err := client.UpdateBlog(ctx, c, blogData); err != nil {
		log.Fatal(err)
	}

	if err := client.ListBlogs(ctx, c); err != nil {
		log.Fatal(err)
	}

	if err := client.DeleteBlog(ctx, c, id); err != nil {
		log.Fatal(err)
	}
}
