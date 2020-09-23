package dsfactory

import (
	"github.com/zeroberto/products-store/discount-calculator/container"
	"github.com/zeroberto/products-store/discount-calculator/container/factory/chronofactory"
	"github.com/zeroberto/products-store/discount-calculator/container/factory/clientfactory"
	"github.com/zeroberto/products-store/discount-calculator/container/factory/dbdriverfactory"
	"github.com/zeroberto/products-store/discount-calculator/datastore"
	"github.com/zeroberto/products-store/discount-calculator/datastore/database"
	"github.com/zeroberto/products-store/discount-calculator/datastore/network"
	"github.com/zeroberto/products-store/discount-calculator/pb/userinfo"
)

// DataStoreFactory is responsible for returning instances of datastores
type DataStoreFactory struct{}

// MakeProductDataStore is responsible for create a ProductDataStore instance
func (dsf *DataStoreFactory) MakeProductDataStore(c container.Container) datastore.ProductDataStore {
	dbdf := &dbdriverfactory.DBDriverFactory{}

	dbDriver := dbdf.MakeNoSQLDBDriver(c, &c.GetAppConfig().DSConfig.ProductConfig)

	return &database.ProductDataStoreMongoDB{DBDriver: dbDriver}
}

// MakeUserDataStore is responsible for create a UserDataStore instance
func (dsf *DataStoreFactory) MakeUserDataStore(c container.Container) datastore.UserDataStore {
	scf := &clientfactory.ServiceClientFactory{}

	config := &c.GetAppConfig().DSConfig.UserConfig

	return &network.UserDataStoreGrpc{
		TS: chronofactory.MakeTimeStamp(c),
		Client: func() (userinfo.UserInfoClient, error) {
			return scf.MakeUserInfoServiceClient(c, config.Host, config.Port)
		},
	}
}
