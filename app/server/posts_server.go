package server

import (
	"context"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
