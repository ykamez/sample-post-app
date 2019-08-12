package rdb

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/wantedly/notifications/infra/record"
	"github.com/ykamez/sample-post-app/infra/store"
)

// NewPostStore creates a new store.PostStore instance.
func NewPostStore(db *sql.DB) store.PostStore {
	return &postStoreImpl{db: db}
}

type postStoreImpl struct {
	db *sql.DB
}

func (s *subscriptionStoreImpl) GetPosts(ctx context.Context) ([]*record.Post, error) {
	posts, err := record.Posts(qm.All(ctx, s.db)
	if err != nil && !isErrNoRows(err) {
		return nil, fail.Wrap(err)
	}

	return posts, nil
}

func (s *subscriptionStoreImpl) Create(ctx context.Context) ( error) {
	err = (&record.Post{title: "Hello, world!"}).Insert(ctx, tx, boil.Infer())
	if err != nil && !isErrNoRows(err) {
		return nil, fail.Wrap(err)
	}

	return nil
}
