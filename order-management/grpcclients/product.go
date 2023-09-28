package grpcclients

import (
	"log"

	"github.com/akshay-misra-demo-apps/go-cart/order-management/api/pb"
	"google.golang.org/grpc"
)

// Product declare a package-level variable to hold the client
var Product pb.ProductServiceClient

// InitProductClient initialize the product gRPC client
func InitProductClient(address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to product-service: %v", err)
	}
	Product = pb.NewProductServiceClient(conn)
}
