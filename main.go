package main

import (
	"log"

	"github.com/akshay-misra-demo-apps/go-cart/user-management/controllers"
	"github.com/akshay-misra-demo-apps/go-cart/user-management/middlewares"
	"github.com/akshay-misra-demo-apps/go-cart/user-management/repositories"
	"github.com/akshay-misra-demo-apps/go-cart/user-management/routes"
	"github.com/akshay-misra-demo-apps/go-cart/user-management/services"
	"github.com/gin-gonic/gin"
)

func init() {
	// check if this is good from DI perspective
	repositories.Init()

	// make grpc call to get all products api and cache all products to hzl
}

func App() {
	server := gin.Default()
	server.GET("/", home)

	userRepository := repositories.Get()["user"]
	userService := services.GetUserService(userRepository)
	userController := controllers.GetUserController(userService)

	routes.NoAuthRoutes(server, userController)

	// Authenticate all below routes
	server.Use(middlewares.Authenticate())
	// Authorize all below routes
	server.Use(middlewares.Authorize())
	routes.UserRoutes(server, userController)

	productRepository := repositories.Get()["product"]
	productService := services.GetProductService(productRepository)
	productController := controllers.GetProductController(productService)

	productGroup := routes.ProductRoutes(server, productController)

	routes.AuthorizedProductRoutes(productGroup, productController)

	if err := server.Run(":4000"); err != nil {
		log.Fatalf("error while creating http server %v", err.Error())
	}
}

// consider using some DI like google.wire to manage dependencies.
func main() {
	App()
}

func home(c *gin.Context) {
	c.JSON(200, gin.H{
		"response": "Welcome to go-cart!",
	})
}
