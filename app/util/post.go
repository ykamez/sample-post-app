package util

import (
	"context"

	api_pb "github.com/ykamez/sample-post-app/api"
	"github.com/ykamez/sample-post-app/infra/record"
)

// PostsToPb converts post objects to pb objects.
func PostsToPb(ctx context.Context, posts []*record.Post) []*api_pb.Post {
	pbs := make([]*api_pb.Post, len(posts))
	for i, m := range posts {
		pbs[i] = PostToPb(ctx, m)
	}
	return pbs
}

// PostToPb converts post object to pb object.
func PostToPb(ctx context.Context, p *record.Post) *api_pb.Post {
	pbp := &api_pb.Post{
		Id:    uint32(p.ID),
		Title: p.Title.String,
	}
	return pbp
}
