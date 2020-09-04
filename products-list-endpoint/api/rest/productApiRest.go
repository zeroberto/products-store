package rest

import (
	"net/http"
	"time"

	"github.com/zeroberto/products-store/products-list-endpoint/api"
	"github.com/zeroberto/products-store/products-list-endpoint/chrono"
	"github.com/zeroberto/products-store/products-list-endpoint/datastore"
)

// ProductAPIRest is responsible for implementing the ProductAPI interface using HTTP REST abstraction
type ProductAPIRest struct {
	PDS datastore.ProductDataStore
	TS  chrono.TimeStamp
}

// Get provides all products, with the option of listing by user,
// if an id is passed via header
func (par *ProductAPIRest) Get(userID int64) api.Response {
	products, err := par.PDS.FindAllWithDiscountByUserID(userID)
	if err != nil {
		return report(err, par.TS.GetCurrentTime())
	}
	return api.Response{
		Code: http.StatusOK,
		Body: products,
	}
}

func report(err error, time time.Time) api.Response {
	code := http.StatusInternalServerError
	return api.Response{
		Code: code,
		Body: api.ResponseBody{
			Time:    time,
			Code:    code,
			Message: err.Error(),
		},
	}
}
