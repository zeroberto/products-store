package database

import (
	"github.com/zeroberto/products-store/discount-calculator/datastore"
	"github.com/zeroberto/products-store/discount-calculator/driver/dbdriver"
	"github.com/zeroberto/products-store/discount-calculator/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// ProductsCollection represents the name of the product collection in the database
	ProductsCollection string = "products"
)

// ProductDataStoreMongoDB is responsible for implementing the ProductDataStore interface,
// using a document database
type ProductDataStoreMongoDB struct {
	DBDriver dbdriver.NoSQLDBDriver
}

// FindByID is responsible for obtaining a product according to the given identifier
func (pds *ProductDataStoreMongoDB) FindByID(ID string) (*model.Product, error) {
	result, err := pds.DBDriver.GetDocByID(ID, ProductsCollection)
	if err != nil {
		return nil, &datastore.Error{Cause: err}
	}

	product := model.Product{
		ID:           result.(primitive.M)["_id"].(primitive.ObjectID).Hex(),
		PriceInCents: result.(primitive.M)["price_in_cents"].(int32),
	}

	return &product, nil
}
