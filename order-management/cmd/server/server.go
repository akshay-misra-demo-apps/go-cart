package server

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/akshay-misra-demo-apps/go-cart/order-management/api/pb"
	"github.com/akshay-misra-demo-apps/go-cart/order-management/controllers"
	"github.com/akshay-misra-demo-apps/go-cart/order-management/internal/service"
	"github.com/akshay-misra-demo-apps/go-cart/order-management/middlewares"
	"github.com/akshay-misra-demo-apps/go-cart/order-management/repositories"
	"github.com/akshay-misra-demo-apps/go-cart/order-management/routes"
	"github.com/akshay-misra-demo-apps/go-cart/order-management/services"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func App() {
	server := NewServer()

	// Start gRPC service in a goroutine
	go func() {
		//defer wg.Done()
		if err := server.StartGRPC(":50052"); err != nil {
			log.Fatalf("failed to start gRPC service: %v", err)
		}
		log.Println("grpc service started.")
	}()

	// Start HTTP service in a goroutine
	go func() {
		//defer wg.Done()
		if err := server.StartHTTP(":4300"); err != nil {
			log.Fatalf("failed to start HTTP service: %v", err)
		}
	}()

	// Listen for interrupt signals
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case <-interrupt:
		log.Println("Received interrupt signal, shutting down...")
		server.Stop()
		log.Println("Servers gracefully stopped.")
	}
}

type Server struct {
	grpcServer *grpc.Server
	httpServer *gin.Engine
}

func NewServer() *Server {
	// Create gRPC service
	grpcServer := Grpc()

	// Create HTTP service
	httpServer := Rest()

	return &Server{
		grpcServer: grpcServer,
		httpServer: httpServer,
	}
}

// StartGRPC run following command to check connection: nc -zv localhost 50051
func (s *Server) StartGRPC(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	log.Println("starting grpc service on ", address)
	return s.grpcServer.Serve(lis)
}

func (s *Server) StartHTTP(address string) (err error) {
	return s.httpServer.Run(address)
}

func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
}

func Grpc() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, service.InitOrderService())
	pb.RegisterCartServiceServer(s, service.InitCartService())
	return s
}

func Rest() *gin.Engine {
	s := gin.Default()
	s.GET("/", home)

	// Authenticate all below routes
	s.Use(middlewares.Authenticate())
	// Authorize all below routes
	s.Use(middlewares.Authorize())

	productRepository := repositories.Get()["product"]
	productService := services.GetProductService(productRepository)
	productController := controllers.GetProductController(productService)

	productGroup := routes.ProductRoutes(s, productController)
	routes.AuthorizedProductRoutes(productGroup, productController)

	return s
}

func home(c *gin.Context) {
	c.JSON(200, gin.H{
		"response": "Welcome to go-cart!",
	})
}

// grpc client code to make remote func call
func pingGRPC() {
	log.Println("ping grpc service started.")
	// Set up a connection to the service
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
		return
	}
	defer conn.Close()

	// Create a client
	client := pb.NewProductServiceClient(conn)

	// Set a timeout for the connection
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	// Try to ping the service
	products, err := client.ReadAllProducts(ctx, &pb.ReadAllProductsRequest{
		Limit: 5,
		Sort:  "-1",
	})

	if err != nil {
		log.Printf("Could not ping service: %v", err)
		return
	}

	if jsonProducts, err := json.Marshal(products); err == nil {
		log.Println("got products using grpc", string(jsonProducts))
	}

	log.Println("grpc service is running")
}
