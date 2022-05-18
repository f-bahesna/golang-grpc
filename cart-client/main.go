package main

import (
	"context"
	"fmt"
	pb "github.com/f-bahesna/grpc-client/cart"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var sku string
	fmt.Print("Write your sku: ")
	fmt.Scanf("%s", &sku)

	conn, err := grpc.Dial("localhost:1500", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error when dial %v", err)
	}

	client := pb.NewCartClient(conn)

	req := pb.CartRequest{Sku: sku}

	cart, err := client.FindCartBySku(context.Background(), &req)
	if err != nil {
		log.Fatalf("error when find cart %v", err)
	}

	fmt.Println(cart)
}
