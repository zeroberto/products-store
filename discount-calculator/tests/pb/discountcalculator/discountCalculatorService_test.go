package discountcalculator

import (
	"errors"
	"reflect"
	"testing"

	"github.com/zeroberto/products-store/discount-calculator/cmd/grpc/grpcservice"
	"github.com/zeroberto/products-store/discount-calculator/datastore"
	"github.com/zeroberto/products-store/discount-calculator/model"
	pbd "github.com/zeroberto/products-store/discount-calculator/pb/discountcalculator"
	"github.com/zeroberto/products-store/discount-calculator/usecase"
)

func TestCalculateDiscount(t *testing.T) {
	expected := &pbd.DiscountResponse{
		ProductId:    "test",
		UserId:       1,
		Pct:          0.05,
		ValueInCents: 6,
	}
	request := &pbd.DiscountRequest{
		ProductId: "test",
		UserId:    1,
	}

	var duc usecase.DiscountUseCase = &discountUseCaseMock{}
	calculateDiscount = func(product *model.Product, user *model.User) *model.Discount {
		return &model.Discount{
			Percentage:   0.05,
			ValueInCents: 6,
		}
	}

	var pds datastore.ProductDataStore = &productDataStoreMock{}
	findProductByID = func(ID string) (*model.Product, error) {
		return &model.Product{ID: "test"}, nil
	}

	var uds datastore.UserDataStore = &userDataStoreMock{}
	findUserByID = func(ID int64) (*model.User, error) {
		return &model.User{ID: 1}, nil
	}

	dcs := &grpcservice.DiscountCalculatorGrpcService{
		DUC: duc,
		PDS: pds,
		UDS: uds,
	}

	got, err := dcs.CalculateDiscount(nil, request)

	if err != nil {
		t.Errorf("CalculateDiscount() failed, expected %v, got %v", nil, err)
	}
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("CalculateDiscount() failed, expected %v, got %v", expected, got)
	}
}

func TestCalculateDiscount_WhenProductError_ThenFailure(t *testing.T) {
	expected := "error message"

	request := &pbd.DiscountRequest{
		ProductId: "test",
		UserId:    1,
	}

	var duc usecase.DiscountUseCase = &discountUseCaseMock{}
	calculateDiscount = func(product *model.Product, user *model.User) *model.Discount {
		return nil
	}

	var uds datastore.UserDataStore = &userDataStoreMock{}
	findUserByID = func(ID int64) (*model.User, error) {
		return nil, nil
	}

	var pds datastore.ProductDataStore = &productDataStoreMock{}
	findProductByID = func(ID string) (*model.Product, error) {
		return nil, errors.New(expected)
	}

	dcs := &grpcservice.DiscountCalculatorGrpcService{
		DUC: duc,
		PDS: pds,
		UDS: uds,
	}

	d, got := dcs.CalculateDiscount(nil, request)

	if d != nil {
		t.Errorf("CalculateDiscount() failed, expected %v, got %v", nil, d)
	}
	if got == nil {
		t.Error("CalculateDiscount() failed, expected an error, got nil")
	}
	if expected != got.Error() {
		t.Errorf("CalculateDiscount() failed, expected %v, got %v", expected, got)
	}
}

func TestCalculateDiscount_WhenUserError_ThenFailure(t *testing.T) {
	expected := "error message"

	request := &pbd.DiscountRequest{
		ProductId: "test",
		UserId:    1,
	}

	var duc usecase.DiscountUseCase = &discountUseCaseMock{}
	calculateDiscount = func(product *model.Product, user *model.User) *model.Discount {
		return nil
	}

	var pds datastore.ProductDataStore = &productDataStoreMock{}
	findProductByID = func(ID string) (*model.Product, error) {
		return nil, nil
	}

	var uds datastore.UserDataStore = &userDataStoreMock{}
	findUserByID = func(ID int64) (*model.User, error) {
		return nil, errors.New(expected)
	}

	dcs := &grpcservice.DiscountCalculatorGrpcService{
		DUC: duc,
		PDS: pds,
		UDS: uds,
	}

	d, got := dcs.CalculateDiscount(nil, request)

	if d != nil {
		t.Errorf("CalculateDiscount() failed, expected %v, got %v", nil, d)
	}
	if got == nil {
		t.Error("CalculateDiscount() failed, expected an error, got nil")
	}
	if expected != got.Error() {
		t.Errorf("CalculateDiscount() failed, expected %v, got %v", expected, got)
	}
}

var calculateDiscount func(product *model.Product, user *model.User) *model.Discount

var findProductByID func(ID string) (*model.Product, error)

var findUserByID func(ID int64) (*model.User, error)

type discountUseCaseMock struct{}

type productDataStoreMock struct{}

type userDataStoreMock struct{}

func (duc *discountUseCaseMock) CalculateDiscount(product *model.Product, user *model.User) *model.Discount {
	return calculateDiscount(product, user)
}

func (duc *productDataStoreMock) FindByID(ID string) (*model.Product, error) {
	return findProductByID(ID)
}

func (duc *userDataStoreMock) FindByID(ID int64) (*model.User, error) {
	return findUserByID(ID)
}
