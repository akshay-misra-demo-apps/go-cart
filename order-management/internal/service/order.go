package service

import (
	"github.com/akshay-misra-demo-apps/go-cart/order-management/api/pb"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
}

func InitOrderService() *OrderService {
	return &OrderService{}
}
