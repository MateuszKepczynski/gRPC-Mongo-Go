package client

import (
	"context"
	"fmt"
	"github.com/grpc-mongo-go/gen/proto"
	clientprotomock "github.com/grpc-mongo-go/mocks/grpc/client"
	clientstreamprotomock "github.com/grpc-mongo-go/mocks/grpc/client/stream"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"testing"
)

const blogID = "65b9018f4e0d763afee645f8"

func TestCreateBlog(t *testing.T) {
	type args struct {
		ctx context.Context
		c   proto.BlogServiceClient
	}
	tests := []struct {
		name       string
		args       args
		want       string
		wantErrMsg string
		setup      func(a *args)
	}{
		{
			name:       "ShouldErrorOnCreateBlog",
			args:       args{ctx: context.Background()},
			wantErrMsg: "foo create",
			setup: func(a *args) {
				clientMock := clientprotomock.NewClientProtoMock(t)
				clientMock.
					EXPECT().
					CreateBlog(a.ctx, &proto.Blog{
						AuthorId: "Matthew",
						Tile:     "Hello",
						Content:  "World",
					}).
					Return(nil, fmt.Errorf("foo create")).
					Once()

				a.c = clientMock
			},
		},
		{
			name: "ShouldCreateBlog",
			args: args{ctx: context.Background()},
			want: blogID,
			setup: func(a *args) {
				clientMock := clientprotomock.NewClientProtoMock(t)
				clientMock.
					EXPECT().
					CreateBlog(a.ctx, &proto.Blog{
						AuthorId: "Matthew",
						Tile:     "Hello",
						Content:  "World",
					}).
					Return(&proto.BlogId{Id: blogID}, nil).
					Once()

				a.c = clientMock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(&tt.args)
			got, err := CreateBlog(tt.args.ctx, tt.args.c)
			if err != nil {
				assert.Errorf(t, err, tt.wantErrMsg)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestReadBlog(t *testing.T) {
	type args struct {
		ctx context.Context
		c   proto.BlogServiceClient
		id  string
	}
	tests := []struct {
		name       string
		args       args
		want       *proto.Blog
		wantErrMsg string
		setup      func(a *args)
	}{
		{
			name: "ShouldErrorOnRead",
			args: args{
				ctx: context.Background(),
				id:  blogID,
			},
			wantErrMsg: "foo read",
			setup: func(a *args) {
				clientMock := clientprotomock.NewClientProtoMock(t)
				clientMock.
					EXPECT().
					ReadBlog(a.ctx, &proto.BlogId{Id: blogID}).
					Return(nil, fmt.Errorf("foo read")).
					Once()

				a.c = clientMock
			},
		},
		{
			name: "ShouldRead",
			args: args{
				ctx: context.Background(),
				id:  blogID,
			},
			want: &proto.Blog{
				Id:       blogID,
				AuthorId: "Foo Author",
				Tile:     "Foo Title",
				Content:  "Foo Content",
			},
			setup: func(a *args) {
				clientMock := clientprotomock.NewClientProtoMock(t)
				clientMock.
					EXPECT().
					ReadBlog(a.ctx, &proto.BlogId{Id: blogID}).
					Return(&proto.Blog{
						Id:       blogID,
						AuthorId: "Foo Author",
						Tile:     "Foo Title",
						Content:  "Foo Content",
					}, nil).
					Once()

				a.c = clientMock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(&tt.args)
			got, err := ReadBlog(tt.args.ctx, tt.args.c, tt.args.id)
			if err != nil {
				assert.Errorf(t, err, tt.wantErrMsg)
				return
			}
			assert.Equal(t, tt.want.Id, got.Id)
		})
	}
}

func TestUpdateBlog(t *testing.T) {
	type args struct {
		ctx context.Context
		c   proto.BlogServiceClient
		b   *proto.Blog
	}
	tests := []struct {
		name       string
		args       args
		wantErrMsg string
		setup      func(a *args)
	}{
		{
			name: "ShouldErrorOnUpdate",
			args: args{
				ctx: context.Background(),
				b: &proto.Blog{
					Id:       blogID,
					AuthorId: "Foo Author",
					Tile:     "Foo Title",
					Content:  "Foo Content",
				},
			},
			wantErrMsg: "foo update",
			setup: func(a *args) {
				clientMock := clientprotomock.NewClientProtoMock(t)
				clientMock.
					EXPECT().
					UpdateBlog(a.ctx, a.b).
					Return(nil, fmt.Errorf("foo update")).
					Once()

				a.c = clientMock
			},
		},
		{
			name: "ShouldUpdate",
			args: args{
				ctx: context.Background(),
				b: &proto.Blog{
					Id:       blogID,
					AuthorId: "Foo Author",
					Tile:     "Foo Title",
					Content:  "Foo Content",
				},
			},
			setup: func(a *args) {
				clientMock := clientprotomock.NewClientProtoMock(t)
				clientMock.
					EXPECT().
					UpdateBlog(a.ctx, a.b).
					Return(nil, nil).
					Once()

				a.c = clientMock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(&tt.args)

			if err := UpdateBlog(tt.args.ctx, tt.args.c, tt.args.b); err != nil {
				assert.Errorf(t, err, tt.wantErrMsg)
			}
		})
	}
}

func TestListBlogs(t *testing.T) {
	type args struct {
		ctx context.Context
		c   proto.BlogServiceClient
	}
	tests := []struct {
		name       string
		args       args
		wantErrMsg string
		setup      func(a *args)
	}{
		{
			name:       "ShouldErrorOnList",
			args:       args{ctx: context.Background()},
			wantErrMsg: "foo list",
			setup: func(a *args) {
				clientMock := clientprotomock.NewClientProtoMock(t)

				clientMock.
					EXPECT().
					ListBlogs(a.ctx, &emptypb.Empty{}).
					Return(nil, fmt.Errorf("foo list")).
					Once()

				a.c = clientMock
			},
		},
		{
			name:       "ShouldErrorOnStream",
			args:       args{ctx: context.Background()},
			wantErrMsg: "foo list",
			setup: func(a *args) {
				clientMock := clientprotomock.NewClientProtoMock(t)
				streamMock := clientstreamprotomock.NewClientStreamProtoMock(t)

				clientMock.
					EXPECT().
					ListBlogs(a.ctx, &emptypb.Empty{}).
					Return(streamMock, nil).
					Once()

				streamMock.
					EXPECT().
					Recv().
					Return(&proto.Blog{
						Id:       blogID,
						AuthorId: "Foo Author",
						Tile:     "Foo Title",
						Content:  "Foo Content",
					}, nil).
					Once().
					On("Recv").
					Return(nil, fmt.Errorf("foo list")).
					Once()

				a.c = clientMock
			},
		},
		{
			name: "ShouldList",
			args: args{ctx: context.Background()},
			setup: func(a *args) {
				clientMock := clientprotomock.NewClientProtoMock(t)
				streamMock := clientstreamprotomock.NewClientStreamProtoMock(t)

				clientMock.
					EXPECT().
					ListBlogs(a.ctx, &emptypb.Empty{}).
					Return(streamMock, nil).
					Once()

				streamMock.
					EXPECT().
					Recv().
					Return(&proto.Blog{
						Id:       blogID,
						AuthorId: "Foo Author",
						Tile:     "Foo Title",
						Content:  "Foo Content",
					}, nil).
					Once().
					On("Recv").
					Return(nil, io.EOF).
					Once()

				a.c = clientMock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(&tt.args)

			if err := ListBlogs(tt.args.ctx, tt.args.c); err != nil {
				assert.Errorf(t, err, tt.wantErrMsg)
			}
		})
	}
}

func TestDeleteBlog(t *testing.T) {
	type args struct {
		ctx context.Context
		c   proto.BlogServiceClient
		id  string
	}
	tests := []struct {
		name       string
		args       args
		wantErrMsg string
		setup      func(a *args)
	}{
		{
			name: "ShouldErrorOnDelete",
			args: args{
				ctx: context.Background(),
				id:  blogID,
			},
			wantErrMsg: "foo delete",
			setup: func(a *args) {
				clientMock := clientprotomock.NewClientProtoMock(t)
				clientMock.
					EXPECT().
					DeleteBlog(a.ctx, &proto.BlogId{Id: a.id}).
					Return(nil, fmt.Errorf("foo delte")).
					Once()

				a.c = clientMock
			},
		},
		{
			name: "ShouldDelete",
			args: args{
				ctx: context.Background(),
				id:  blogID,
			},
			setup: func(a *args) {
				clientMock := clientprotomock.NewClientProtoMock(t)
				clientMock.
					EXPECT().
					DeleteBlog(a.ctx, &proto.BlogId{Id: a.id}).
					Return(nil, nil).
					Once()

				a.c = clientMock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(&tt.args)

			if err := DeleteBlog(tt.args.ctx, tt.args.c, tt.args.id); err != nil {
				assert.Errorf(t, err, tt.wantErrMsg)
			}
		})
	}
}
