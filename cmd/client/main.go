package main

import (
	"github.com/grpc-mongo-go/gen/proto"
	"github.com/grpc-mongo-go/internal/blog/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := proto.NewBlogServiceClient(conn)

	id, err := client.CreateBlog(c)
	if err != nil {
		log.Fatal(err)
	}

	blogData, err := client.ReadBlog(c, id)
	if err != nil {
		log.Fatal(err)
	}

	blogData.Content = "Modified " + time.Now().String()
	if err := client.UpdateBlog(c, blogData); err != nil {
		log.Fatal(err)
	}

	if err := client.ListBlogs(c); err != nil {
		log.Fatal(err)
	}

	if err := client.DeleteBlog(c, id); err != nil {
		log.Fatal(err)
	}
}
