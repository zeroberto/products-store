package api

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/zeroberto/products-store/products-list-endpoint/api"
	"github.com/zeroberto/products-store/products-list-endpoint/api/rest"
	"github.com/zeroberto/products-store/products-list-endpoint/chrono/provider"
	"github.com/zeroberto/products-store/products-list-endpoint/datastore"
	"github.com/zeroberto/products-store/products-list-endpoint/model"
)

func TestProductAPIGet(t *testing.T) {
	products := []model.Product{
		model.Product{ID: "1"},
		model.Product{ID: "2"},
	}
	expected := api.Response{
		Code: 200,
		Body: products,
	}

	var pds datastore.ProductDataStore = &productDataStoreMock{}
	findAllWithDiscountByUserID = func(userID int64) ([]model.Product, error) {
		return products, nil
	}

	var papi api.ProductAPI = &rest.ProductAPIRest{
		PDS: pds,
		TS:  &provider.TimeStampImpl{},
	}
	got := papi.Get(1)

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Get() failed, expected %v, got %v", expected, got)
	}
}

func TestProductAPIGet_WhenProductDataStoreError_ThenFailure(t *testing.T) {
	expectedCode := 500
	expectedTime := time.Now()
	expected := api.Response{
		Code: expectedCode,
		Body: api.ResponseBody{
			Code:    expectedCode,
			Message: "Could not get product list",
			Time:    expectedTime,
		},
	}

	var pds datastore.ProductDataStore = &productDataStoreMock{}
	findAllWithDiscountByUserID = func(userID int64) ([]model.Product, error) {
		return nil, errors.New("Could not get product list")
	}

	var papi api.ProductAPI = &rest.ProductAPIRest{
		PDS: pds,
		TS:  &timeStampMock{Time: expectedTime},
	}
	got := papi.Get(1)

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Get() failed, expected %v, got %v", expected, got)
	}
}

var findAllWithDiscountByUserID func(userID int64) ([]model.Product, error)

type productDataStoreMock struct{}

type timeStampMock struct {
	Time time.Time
}

func (eruc *productDataStoreMock) FindAllWithDiscountByUserID(userID int64) ([]model.Product, error) {
	return findAllWithDiscountByUserID(userID)
}

func (tp *timeStampMock) GetCurrentTime() time.Time {
	return tp.Time
}
