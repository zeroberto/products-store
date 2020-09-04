package apifactory

import (
	"github.com/pkg/errors"
	"github.com/zeroberto/products-store/products-list-endpoint/api"
	"github.com/zeroberto/products-store/products-list-endpoint/api/rest"
	"github.com/zeroberto/products-store/products-list-endpoint/container"
	"github.com/zeroberto/products-store/products-list-endpoint/container/factory/chronofactory"
	"github.com/zeroberto/products-store/products-list-endpoint/container/factory/datastorefactory"
)

// APIFactory is responsible for providing instances of api structs
type APIFactory struct{}

// MakeProductAPI is responsible for providing an instance of ProductAPI
// with ProductAPIRest implementation
func (af *APIFactory) MakeProductAPI(c container.Container) (api.ProductAPI, error) {
	ts := chronofactory.MakeTimeStamp(c)

	dsf := &datastorefactory.DataStoreFactory{}
	pds, err := dsf.MakeProductDataStore(c)
	if err != nil {
		return nil, errors.Wrap(err, "Can not create a Product API")
	}

	return &rest.ProductAPIRest{PDS: pds, TS: ts}, nil
}
