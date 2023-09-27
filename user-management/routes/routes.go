package routes

import (
	"github.com/akshay-misra-demo-apps/go-cart/user-management/controllers"
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
