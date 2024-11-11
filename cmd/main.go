package main

import (
	"context"
	"log"
	"net"

	"github.com/justIGreK/MonkeyKeeper-User/cmd/handler"
	"github.com/justIGreK/MonkeyKeeper-User/internal/repository"
	"github.com/justIGreK/MonkeyKeeper-User/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()

	db := repository.CreateMongoClient(ctx)
	userDB := repository.NewUserRepository(db)
	userSRV := service.NewUserService(userDB)
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	handler := handler.NewHandler(grpcServer, userSRV)
	handler.RegisterServices()
	reflection.Register(grpcServer)

	log.Printf("Starting gRPC server on :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
