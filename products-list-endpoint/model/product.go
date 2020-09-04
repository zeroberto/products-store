package model

// Product is a model structure that represents a store product
type Product struct {
	ID           string    `json:"id"`
	PriceInCents int32     `json:"price_in_cents"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Discount     *Discount `json:"discount"`
}
