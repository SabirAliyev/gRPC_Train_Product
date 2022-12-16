package main

import (
	"context"
	pb "example.com/go-productmgmt-grpc/productmgmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("could not close client connection")
		}
	}(conn)

	c := pb.NewProductManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var newProduct = make(map[string]float64)
	newProduct["Macbook Pro"] = 2500.99
	newProduct["Lenovo ideaPad"] = 1200.99

	for name, value := range newProduct {
		r, err := c.CreateProduct(ctx, &pb.NewProduct{Name: name, Value: value})
		if err != nil {
			log.Fatalf("could not create product %v", err)
		}

		log.Printf(`Product Details:
NAME: %s
VALUE: %f
ID: %d`, r.GetName(), r.GetValue(), r.GetId())
	}
}
