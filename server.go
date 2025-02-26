package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"pizza/pb"
)

type server struct {
	pb.UnimplementedPizzaServiceServer
	products map[string]*pb.ProductResponse
}

func newServer() *server {
	return &server{
		products: map[string]*pb.ProductResponse{
			"1": {ProductId: "1", Name: "Margherita", Price: 8.5, Category: pb.ProductCategory_PRODUCT_CATEGORY_PIZZA},
			"2": {ProductId: "2", Name: "Pepperoni", Price: 9.5, Category: pb.ProductCategory_PRODUCT_CATEGORY_PIZZA},
			"3": {ProductId: "3", Name: "Carbonara", Price: 10.0, Category: pb.ProductCategory_PRODUCT_CATEGORY_PASTA},
		},
	}
}

func (s *server) GetProduct(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	product, exists := s.products[req.GetProductId()]
	if !exists {
		return nil, grpc.Errorf(grpc.Code(pb.ProductCategory_PRODUCT_CATEGORY_UNSPECIFIED), "Product not found")
	}
	return product, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPizzaServiceServer(s, newServer())

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
