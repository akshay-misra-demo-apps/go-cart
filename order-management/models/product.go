package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cart struct {
	Entity
	Id         primitive.ObjectID `json:"id" bson:"_id"`
	UserId     primitive.ObjectID `json:"userId" bson:"userId"`
	OrderId    primitive.ObjectID `json:"orderId" bson:"orderId"`
	Items      []Product          `json:"items" bson:"items"`
	Quantity   int                `json:"quantity" bson:"quantity"`
	TotalPrice string             `json:"totalPrice" bson:"totalPrice"`
	TotalTax   string             `json:"totalTax" bson:"totalTax"`
	CouponCode string             `json:"couponCode" bson:"couponCode"`
}

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

type OrderItem struct {
	Entity
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	ProductId primitive.ObjectID `json:"productId" bson:"productId"`
	BasePrice string             `json:"basePrice" bson:"basePrice"`
	TaxAmount string             `json:"taxAmount" bson:"taxAmount"`
}

type Order struct {
	Entity
	Id         primitive.ObjectID `json:"id" bson:"_id"`
	UserId     primitive.ObjectID `json:"userId" bson:"userId"`
	OrderItem  []*OrderItem       `json:"orderItem" bson:"orderItem"`
	TotalPrice string             `json:"totalPrice" bson:"totalPrice"`
	TotalTax   string             `json:"totalTax" bson:"totalTax"`
}
