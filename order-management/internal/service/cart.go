package service

import (
	"context"

	"github.com/akshay-misra-demo-apps/go-cart/order-management/api/pb"
)

type CartService struct {
	pb.UnimplementedCartServiceServer
}

func InitCartService() *CartService {
	return &CartService{}
}

func (s *CartService) CreateCart(context.Context, *pb.CreateCartRequest) (*pb.CreateCartResponse, error) {
	return nil, nil
}
func (s *CartService) ReadCart(context.Context, *pb.ReadCartRequest) (*pb.ReadCartResponse, error) {
	return nil, nil
}
func (s *CartService) UpdateCart(context.Context, *pb.UpdateCartRequest) (*pb.UpdateCartResponse, error) {
	return nil, nil
}
func (s *CartService) DeleteCart(context.Context, *pb.DeleteCartRequest) (*pb.DeleteCartResponse, error) {
	return nil, nil
}
func (s *CartService) AddProductToCart(context.Context, *pb.AddProductToCartRequest) (*pb.AddProductToCartResponse, error) {
	return nil, nil
}
func (s *CartService) RemoveProductFromCart(context.Context, *pb.RemoveProductFromCartRequest) (*pb.RemoveProductFromCartResponse, error) {
	return nil, nil
}
