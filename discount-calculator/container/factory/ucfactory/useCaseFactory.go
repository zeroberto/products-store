package ucfactory

import (
	"github.com/zeroberto/products-store/discount-calculator/container"
	"github.com/zeroberto/products-store/discount-calculator/container/factory/chronofactory"
	"github.com/zeroberto/products-store/discount-calculator/usecase"
	"github.com/zeroberto/products-store/discount-calculator/usecase/discountuc"
)

// UseCaseFactory is responsible for returning instances of use cases
type UseCaseFactory struct{}

// MakeDiscountUseCase is responsible for create a DiscountUseCase instance
func (ucf *UseCaseFactory) MakeDiscountUseCase(c container.Container) usecase.DiscountUseCase {
	return &discountuc.RuleCalculatorUseCase{
		TimeStamp: chronofactory.MakeTimeStamp(c),
		DRUC:      ucf.MakeDiscountRuleUseCase(c),
	}
}

// MakeDiscountRuleUseCase is responsible for create a DiscountRuleUseCase instance
func (ucf *UseCaseFactory) MakeDiscountRuleUseCase(c container.Container) usecase.DiscountRuleUseCase {
	return &discountuc.RuleProviderUseCase{
		TimeStamp: chronofactory.MakeTimeStamp(c),
	}
}
