package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type IRepository interface {
	GetCollection() *mongo.Collection
}
