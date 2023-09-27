package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	ConnectionString string
}

type DBConnection struct {
	Client *mongo.Client
}

func (d Mongo) Connect() (connection DBConnection, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(d.ConnectionString))
	if err != nil {
		return
	}

	//defer func() {
	//	if err = client.Disconnect(ctx); err != nil {
	//		panic(err)
	//	}
	//}()

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return
	}

	fmt.Println("Connected to MongoDB!")

	return DBConnection{Client: client}, err
}
