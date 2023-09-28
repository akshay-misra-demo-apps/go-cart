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
	Category    string             `json:"category" bson:"category"`
	ImageURL    string             `json:"imageURL,omitempty" bson:"imageURL"`
	Tags        []string           `json:"tags,omitempty" bson:"tags,omitempty"`
	Version     int                `json:"version" bson:"version"`
}

func (p Product) Clone() *Product {
	return &Product{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		BasePrice:   p.BasePrice,
		Category:    p.Category,
		ImageURL:    p.ImageURL,
		Version:     p.Version,
	}
}
