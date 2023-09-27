package controllers

import (
	"fmt"
	"strings"

	"github.com/akshay-misra-demo-apps/go-cart/user-management/interfaces"
	"github.com/akshay-misra-demo-apps/go-cart/user-management/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	Service interfaces.IUser
}

func GetUserController(service interfaces.IUser) *UserController {
	return &UserController{
		Service: service,
	}
}

func (u UserController) Register(c *gin.Context) {
	fmt.Println("register route")

	var user models.User
	c.Bind(&user)

	user.Id = primitive.NewObjectID()

	created, err := u.Service.Register(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("internal server error: %v", err.Error()),
		})
		return
	}

	c.JSON(201, gin.H{
		"data": created,
	})
}

func (u UserController) Login(c *gin.Context) {
	fmt.Println("login route")
	var user models.Login
	c.Bind(&user)

	auth, err := u.Service.Authenticate(&user)
	if err == nil {
		c.JSON(200, gin.H{
			"data": auth,
		})
		return
	}

	if strings.Contains(err.Error(), "not found") {
		c.JSON(404, gin.H{
			"message": fmt.Sprintf("%v", err.Error()),
		})
	} else {
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("internal server error: %v", err.Error()),
		})
	}
}

func (u UserController) Logout(c *gin.Context) {
	fmt.Println("logout route")
}

func (u UserController) GetUser(c *gin.Context) {
	fmt.Println("profile with id: ", c.Param("id"))
}

func (u UserController) PatchUser(c *gin.Context) {
	fmt.Println("patch existing user", c.Param("id"))

	if c.Param("id") == "" {
		c.JSON(400, gin.H{
			"message": "user 'id' is missing in request!",
		})
		return
	}

	var user models.User
	c.Bind(&user)

	objectId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("internal server error: %v", err.Error()),
		})
		return
	}
	user.Id = objectId

	updated, err := u.Service.Patch(&user)
	if err == nil {
		c.JSON(200, gin.H{
			"data": updated,
		})
		return
	}

	if strings.Contains(err.Error(), "not found") {
		c.JSON(404, gin.H{
			"message": fmt.Sprintf("%v", err.Error()),
		})
	} else {
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("internal server error: %v", err.Error()),
		})
	}
}

func (u UserController) DeleteUser(c *gin.Context) {
	fmt.Println("profile with id: ", c.Param("id"))
}
