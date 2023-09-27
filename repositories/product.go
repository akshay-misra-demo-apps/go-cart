package repositories

import (
	"github.com/akshay-misra-demo-apps/go-cart/user-management/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	connection database.DBConnection
}

// GetCollection Call this everytime we want to interact with the mongo collection, like insert, read etc.
func (r ProductRepository) GetCollection() *mongo.Collection {
	// Get Products collection
	return r.connection.Client.
		Database("go-cart").
		Collection("Products")
}

func initProductRepo(connection database.DBConnection) {
	productRepository := ProductRepository{
		connection: connection,
	}
	repositories["product"] = productRepository
}
