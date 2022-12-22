package main

import (
	"context"
	pb "example.com/go-productmgmt-grpc/productmgmt"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strconv"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	fmt.Println("Enter service IP Address or keep empty to use default: ")
	var ipAddress string
	_, err := fmt.Scanln(&ipAddress)
	if err != nil {
		log.Fatalf("could not establish host address")
	}

	if ipAddress == "" {
		ipAddress = address
		fmt.Printf("service IP address set as %v (default) \n", address)
	}

	conn, err := grpc.Dial(ipAddress, grpc.FailOnNonTempDialError(true), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("could not close client connection")
		}
	}(conn)

	userInterface(conn)
}

func userInterface(conn *grpc.ClientConn) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()
	var exit bool

	for exit != true {
		fmt.Printf(" 1 - find product by ID, \n 2 - create new product, \n 0 - exit program \n")
		var mode string
		_, err := fmt.Scanln(&mode)

		if err == nil {
			if mode == "1" {
				err = getProduct(ctx, conn)
			} else if mode == "2" {
				err = addProduct(ctx, conn)
			} else if mode == "0" {
				fmt.Println("Exiting program")
				exit = true
			} else {
				println("unknown command")
			}
		} else {
			log.Printf("unknown error: %v", err)
		}
	}
}

func addProduct(ctx context.Context, conn *grpc.ClientConn) error {
	c := pb.NewProductManagementClient(conn)

	var name string
	var value string

	fmt.Println("Enter product Name and Value: ")
	_, err := fmt.Scan(&name, &value)
	if err != nil {
		log.Fatalf("could not read input: %v", err)
	}
	valueFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Fatalf("could not convert value to float: %v", err)
	}

	// sending command to server
	prod, err := c.CreateProduct(ctx, &pb.NewProduct{Name: name, Value: valueFloat})
	if err != nil {
		log.Fatalf("could not create product: %v", err)
	}
	fmt.Println("New product successfully created!")
	log.Printf("\n Product Details: \n NAME: %s \n VALUE: %f \n ID: %d \n",
		prod.GetName(), prod.GetValue(), prod.GetId())

	return err
}

func getProduct(ctx context.Context, conn *grpc.ClientConn) error {
	c := pb.NewProductManagementClient(conn)

	var input string
	fmt.Println("Enter ID number")

	if _, err := fmt.Scan(&input); err != nil {
		log.Fatalf("could not read input: %v", err)
	}
	idInt, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		log.Fatalf("could not convert input %v", err)
	}
	id := int32(idInt)

	prod, err := c.GetProduct(ctx, &pb.Id{Id: id})
	if err != nil {
		log.Fatalf("a problem was ocured in getting product: %v", err)
	}

	if prod != nil {
		log.Printf("\n Product Details: \n ID: %d \n NAME: %s \n DESCR: %s \n VALUE: %f \n \n",
			prod.GetId(), prod.GetName(), prod.GetDescription(), prod.GetValue())
	} else {
		fmt.Printf("could not find product by id %d", id)
	}

	return err
}
