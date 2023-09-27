package routes

import (
	"github.com/akshay-misra-demo-apps/go-cart/controllers"
	"github.com/gin-gonic/gin"
)

func NoAuthRoutes(r *gin.Engine, c *controllers.UserController) *gin.RouterGroup {
	// user management routes
	user := r.Group("/auth/user")

	user.POST("/register", c.Register)
	user.POST("/login", c.Login)
	user.POST("/logout", c.Logout)

	return user
}

func UserRoutes(r *gin.Engine, c *controllers.UserController) *gin.RouterGroup {
	// user management routes
	user := r.Group("/api/user")
	user.GET("/:id", c.GetUser)
	user.PATCH("/:id", c.PatchUser)
	user.DELETE("/:id", c.DeleteUser)

	return user
}

func ProductRoutes(r *gin.Engine, c *controllers.ProductController) *gin.RouterGroup {
	// product catalog management routes
	product := r.Group("/api/product")
	product.GET("/", c.GetProducts)
	product.GET("/:id", c.GetProduct)

	return product
}

func AuthorizedProductRoutes(product *gin.RouterGroup, c *controllers.ProductController) *gin.RouterGroup {
	// product catalog management routes
	product.POST("/", c.CreateProduct)
	product.PATCH("/:id", c.PatchProduct)
	product.DELETE("/:id", c.DeleteProduct)

	return product
}
