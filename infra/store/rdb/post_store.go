package rdb

import (
	"context"
	"database/sql"

	"github.com/srvc/fail"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/ykamez/sample-post-app/infra/record"
	"github.com/ykamez/sample-post-app/infra/store"
)

// NewPostStore creates a new store.PostStore instance.
func NewPostStore(db *sql.DB) store.PostStore {
	return &postStoreImpl{db: db}
}

type postStoreImpl struct {
	db *sql.DB
}

func (s *postStoreImpl) GetPosts(ctx context.Context) ([]*record.Post, error) {
	posts, err := record.Posts(qm.Where("id IN (1,2,3,4,5)")).All(ctx, s.db)
	if err != nil {
		return nil, fail.Wrap(err)
	}

	return posts, nil
}

// 作成したPostを返したい
// func (s *postStoreImpl) Create(ctx context.Context) error {
// 	err := (&record.Post{title: "Hello, world!"}).Insert(ctx, tx, boil.Infer())
// 	if err != nil && !isErrNoRows(err) {
// 		return nil, fail.Wrap(err)
// 	}

// 	return nil
// }
