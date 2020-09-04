package datastorefactory

import (
	"github.com/pkg/errors"
	"github.com/zeroberto/products-store/products-list-endpoint/container"
	"github.com/zeroberto/products-store/products-list-endpoint/container/factory/clientfactory"
	"github.com/zeroberto/products-store/products-list-endpoint/datastore"
	"github.com/zeroberto/products-store/products-list-endpoint/datastore/network"
)

// DataStoreFactory is responsible for returning instances of datastores
type DataStoreFactory struct{}

// MakeProductDataStore is responsible for create a ProductDataStore instance
func (dsf *DataStoreFactory) MakeProductDataStore(c container.Container) (datastore.ProductDataStore, error) {
	scf := &clientfactory.ServiceClientFactory{}

	plsc, err := scf.MakeProductsListServiceClient(c)
	if err != nil {
		return nil, errors.Wrap(err, "Can not create a ProductDataStore instance")
	}

	return &network.ProductDataStoreGrpc{Client: plsc}, nil
}
