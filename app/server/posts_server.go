package server

import (
	"context"
	"fmt"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/srvc/fail"

	api_pb "github.com/ykamez/sample-post-app/api"
	"github.com/ykamez/sample-post-app/app/util"
	"github.com/ykamez/sample-post-app/infra/store"
)

// PostServiceServer is a composite interface of api_pb.PostServiceServer and grapiserver.Server.
type PostServiceServer interface {
	api_pb.PostServiceServer
	grapiserver.Server
}

// NewPostServiceServer creates a new PostServiceServer instance.
func NewPostServiceServer(
	store store.PostStore,
) interface {
	api_pb.PostServiceServer
	grapiserver.Server
} {
	return &postServiceServerImpl{
		store: store,
	}
}

type postServiceServerImpl struct {
	store store.PostStore
}

func (s *postServiceServerImpl) ListPosts(ctx context.Context, req *api_pb.ListPostsRequest) (*api_pb.ListPostsResponse, error) {
	fmt.Println("ListPost called")
	posts, err := s.store.GetPosts(ctx)
	if err != nil {
		return nil, fail.Wrap(err)
	}
	resp := &api_pb.ListPostsResponse{
		Posts: util.PostsToPb(ctx, posts),
	}

	return resp, nil
}

func (s *postServiceServerImpl) CreatePost(ctx context.Context, req *api_pb.CreatePostsRequest) (*api_pb.Post, error) {

	fmt.Println("CreatePost called")
	// post, err := s.store.Create(ctx)
	resp := &api_pb.Post{
		// Posts: util.PostsToPb(subs, userSubs, appMeta.Language),
	}

	return resp, nil
}
