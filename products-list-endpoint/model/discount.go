package model

// Discount is a model structure that represents a product discount
type Discount struct {
	Percentage   float32 `json:"pct"`
	ValueInCents int32   `json:"value_in_cents"`
}
