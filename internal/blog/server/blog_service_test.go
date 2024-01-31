package server

import (
	"context"
	"fmt"
	"github.com/grpc-mongo-go/gen/proto"
	"github.com/grpc-mongo-go/internal/blog/db"
	servermock "github.com/grpc-mongo-go/mocks/grpc/server"
	"github.com/grpc-mongo-go/mocks/mongo/collection"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"reflect"
	"testing"
)

const (
	blogID  = "65b9018f4e0d763afee645f8"
	wrongID = "¤"
)

var blogOID, _ = primitive.ObjectIDFromHex(blogID)

func TestServer_CreateBlog(t *testing.T) {
	type fields struct {
		collection CollectionInterface
	}
	type args struct {
		ctx context.Context
		req *proto.Blog
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       *proto.BlogId
		wantErrMsg string
		setup      func(f *fields)
	}{
		{
			name: "ShouldErrorOnInsertData",
			args: args{
				ctx: context.Background(),
				req: &proto.Blog{
					AuthorId: "1",
					Tile:     "foo title",
					Content:  "foo content",
				},
			},
			fields: fields{
				collection: nil,
			},
			want: &proto.BlogId{
				Id: blogID,
			},
			wantErrMsg: "Internal error: foo bar\n",
			setup: func(f *fields) {
				expectedBlogItem := db.BlogItem{
					AuthorID: "1",
					Title:    "foo title",
					Content:  "foo content",
				}
				colMock := collection.NewCollectionMock(t)
				colMock.
					EXPECT().
					InsertOne(context.Background(), expectedBlogItem).
					Return(nil, fmt.Errorf("foo bar")).
					Once()

				f.collection = colMock
			},
		},
		{
			name: "ShouldErrorOnWrongID",
			args: args{
				ctx: context.Background(),
				req: &proto.Blog{
					AuthorId: "1",
					Tile:     "foo title",
					Content:  "foo content",
				},
			},
			fields: fields{
				collection: nil,
			},
			want: &proto.BlogId{
				Id: blogID,
			},
			wantErrMsg: "Cannot convert insertedID to OID <nil>\n",
			setup: func(f *fields) {
				expectedBlogItem := db.BlogItem{
					AuthorID: "1",
					Title:    "foo title",
					Content:  "foo content",
				}
				expectedRes := &mongo.InsertOneResult{
					InsertedID: byte(1) < 3,
				}
				colMock := collection.NewCollectionMock(t)
				colMock.
					EXPECT().
					InsertOne(context.Background(), expectedBlogItem).
					Return(expectedRes, nil).
					Once()

				f.collection = colMock
			},
		},
		{
			name: "ShouldCreateBlog",
			args: args{
				ctx: context.Background(),
				req: &proto.Blog{
					AuthorId: "1",
					Tile:     "foo title",
					Content:  "foo content",
				},
			},
			fields: fields{
				collection: nil,
			},
			want: &proto.BlogId{
				Id: blogID,
			},
			setup: func(f *fields) {
				expectedBlogItem := db.BlogItem{
					AuthorID: "1",
					Title:    "foo title",
					Content:  "foo content",
				}

				expectedRes := &mongo.InsertOneResult{
					InsertedID: blogOID,
				}
				colMock := collection.NewCollectionMock(t)
				colMock.
					EXPECT().
					InsertOne(context.Background(), expectedBlogItem).
					Return(expectedRes, nil).
					Once()

				f.collection = colMock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(&tt.fields)
			s := &Server{
				collection: tt.fields.collection,
			}
			got, err := s.CreateBlog(tt.args.ctx, tt.args.req)
			if err != nil {
				e, _ := status.FromError(err)

				assert.Equal(t, tt.wantErrMsg, e.Message())
				return
			}
			assert.Equal(t, tt.want.Id, got.Id)
		})
	}
}

func TestServer_ReadBlog(t *testing.T) {
	type fields struct {
		collection CollectionInterface
	}
	type args struct {
		ctx context.Context
		req *proto.BlogId
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       *proto.Blog
		wantErrMsg string
		setup      func(f *fields)
	}{
		{
			name: "ShouldErrorOnGettingOID",
			args: args{
				ctx: context.Background(),
				req: &proto.BlogId{Id: "¤"},
			},
			wantErrMsg: "Cannot convert insertedID to OID - the provided hex string is not a valid ObjectID\n",
			setup:      func(f *fields) {},
		},
		{
			name: "ShouldErrorOnDecode",
			args: args{
				ctx: context.Background(),
				req: &proto.BlogId{Id: blogID},
			},
			wantErrMsg: "Cannot decode MongoDB response - foo bar decode\n",
			setup: func(f *fields) {
				collectionMock := collection.NewCollectionMock(t)

				collectionMock.
					EXPECT().
					FindOne(context.Background(), bson.M{"_id": blogOID}).
					Return(mongo.NewSingleResultFromDocument(db.BlogItem{}, fmt.Errorf("foo bar decode"), nil)).
					Once()

				f.collection = collectionMock
			},
		},
		{
			name: "ShouldReadBlog",
			args: args{
				ctx: context.Background(),
				req: &proto.BlogId{Id: blogID},
			},
			want: &proto.Blog{Id: blogID},
			setup: func(f *fields) {
				collectionMock := collection.NewCollectionMock(t)

				collectionMock.
					EXPECT().
					FindOne(context.Background(), bson.M{"_id": blogOID}).
					Return(mongo.NewSingleResultFromDocument(db.BlogItem{ID: blogOID}, nil, nil)).
					Once()

				f.collection = collectionMock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(&tt.fields)
			s := &Server{
				collection: tt.fields.collection,
			}
			got, err := s.ReadBlog(tt.args.ctx, tt.args.req)
			if err != nil {
				e, _ := status.FromError(err)
				assert.Equal(t, tt.wantErrMsg, e.Message())
				return
			}
			assert.Equal(t, true, reflect.DeepEqual(tt.want, got))
		})
	}
}

