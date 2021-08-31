package main

import (
	"github.com/SevereCloud/vksdk/api"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	grpcService "mutual/internal/grpc"
	"mutual/pkg/config"
	grpc2 "mutual/pkg/grpc"
	"net"
	"net/http"
)

func main() {
	cfg, err := config.GetGrpcServerConfig()
	if err != nil {
		log.Fatal(err)
	}

	dsn := cfg.Db.Dsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("ошибка подключения базы данных: %s", err)
	}

	vk := api.NewVK(cfg.Vk.Token)

	mutualServer := grpcService.NewGrpcServer(db, vk)

	server := grpc.NewServer()

	grpc2.RegisterMutualServer(server, mutualServer)

	ln, err := net.Listen("tcp", ":"+cfg.Grpc.Port)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	if err := server.Serve(ln); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
