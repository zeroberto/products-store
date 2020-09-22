package model

// DiscountRule represents a simple strategic interface to create n types of
// discount calculations
type DiscountRule interface {
	// CalculateDiscount is responsible for making a discount calculation
	// following a specific rule
	CalculateDiscount() float32
}