func TestServer_UpdateBlog(t *testing.T) {
	type fields struct {
		collection CollectionInterface
	}
	type args struct {
		ctx context.Context
		req *proto.Blog
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       *emptypb.Empty
		wantErrMsg string
		setup      func(f *fields)
	}{
		{
			name: "ShouldErrorOnWrongOID",
			args: args{
				ctx: context.Background(),
				req: &proto.Blog{
					Id:       wrongID,
					AuthorId: "Foo Author",
					Tile:     "Foo Title",
					Content:  "Foo Content",
				},
			},
			wantErrMsg: "Cannot convert insertedID to OID - the provided hex string is not a valid ObjectID\n",
			setup: func(f *fields) {
			},
		},
		{
			name: "ShouldErrorOnUpdate",
			args: args{
				ctx: context.Background(),
				req: &proto.Blog{
					Id:       blogID,
					AuthorId: "Foo Author",
					Tile:     "Foo Title",
					Content:  "Foo Content",
				},
			},
			wantErrMsg: "Could not update - foo update\n",
			setup: func(f *fields) {
				collectionMock := collection.NewCollectionMock(t)

				collectionMock.
					EXPECT().
					UpdateOne(
						context.Background(),
						bson.M{"_id": blogOID},
						bson.M{"$set": &db.BlogItem{
							AuthorID: "Foo Author",
							Title:    "Foo Title",
							Content:  "Foo Content",
						}},
					).
					Return(nil, fmt.Errorf("foo update")).
					Once()
				f.collection = collectionMock
			},
		},
		{
			name: "ShouldErrorOnNotMatched",
			args: args{
				ctx: context.Background(),
				req: &proto.Blog{
					Id:       blogID,
					AuthorId: "Foo Author",
					Tile:     "Foo Title",
					Content:  "Foo Content",
				},
			},
			wantErrMsg: "Cannot find blog with ID",
			setup: func(f *fields) {
				collectionMock := collection.NewCollectionMock(t)

				collectionMock.
					EXPECT().
					UpdateOne(
						context.Background(),
						bson.M{"_id": blogOID},
						bson.M{"$set": &db.BlogItem{
							AuthorID: "Foo Author",
							Title:    "Foo Title",
							Content:  "Foo Content",
						}},
					).
					Return(&mongo.UpdateResult{MatchedCount: 0}, nil).
					Once()
				f.collection = collectionMock
			},
		},
		{
			name: "ShouldUpdateBlog",
			args: args{
				ctx: context.Background(),
				req: &proto.Blog{
					Id:       blogID,
					AuthorId: "Foo Author",
					Tile:     "Foo Title",
					Content:  "Foo Content",
				},
			},
			want: &emptypb.Empty{},
			setup: func(f *fields) {
				collectionMock := collection.NewCollectionMock(t)

				collectionMock.
					EXPECT().
					UpdateOne(
						context.Background(),
						bson.M{"_id": blogOID},
						bson.M{"$set": &db.BlogItem{
							AuthorID: "Foo Author",
							Title:    "Foo Title",
							Content:  "Foo Content",
						}},
					).
					Return(&mongo.UpdateResult{MatchedCount: 1}, nil).
					Once()

				f.collection = collectionMock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(&tt.fields)
			s := &Server{
				collection: tt.fields.collection,
			}
			got, err := s.UpdateBlog(tt.args.ctx, tt.args.req)
			if err != nil {
				e, _ := status.FromError(err)
				assert.Equal(t, tt.wantErrMsg, e.Message())
				return
			}
			assert.Equal(t, true, reflect.DeepEqual(tt.want, got))
		})
	}
}

