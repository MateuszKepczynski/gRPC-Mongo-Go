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

// DocumentToBlog converts a BlogItem document to a proto.Blog.
// It takes a BlogItem pointer as receiver and returns a new proto.Blog
// instance populated with the corresponding data from the BlogItem.
//
// Returns:
//   - *proto.Blog: A proto.Blog instance with data from the BlogItem.
func (b *BlogItem) DocumentToBlog() *proto.Blog {
	return &proto.Blog{
		Id:       b.ID.Hex(),
		AuthorId: b.AuthorID,
		Tile:     b.Title,
		Content:  b.Content,
	}
}
