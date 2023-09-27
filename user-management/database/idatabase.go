package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type IDatabase interface {
	Connect() (*mongo.Client, error)
}
