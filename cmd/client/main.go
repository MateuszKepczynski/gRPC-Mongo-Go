package main

import (
	"github.com/grpc-mongo-go/internal/client"
	"github.com/grpc-mongo-go/proto/blog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
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
	_, err = client.ReadBlog(c, id)
	if err != nil {
		log.Fatal(err)
	}

}
