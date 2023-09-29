package server

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/akshay-misra-demo-apps/go-cart/product-management/api/pb"
	"github.com/akshay-misra-demo-apps/go-cart/product-management/controllers"
	"github.com/akshay-misra-demo-apps/go-cart/product-management/internal/service"
	"github.com/akshay-misra-demo-apps/go-cart/product-management/middlewares"
	"github.com/akshay-misra-demo-apps/go-cart/product-management/repositories"
	"github.com/akshay-misra-demo-apps/go-cart/product-management/routes"
	"github.com/akshay-misra-demo-apps/go-cart/product-management/services"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func App() {
	server := NewServer()

	// Start gRPC service in a goroutine
	go func() {
		//defer wg.Done()
		if err := server.StartGRPC(":9090"); err != nil {
			log.Fatalf("failed to start gRPC service: %v", err)
		}
		log.Println("grpc service started.")
	}()

	// Start HTTP service in a goroutine
	go func() {
		//defer wg.Done()
		if err := server.StartHTTP(":4200"); err != nil {
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
	pb.RegisterProductServiceServer(s, service.InitProductService())
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
