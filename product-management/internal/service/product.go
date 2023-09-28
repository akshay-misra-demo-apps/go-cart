package service

import (
	"context"
	"log"

	"github.com/akshay-misra-demo-apps/go-cart/product-management/api/pb"
)

type ProductService struct {
	pb.UnimplementedProductServiceServer
}

func InitProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) CreateProduct(ctx context.Context, in *pb.Product) (*pb.CreateProductResponse, error) {
	// TODO: Implement logic for CreateProduct
	return &pb.CreateProductResponse{Product: in}, nil
}

func (s *ProductService) ReadProduct(ctx context.Context, in *pb.ProductID) (*pb.ReadProductResponse, error) {
	// TODO: Implement logic for ReadProduct
	return &pb.ReadProductResponse{Product: &pb.Product{}}, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, in *pb.Product) (*pb.UpdateProductResponse, error) {
	// TODO: Implement logic for UpdateProduct
	return &pb.UpdateProductResponse{Product: in}, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, in *pb.ProductID) (*pb.DeleteProductResponse, error) {
	// TODO: Implement logic for DeleteProduct
	return &pb.DeleteProductResponse{Success: true}, nil
}

func (s *ProductService) SearchProducts(ctx context.Context, in *pb.SearchRequest) (*pb.SearchProductsResponse, error) {
	// TODO: Implement logic for SearchProducts
	return &pb.SearchProductsResponse{Products: &pb.ProductList{}}, nil
}

func (s *ProductService) ReadAllProducts(ctx context.Context, in *pb.ReadAllProductsRequest) (*pb.ReadAllProductsResponse, error) {
	// TODO: Implement logic for ReadAllProducts
	list := &pb.ProductList{
		Products: []*pb.Product{
			{
				Id:          "",
				Name:        "",
				Category:    "Electronics",
				BasePrice:   "500",
				Description: "iPhone 12 pro max, white, 6gb, 64gb",
				Tags:        []string{"electronics", "mobile"},
			},
		},
	}
	log.Printf("response from grpc api ReadAllProducts: %v", list)
	return &pb.ReadAllProductsResponse{Products: list}, nil
}
