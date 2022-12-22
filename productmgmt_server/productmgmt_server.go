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
	}
}

type ProductManagementServer struct {
	pb.UnimplementedProductManagementServer
}

func (s *ProductManagementServer) CreateProduct(ctx context.Context, prod *pb.NewProduct) (*pb.Product, error) {
	log.Printf("Recieved new product: %v", prod.GetName())

	product := models.Product{
		Name:  prod.Name,
		Value: prod.Value,
	}

	id, err := appg.insertProduct(&product)
	if err != nil {
		log.Fatalf("could not create product: %v", err)
	}

	return &pb.Product{Name: prod.GetName(), Value: prod.GetValue(), Id: id}, nil
}

func (app *application) insertProduct(product *models.Product) (int32, error) {
	id, err := app.products.Insert(product)
	if err != nil {
		log.Fatalf("could not insert product %v", err)
	}

	return id, err
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
