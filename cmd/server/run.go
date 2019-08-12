package main

import (
	"context"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/ykamez/sample-post-app/app/server"
)

func run() error {
	// Application context
	ctx := context.Background()

	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
			server.NewPostServiceServer(),
		),
	)
	return s.ServeContext(ctx)
}
