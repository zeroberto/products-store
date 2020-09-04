package datastore

import "github.com/zeroberto/products-store/products-list-endpoint/model"

// ProductDataStore is responsible for obtaining product data
type ProductDataStore interface {
	// FindAllWithDiscountByUserID is responsible for returning a list of discounted
	// products per user from a given repository
	FindAllWithDiscountByUserID(userID int64) ([]model.Product, error)
}
