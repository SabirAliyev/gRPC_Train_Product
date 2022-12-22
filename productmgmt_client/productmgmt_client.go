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
	fmt.Println("Enter an IP Address or keep empty to use default: ")
	var ipAddress string
	fmt.Scanln(&ipAddress)

	if ipAddress == "" {
		ipAddress = address
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

	c := pb.NewProductManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()
	var exit bool

	for exit != true {
		fmt.Printf("1 - find product by ID, \n 2 - create new product")
		var mode string
		_, err := fmt.Scanln(&mode)

		if err == nil {
			if mode == "1" {
				// var id int32
				fmt.Println("Enter ID number")

			} else if mode == "2" {
				var name string
				var value string

				fmt.Println("Enter product Name and Value: ")
				_, err := fmt.Scan(&name, &value)
				if err != nil {
					log.Fatalf("could not add product: %v", err)
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
				log.Printf("Product Details: \n NAME: %s \n VALUE: %f \n ID: %d",
					prod.GetName(), prod.GetValue(), prod.GetId())
			} else if mode == "0" {
				exit = true
			} else {
				println("unknown command")
			}
		} else {
			log.Printf("unknown error: %v", err)
		}
	}
}