func TestServer_ListBlogs(t *testing.T) {
	type fields struct {
		collection CollectionInterface
	}
	type args struct {
		in0    *emptypb.Empty
		stream proto.BlogService_ListBlogsServer
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErrMsg string
		setup      func(f *fields, a *args)
	}{
		{
			name:       "ShouldErrorOnFind",
			wantErrMsg: "Could not find blogs - foo find\n",
			setup: func(f *fields, a *args) {
				collectionMock := collection.NewCollectionMock(t)

				collectionMock.
					EXPECT().
					Find(context.Background(), mock.Anything).
					Return(nil, fmt.Errorf("foo find")).
					Once()

				f.collection = collectionMock
			},
		},
		{
			name:       "ShouldErrorOnStream",
			wantErrMsg: "Cannot send data - foo stream\n",
			setup: func(f *fields, a *args) {
				streamMock := servermock.NewServerStreamMock(t)
				collectionMock := collection.NewCollectionMock(t)

				collectionMock.
					EXPECT().
					Find(context.Background(), mock.Anything).
					Return(mongo.NewCursorFromDocuments(
						[]interface{}{
							db.BlogItem{
								ID:       blogOID,
								AuthorID: "Foo Author",
								Title:    "Foo Title",
								Content:  "Foo Content",
							},
						}, nil, nil),
					).
					Once()

				streamMock.
					EXPECT().
					Send(mock.Anything).
					Return(fmt.Errorf("foo stream")).
					Once()

				a.stream = streamMock
				f.collection = collectionMock
			},
		},
		{
			name:       "ShouldErrorOnUnknownError",
			wantErrMsg: "Unknow error - foo unknown\n",
			setup: func(f *fields, a *args) {
				streamMock := servermock.NewServerStreamMock(t)
				collectionMock := collection.NewCollectionMock(t)

				collectionMock.
					EXPECT().
					Find(context.Background(), mock.Anything).
					Return(mongo.NewCursorFromDocuments(
						[]interface{}{
							db.BlogItem{
								ID:       blogOID,
								AuthorID: "Foo Author",
								Title:    "Foo Title",
								Content:  "Foo Content",
							},
						}, fmt.Errorf("foo unknown"), nil),
					).
					Once()

				a.stream = streamMock
				f.collection = collectionMock
			},
		},
		{
			name: "ShouldListBlogs",
			setup: func(f *fields, a *args) {
				streamMock := servermock.NewServerStreamMock(t)
				collectionMock := collection.NewCollectionMock(t)

				collectionMock.
					EXPECT().
					Find(context.Background(), mock.Anything).
					Return(mongo.NewCursorFromDocuments(
						[]interface{}{
							db.BlogItem{
								ID:       blogOID,
								AuthorID: "Foo Author",
								Title:    "Foo Title",
								Content:  "Foo Content",
							},
						}, nil, nil),
					).
					Once()

				streamMock.
					EXPECT().
					Send(mock.Anything).
					Return(nil).
					Once()

				a.stream = streamMock
				f.collection = collectionMock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(&tt.fields, &tt.args)
			s := &Server{
				collection: tt.fields.collection,
			}

			if err := s.ListBlogs(tt.args.in0, tt.args.stream); err != nil {
				e, _ := status.FromError(err)
				assert.Equal(t, tt.wantErrMsg, e.Message())
			}
		})
	}
}

func TestServer_DeleteBlog(t *testing.T) {
	type fields struct {
		collection CollectionInterface
	}
	type args struct {
		ctx context.Context
		req *proto.BlogId
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErrMsg string
		setup      func(f *fields)
	}{
		{
			name: "ShouldErrorOnWrongOID",
			args: args{
				ctx: context.Background(),
				req: &proto.BlogId{Id: wrongID},
			},
			wantErrMsg: "Cannot convert insertedID to OID - the provided hex string is not a valid ObjectID\n",
			setup: func(f *fields) {
			},
		},
		{
			name: "ShouldErrorOnDelete",
			args: args{
				ctx: context.Background(),
				req: &proto.BlogId{Id: blogID},
			},
			wantErrMsg: fmt.Sprintf("Cannot delete object under ID '%s' with error: %v", blogID, "foo delete"),
			setup: func(f *fields) {
				collectionMock := collection.NewCollectionMock(t)
				collectionMock.
					EXPECT().
					DeleteOne(context.Background(), mock.Anything).
					Return(nil, fmt.Errorf("foo delete")).
					Once()

				f.collection = collectionMock
			},
		},
		{
			name: "ShouldErrorOnDeleteCount",
			args: args{
				ctx: context.Background(),
				req: &proto.BlogId{Id: blogID},
			},
			wantErrMsg: fmt.Sprintf("Cannot find any objects with given ID '%s'", blogID),
			setup: func(f *fields) {
				collectionMock := collection.NewCollectionMock(t)
				collectionMock.
					EXPECT().
					DeleteOne(context.Background(), mock.Anything).
					Return(&mongo.DeleteResult{DeletedCount: 0}, nil).
					Once()

				f.collection = collectionMock
			},
		},
		{
			name: "ShouldDelete",
			args: args{
				ctx: context.Background(),
				req: &proto.BlogId{Id: blogID},
			},
			setup: func(f *fields) {
				collectionMock := collection.NewCollectionMock(t)
				collectionMock.
					EXPECT().
					DeleteOne(context.Background(), mock.Anything).
					Return(&mongo.DeleteResult{DeletedCount: 1}, nil).
					Once()

				f.collection = collectionMock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(&tt.fields)
			s := &Server{
				collection: tt.fields.collection,
			}

			if _, err := s.DeleteBlog(tt.args.ctx, tt.args.req); err != nil {
				e, _ := status.FromError(err)
				assert.Equal(t, tt.wantErrMsg, e.Message())
			}
		})
	}
}
