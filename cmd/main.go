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
	return &rpc.SearchResponse{Results: fmt.Sprintf("Here's what I got from you %+v", req)}, nil
}

func main() {
	server := rpc.NewSearchServer(search{}, &twirp.ServerHooks{})

	srv := &http.Server{
		Addr:    "localhost:1234",
		Handler: server,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println(err)
	}
}
