package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Entity
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description,omitempty" bson:"description"`
	BasePrice   string             `json:"basePrice" bson:"basePrice"`
	ImageURL    string             `json:"imageURL,omitempty" bson:"imageURL"`
	Version     int                `json:"version" bson:"version"`
}

func (p Product) Clone() *Product {
	return &Product{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		BasePrice:   p.BasePrice,
		ImageURL:    p.ImageURL,
		Version:     p.Version,
	}
}
