package datastore

import (
	"errors"
	"reflect"
	"testing"

	"github.com/zeroberto/products-store/discount-calculator/datastore"
	"github.com/zeroberto/products-store/discount-calculator/datastore/database"
	"github.com/zeroberto/products-store/discount-calculator/driver/dbdriver"
	"github.com/zeroberto/products-store/discount-calculator/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestFindByID(t *testing.T) {
	expectedID := primitive.NewObjectID()
	var expectedPriceInCents int32 = 100
	expected := &model.Product{
		ID:           expectedID.Hex(),
		PriceInCents: expectedPriceInCents,
	}

	var dbDriver dbdriver.NoSQLDBDriver = &noSQLDBDriver{}
	getDocByID = func(hexID string, collection string) (interface{}, error) {
		return bson.M{
			"_id":            expectedID,
			"price_in_cents": expectedPriceInCents,
		}, nil
	}
	var pds datastore.ProductDataStore = &database.ProductDataStoreMongoDB{
		DBDriver: dbDriver,
	}

	got, err := pds.FindByID(expectedID.Hex())

	if err != nil {
		t.Errorf("FindByID() failed, expected %v, got %v", nil, err)
	}
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("FindByID() failed, expected %v, got %v", expected, got)
	}
}

func TestFindByID_WhenNoSQLDBDriverError_ThenFailure(t *testing.T) {
	expected := "test error"

	var dbDriver dbdriver.NoSQLDBDriver = &noSQLDBDriver{}
	getDocByID = func(hexID string, collection string) (interface{}, error) {
		return nil, errors.New(expected)
	}
	var pds datastore.ProductDataStore = &database.ProductDataStoreMongoDB{
		DBDriver: dbDriver,
	}

	got, err := pds.FindByID("test")

	if got != nil {
		t.Errorf("FindByID() failed, expected %v, got %v", nil, got)
	}
	if err == nil {
		t.Error("FindByID() failed, expected an error, got nil")
	}
	if expected != err.Error() {
		t.Errorf("FindByID() failed, expected %v, got %v", expected, got)
	}
}

var getDocByID func(hexID string, collection string) (interface{}, error)

type noSQLDBDriver struct{}

func (driver *noSQLDBDriver) GetDocByID(hexID string, collection string) (interface{}, error) {
	return getDocByID(hexID, collection)
}
