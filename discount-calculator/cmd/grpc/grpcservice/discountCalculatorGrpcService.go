package grpcservice

import (
	"context"

	"github.com/zeroberto/products-store/discount-calculator/datastore"
	"github.com/zeroberto/products-store/discount-calculator/model"
	pbd "github.com/zeroberto/products-store/discount-calculator/pb/discountcalculator"
	"github.com/zeroberto/products-store/discount-calculator/usecase"
)

// DiscountCalculatorGrpcService is responsible for providing communication with the grpc service
type DiscountCalculatorGrpcService struct {
	DUC usecase.DiscountUseCase
	PDS datastore.ProductDataStore
	UDS datastore.UserDataStore
}

// CalculateDiscount is responsible for calculating the discount for a product
func (dcgs *DiscountCalculatorGrpcService) CalculateDiscount(ctx context.Context, dreq *pbd.DiscountRequest) (*pbd.DiscountResponse, error) {
	p, err := dcgs.PDS.FindByID(dreq.GetProductId())
	if err != nil {
		return nil, err
	}

	u, err := dcgs.UDS.FindByID(dreq.GetUserId())
	if err != nil {
		return nil, err
	}

	d := dcgs.DUC.CalculateDiscount(p, u)

	return toDiscountResponse(p.ID, u.ID, d), nil
}

func toDiscountResponse(productID string, userID int64, d *model.Discount) *pbd.DiscountResponse {
	return &pbd.DiscountResponse{
		ProductId:    productID,
		UserId:       userID,
		Pct:          d.Percentage,
		ValueInCents: d.ValueInCents,
	}
}
