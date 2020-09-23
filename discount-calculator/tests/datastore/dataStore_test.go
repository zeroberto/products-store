package datastore

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/zeroberto/products-store/discount-calculator/chrono/provider"
	"github.com/zeroberto/products-store/discount-calculator/datastore"
	"github.com/zeroberto/products-store/discount-calculator/datastore/database"
	"github.com/zeroberto/products-store/discount-calculator/datastore/network"
	"github.com/zeroberto/products-store/discount-calculator/driver/dbdriver"
	"github.com/zeroberto/products-store/discount-calculator/model"
	"github.com/zeroberto/products-store/discount-calculator/pb/userinfo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
)

func TestProductDataStoreFindByID(t *testing.T) {
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

func TestProductDataStoreFindByID_WhenNoSQLDBDriverError_ThenFailure(t *testing.T) {
	expected := "test error"

	var dbDriver dbdriver.NoSQLDBDriver = &noSQLDBDriver{}
	getDocByID = func(hexID string, collection string) (interface{}, error) {
		return nil, errors.New(expected)
	}
	var pds datastore.ProductDataStore = &database.ProductDataStoreMongoDB{
		DBDriver: dbDriver,
	}

	p, got := pds.FindByID("test")

	if p != nil {
		t.Errorf("FindByID() failed, expected %v, got %v", nil, p)
	}
	if got == nil {
		t.Error("FindByID() failed, expected an error, got nil")
	}
	if expected != got.Error() {
		t.Errorf("FindByID() failed, expected %v, got %v", expected, got)
	}
}

func TestUserDataStoreFindByID(t *testing.T) {
	time.Local = time.UTC

	var expectedID int64 = 1
	expectedDateOfBirth := time.Unix(0, time.Date(2020, time.May, 1, 0, 0, 0, 0, time.UTC).UnixNano())
	expected := &model.User{
		ID:          expectedID,
		DateOfBirth: expectedDateOfBirth,
	}

	var client userinfo.UserInfoClient = &userInfoClientMock{}
	getUserInfo = func(ctx context.Context, in *userinfo.UserInfoRequest, opts ...grpc.CallOption) (*userinfo.UserInfoResponse, error) {
		return &userinfo.UserInfoResponse{
			Id:          1,
			DateOfBirth: time.Date(2020, time.May, 1, 0, 0, 0, 0, time.UTC).UnixNano(),
		}, nil
	}
	var uds datastore.UserDataStore = &network.UserDataStoreGrpc{
		TS: &provider.TimeStampImpl{},
		Client: func() (userinfo.UserInfoClient, error) {
			return client, nil
		},
	}

	got, err := uds.FindByID(expectedID)

	if err != nil {
		t.Errorf("FindByID() failed, expected %v, got %v", nil, err)
	}
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("FindByID() failed, expected %v, got %v", expected, got)
	}
}

func TestUserDataStoreFindByID_WhenClientError_ThenFailure(t *testing.T) {
	errorMessage := "test error"
	expected := "Could not retrieve user: " + errorMessage

	var client userinfo.UserInfoClient = &userInfoClientMock{}
	getUserInfo = func(ctx context.Context, in *userinfo.UserInfoRequest, opts ...grpc.CallOption) (*userinfo.UserInfoResponse, error) {
		return nil, errors.New(errorMessage)
	}
	var uds datastore.UserDataStore = &network.UserDataStoreGrpc{
		TS: &provider.TimeStampImpl{},
		Client: func() (userinfo.UserInfoClient, error) {
			return client, nil
		},
	}

	u, got := uds.FindByID(1)

	if u != nil {
		t.Errorf("FindByID() failed, expected %v, got %v", nil, u)
	}
	if got == nil {
		t.Error("FindByID() failed, expected an error, got nil")
	}
	if expected != got.Error() {
		t.Errorf("FindByID() failed, expected %v, got %v", expected, got)
	}
}

func TestUserDataStoreFindByID_WhenClientConnError_ThenFailure(t *testing.T) {
	errorMessage := "test error"
	expected := "Could not connect to user service: " + errorMessage

	var uds datastore.UserDataStore = &network.UserDataStoreGrpc{
		TS: &provider.TimeStampImpl{},
		Client: func() (userinfo.UserInfoClient, error) {
			return nil, errors.New(errorMessage)
		},
	}

	u, got := uds.FindByID(1)

	if u != nil {
		t.Errorf("FindByID() failed, expected %v, got %v", nil, u)
	}
	if got == nil {
		t.Error("FindByID() failed, expected an error, got nil")
	}
	if expected != got.Error() {
		t.Errorf("FindByID() failed, expected %v, got %v", expected, got)
	}
}

var getDocByID func(hexID string, collection string) (interface{}, error)

var getUserInfo func(
	ctx context.Context,
	in *userinfo.UserInfoRequest,
	opts ...grpc.CallOption,
) (*userinfo.UserInfoResponse, error)

type noSQLDBDriver struct{}

type userInfoClientMock struct{}

func (driver *noSQLDBDriver) GetDocByID(hexID string, collection string) (interface{}, error) {
	return getDocByID(hexID, collection)
}

func (client *userInfoClientMock) GetUserInfo(
	ctx context.Context,
	in *userinfo.UserInfoRequest,
	opts ...grpc.CallOption,
) (*userinfo.UserInfoResponse, error) {
	return getUserInfo(ctx, in, opts...)
}
