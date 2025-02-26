package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/codevibe-de/goadv--products/product/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedProductServiceServer
}

func (s *server) ListProducts(in *pb.ProductListRequest, stream pb.ProductService_ListProductsServer) error {
	for i := 0; i < 5; i++ {
		if err := stream.Send(&pb.ProductResponse{
			ProductId: fmt.Sprintf("P-%03d", i),
			Name:      fmt.Sprintf("Pizza %d", i),
			Price:     7.99 + float64(i),
		}); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) GetProduct(ctx context.Context, in *pb.ProductRequest) (*pb.ProductResponse, error) {
	return &pb.ProductResponse{
		ProductId: in.ProductId,
		Name:      "Pizza",
		Price:     7.99,
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
