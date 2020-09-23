package grpcservice

import (
	"context"

	pbd "github.com/zeroberto/products-store/discount-calculator/pb/discountcalculator"
	"github.com/zeroberto/products-store/discount-calculator/usecase"
)

// DiscountCalculatorGrpcService is responsible for providing communication with the grpc service
type DiscountCalculatorGrpcService struct {
	DUC usecase.DiscountUseCase
}

// CalculateDiscount is responsible for calculating the discount for a product
func (dcgs *DiscountCalculatorGrpcService) CalculateDiscount(ctx context.Context, dreq *pbd.DiscountRequest) (*pbd.DiscountResponse, error) {
	return nil, nil
}
