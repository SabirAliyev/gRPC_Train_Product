package main

import (
	"context"
	"example.com/go-productmgmt-grpc/models"
	"example.com/go-productmgmt-grpc/models/postgresql"
	pb "example.com/go-productmgmt-grpc/productmgmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

const (
	address = ":50051"
)

var appg application

type application struct {
	products interface {
		Insert(product *models.Product) (int32, error)
		FindById(int322 int32) (models.Product, error)
	}
}

func main() {
	app := &application{}

	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("could not initiate database: %v", err)
	}
	defer pool.Close()

	app.products = &postgresql.ProductModel{Pool: pool}

	appg = *app

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterProductManagementServer(s, &ProductManagementServer{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
