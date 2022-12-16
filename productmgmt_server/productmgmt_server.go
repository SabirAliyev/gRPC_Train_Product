package main

import (
	"context"
	pb "example.com/go-productmgmt-grpc/productmgmt"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
)

const (
	address = ":50051"
)

type ProductManagementServer struct {
	pb.UnimplementedProductManagementServer
}

func (s *ProductManagementServer) CreateProduct(ctx context.Context, in *pb.NewProduct) (*pb.Product, error) {
	log.Printf("Recieved: %v", in.GetName())

	var productId = int32(rand.Intn(1000))

	return &pb.Product{Name: in.GetName(), Value: in.GetValue(), Id: productId}, nil
}

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterProductManagementServer(s, &ProductManagementServer{})
	log.Printf("sever listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
