package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/akshay-misra-demo-apps/go-cart/order-management/api/pb"
	"github.com/akshay-misra-demo-apps/go-cart/order-management/cmd/server"
	"github.com/akshay-misra-demo-apps/go-cart/order-management/grpcclients"
	"github.com/akshay-misra-demo-apps/go-cart/order-management/repositories"
)

func init() {
	// check if this is good from DI perspective
	repositories.Init()
	// Initialize product gRPC client
	grpcclients.InitProductClient(":9090")

	// make service call to get all products api and cache all products to hzl
}

// consider using some DI like google.wire to manage dependencies.
func main() {
	// Set a timeout for the connection
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	// make grpc call to product service and cache all products to hzl
	products, err := grpcclients.Product.ReadAllProducts(ctx, &pb.ReadAllProductsRequest{
		Limit: 5,
		Sort:  "-1",
	})
	if err != nil {
		log.Printf("Could not ping service: %v", err)
		return
	}

	if jsonProducts, err := json.Marshal(products); err == nil {
		log.Println("got products using grpc:: ", string(jsonProducts))
	}
	server.App()
}
