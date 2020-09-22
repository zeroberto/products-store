package usecase

import (
	"time"

	"github.com/zeroberto/products-store/discount-calculator/model"
)

// DiscountUseCase represents the business rules for discounts
type DiscountUseCase interface {
	// CalculateDiscount is responsible for calculating the discounts
	CalculateDiscount(product *model.Product, user *model.User) *model.Discount
}

// DiscountRuleUseCase represents the business logic for discount rules
type DiscountRuleUseCase interface {
	// GetDiscountRules is responsible for providing all discount rules
	GetDiscountRules(time time.Time, user *model.User) []model.DiscountRule
}
