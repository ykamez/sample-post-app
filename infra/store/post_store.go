package store

import (
	"context"

	"github.com/ykamez/sample-post-app/infra/record"
)

// PostStore is an interface for CRUD post records
type PostStore interface {
	GetPosts(ctx context.Context) ([]*record.Post, error)
	Create(ctx context.Context, n *record.Post) error
}
