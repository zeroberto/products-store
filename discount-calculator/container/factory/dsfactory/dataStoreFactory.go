package dsfactory

import (
	"github.com/zeroberto/products-store/discount-calculator/container"
	"github.com/zeroberto/products-store/discount-calculator/container/factory/dbdriverfactory"
	"github.com/zeroberto/products-store/discount-calculator/datastore"
	"github.com/zeroberto/products-store/discount-calculator/datastore/database"
)

// DataStoreFactory is responsible for returning instances of datastores
type DataStoreFactory struct{}

// MakeProductDataStore is responsible for create a ProductDataStore instance
func (dsf *DataStoreFactory) MakeProductDataStore(c container.Container) datastore.ProductDataStore {
	dbdf := &dbdriverfactory.DBDriverFactory{}

	dbDriver := dbdf.MakeNoSQLDBDriver(c, &c.GetAppConfig().DSConfig.ProductConfig)

	return &database.ProductDataStoreMongoDB{DBDriver: dbDriver}
}
