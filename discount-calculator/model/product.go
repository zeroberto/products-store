package model

// Product is a model structure that represents a store product
type Product struct {
	ID           string
	PriceInCents int32
	Title        string
	Description  string
	Discount     *Discount
}
