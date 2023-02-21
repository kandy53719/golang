package services

import context "context"

type ProductService struct{}

func (ps ProductService) GetProductStock(ctx context.Context, request *ProductRequest) (*ProductResponse, error) {
	return &ProductResponse{ProductStock: 20}, nil
}

func (ps ProductService) mustEmbedUnimplementedProductServiceServer() {}
