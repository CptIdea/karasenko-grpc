package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"mutual/pkg/config"
	grpc2 "mutual/pkg/grpc"
	"net/http"
)

func main() {
	cfg, err := config.GetHttpServerConfig()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = grpc2.RegisterMutualHandlerFromEndpoint(ctx, mux, "localhost:"+cfg.Grpc.Port, opts)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":"+cfg.Http.Port, mux))

}
