package main

import (
	"context"
	"example.com/go-productmgmt-grpc/models"
	pb "example.com/go-productmgmt-grpc/productmgmt"
	"log"
)

type ProductManagementServer struct {
	pb.UnimplementedProductManagementServer
}

func (s *ProductManagementServer) CreateProduct(ctx context.Context, prod *pb.NewProduct) (*pb.Product, error) {
	log.Printf("Recieved new product: %v", prod.GetName())

	product := models.Product{
		Name:        prod.Name,
		Description: prod.Description,
		Value:       prod.Value,
	}

	id, err := appg.products.Insert(&product)
	if err != nil {
		log.Fatalf("could not insert product %v", err)
	}

	// return protobuf Product
	return &pb.Product{Name: prod.GetName(), Value: prod.GetValue(), Id: id}, nil
}

func (s *ProductManagementServer) GetProduct(ctx context.Context, id *pb.Id) (*pb.Product, error) {
	log.Printf("Recieved product ID: %v", id)

	var sendId = id.GetId()

	prod, err := appg.products.FindById(sendId)
	if err != nil {
		log.Fatalf("Error was occured on getting product by Id: (Id number: %v). %v", id, err)
	}

	// return protobuf Product
	return &pb.Product{Id: prod.Id, Name: prod.Name, Description: prod.Description, Value: prod.Value}, nil
}
