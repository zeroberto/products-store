package dsfactory

import (
	"github.com/zeroberto/products-store/users-data/container"
	"github.com/zeroberto/products-store/users-data/container/factory/driverfactory"
	"github.com/zeroberto/products-store/users-data/datastore"
	"github.com/zeroberto/products-store/users-data/datastore/sqldatastore"
)

// DataStoreFactory is responsible for providing instances of data store structs
type DataStoreFactory struct{}

// MakeUserInfoDataStore is responsible for providing an instance of UserInfoDataStore
func (dsf *DataStoreFactory) MakeUserInfoDataStore(c container.Container) (datastore.UserInfoDataStore, error) {
	df := driverfactory.DBDriverFactory{}

	driver, err := df.MakeSQLDBDriver(c, &c.GetAppConfig().DSConfig.SQLConfig.UserInfoConfig)
	if err != nil {
		return nil, err
	}

	eds := &sqldatastore.UserInfoDataStoreSQL{
		SQLDriver: driver,
	}

	return eds, nil
}
