package di

import (
	"database/sql"

	"github.com/srvc/fail"
	"github.com/ykamez/sample-post-app/infra/store"
	"github.com/ykamez/sample-post-app/infra/store/rdb"
)

// AppComponent is a DI container of notifications app.
type AppComponent interface {
	PostStore() store.PostStore
}

// CreateAppComponent initializes a new AppComponent instnace.
func CreateAppComponent() (AppComponent, error) {
	db, err := sql.Open("sample_post_app", "postgres://postgres:@localhost/sample_post_app?sslmode=disable")
	if err != nil {
		return nil, fail.Wrap(err)
	}
	db.SetMaxIdleConns(40)
	db.SetMaxOpenConns(40)

	return &appComponentImpl{
		DB: db,
	}, nil
}

type appComponentImpl struct {
	DB        *sql.DB
	postStore store.PostStore
}

func (c *appComponentImpl) PostStore() store.PostStore {
	return rdb.NewPostStore(c.DB)
}
