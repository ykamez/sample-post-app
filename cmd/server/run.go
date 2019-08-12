package main

import (
	"context"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/ykamez/sample-post-app/app/di"
	"github.com/ykamez/sample-post-app/app/server"
)

func run() error {
	// Application context
	ctx := context.Background()
	app, err := di.CreateAppComponent()

	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
			server.NewPostServiceServer(app.PostStore),
		),
	)
	return s.ServeContext(ctx)
}
