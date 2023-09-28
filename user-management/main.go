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
}

func App() {
	server := gin.Default()
	server.GET("/", home)

	userRepository := repositories.Get()["user"]
	userService := services.GetUserService(userRepository)
	userController := controllers.GetUserController(userService)

	// register routes without authentication
	routes.NoAuthRoutes(server, userController)
	// Authenticate all below routes
	server.Use(middlewares.Authenticate())
	// Authorize all below routes
	server.Use(middlewares.Authorize())
	routes.UserRoutes(server, userController)

	if err := server.Run(":4000"); err != nil {
		log.Fatalf("error while creating http service %v", err.Error())
	}
}

// Consider using some DI like google.wire to manage dependencies.
func main() {
	App()
}

func home(c *gin.Context) {
	c.JSON(200, gin.H{
		"response": "Welcome to go-cart user management service!",
	})
}
