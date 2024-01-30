package main

import (
	"github.com/grpc-mongo-go/internal/client"
	"github.com/grpc-mongo-go/proto/blog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const address = "localhost:5051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := blog.NewBlogServiceClient(conn)

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
}
