package routes

import (
	"github.com/akshay-misra-demo-apps/go-cart/product-management/controllers"
	"github.com/gin-gonic/gin"
)

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
