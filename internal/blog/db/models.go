package db

import (
	"github.com/grpc-mongo-go/gen/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"authorID"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

func (b *BlogItem) DocumentToBlog() *proto.Blog {
	return &proto.Blog{
		Id:       b.ID.Hex(),
		AuthorId: b.AuthorID,
		Tile:     b.Title,
		Content:  b.Content,
	}
}
