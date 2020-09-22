package discountuc

import (
	"math"

	"github.com/zeroberto/products-store/discount-calculator/chrono"
	"github.com/zeroberto/products-store/discount-calculator/model"
	"github.com/zeroberto/products-store/discount-calculator/usecase"
)

// RuleCalculatorUseCase is responsible for implementing the DiscountUseCase interface
type RuleCalculatorUseCase struct {
	TimeStamp chrono.TimeStamp
	DRUC      usecase.DiscountRuleUseCase
}

// CalculateDiscount is responsible for calculating the discounts
func (rcuc *RuleCalculatorUseCase) CalculateDiscount(product *model.Product, user *model.User) *model.Discount {
	var discountPct float32

	for _, rule := range rcuc.DRUC.GetDiscountRules(rcuc.TimeStamp.GetCurrentTime(), user) {
		discountPct += rule.CalculateDiscount()
	}

	if discountPct > MaxPct {
		discountPct = MaxPct
	}

	return &model.Discount{
		Percentage:   discountPct,
		ValueInCents: int32(math.Round(float64(product.PriceInCents) * float64(discountPct))),
	}
}
