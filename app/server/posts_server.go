package server

import (
	"context"
	"fmt"

	"github.com/izumin5210/grapi/pkg/grapiserver"

	api_pb "github.com/ykamez/sample-post-app/api"
)

// PostServiceServer is a composite interface of api_pb.PostServiceServer and grapiserver.Server.
type PostServiceServer interface {
	api_pb.PostServiceServer
	grapiserver.Server
}

// NewPostServiceServer creates a new PostServiceServer instance.
func NewPostServiceServer() PostServiceServer {
	return &postServiceServerImpl{}
}

type postServiceServerImpl struct {
}

func (s *postServiceServerImpl) ListPosts(ctx context.Context, req *api_pb.ListPostsRequest) (*api_pb.ListPostsResponse, error) {
	fmt.Println("ListPost called")
	resp := &api_pb.ListPostsResponse{
		// Posts:
	}

	return resp, nil
}

func (s *postServiceServerImpl) CreatePost(ctx context.Context, req *api_pb.CreatePostsRequest) (*api_pb.Post, error) {

	fmt.Println("CreatePost called")
	resp := &api_pb.Post{
		// Posts: util.PostsToPb(subs, userSubs, appMeta.Language),
	}

	return resp, nil
}
