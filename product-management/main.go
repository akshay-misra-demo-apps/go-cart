package main

import (
	"github.com/akshay-misra-demo-apps/go-cart/product-management/cmd/server"
	"github.com/akshay-misra-demo-apps/go-cart/product-management/repositories"
)

func init() {
	// check if this is good from DI perspective
	repositories.Init()

	// make service call to get all products api and cache all products to hzl
}

// consider using some DI like google.wire to manage dependencies.
func main() {
	server.App()
}
