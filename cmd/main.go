package main

import (
	"context"
	"fmt"
	"github.com/twitchtv/twirp"
	rpc "minimaltwirp/rpc/search"
	"net/http"
)

type search struct{}

func (search) Search(ctx context.Context, req *rpc.SearchRequest) (*rpc.SearchResponse, error) {
	message := fmt.Sprintf("Here's what I got from you %+v", req)
	fmt.Println("user_query", req.GetUserQuery())
	fmt.Println("repository_id", req.GetRepositoryId())
	fmt.Println()
	return &rpc.SearchResponse{Results: message}, nil
}

func main() {
	server := rpc.NewSearchServer(search{}, &twirp.ServerHooks{})

	addr := "localhost:1234"
	srv := &http.Server{
		Addr:    addr,
		Handler: server,
	}

	fmt.Println("Listening on", addr)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println(err)
	}
}
