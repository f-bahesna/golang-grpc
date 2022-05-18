package main

import (
	"context"
	pb "github.com/f-bahesna/grpc-service/cart"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type cartService struct {
	pb.UnimplementedCartServer
	carts []*pb.CartData
}

func (c *cartService) FindCartBySku(ctx context.Context, in *pb.CartRequest) (*pb.CartData, error) {
	for _, cart := range c.carts {
		if cart.Sku == in.Sku {
			return cart, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "cart not found")
}

func newCartService() *cartService {
	var carts []*pb.CartData
	carts = append(carts, &pb.CartData{Sku: "1", Name: "Marvel Book", Description: "This is Marvel Book for you"})
	carts = append(carts, &pb.CartData{Sku: "2", Name: "Tarzan Book", Description: "This is Tarzan Book for you"})

	cs := &cartService{}
	cs.carts = carts
	return cs
}

func main() {
	lis, err := net.Listen("tcp", ":1500")
	if err != nil {
		log.Fatalf("error when starting %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCartServer(s, newCartService())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("error when serve %v", err)
	}
}
