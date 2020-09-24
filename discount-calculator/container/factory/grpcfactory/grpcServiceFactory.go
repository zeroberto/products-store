package grpcfactory

import (
	"github.com/zeroberto/products-store/discount-calculator/cmd/grpc/grpcservice"
	"github.com/zeroberto/products-store/discount-calculator/container"
	"github.com/zeroberto/products-store/discount-calculator/container/factory/dsfactory"
	"github.com/zeroberto/products-store/discount-calculator/container/factory/ucfactory"
)

// GrpcServiceFactory is responsible for providing instances of grpc service structs
type GrpcServiceFactory struct{}

// MakeDiscountCalculatorGrpcService is responsible for providing an instance of DiscountCalculatorGrpcService
func (gsf *GrpcServiceFactory) MakeDiscountCalculatorGrpcService(c container.Container) *grpcservice.DiscountCalculatorGrpcService {
	dsf := &dsfactory.DataStoreFactory{}
	ucf := &ucfactory.UseCaseFactory{}

	return &grpcservice.DiscountCalculatorGrpcService{
		DUC: ucf.MakeDiscountUseCase(c),
		PDS: dsf.MakeProductDataStore(c),
		UDS: dsf.MakeUserDataStore(c),
	}
}
