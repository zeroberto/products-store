package grpcservice

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/zeroberto/products-store/users-data/cmd/grpc/grpcservice"
	"github.com/zeroberto/products-store/users-data/datastore"
	"github.com/zeroberto/products-store/users-data/model"
	pb "github.com/zeroberto/products-store/users-data/pb/userinfo"
)

func TestGetUserInfo(t *testing.T) {
	currentTime := time.Now()
	ui := &model.UserInfo{
		ID:          1,
		FirstName:   "test",
		LastName:    "test",
		DateOfBirth: currentTime,
		CreatedAt:   currentTime,
	}
	expected := &pb.UserInfoResponse{
		Id:          ui.ID,
		FirstName:   ui.FirstName,
		LastName:    ui.LastName,
		DateOfBirth: ui.DateOfBirth.UnixNano(),
		CreatedAt:   ui.CreatedAt.UnixNano(),
	}
	req := &pb.UserInfoRequest{
		Id: ui.ID,
	}

	var uids datastore.UserInfoDataStore = &userInfoDataStoreMock{}
	findByID = func(ID int64) (*model.UserInfo, error) {
		return ui, nil
	}
	uigs := &grpcservice.UserInfoGrpcService{UIDS: uids}

	got, err := uigs.GetUserInfo(nil, req)

	if err != nil {
		t.Errorf("GetUserInfo() failed, expected %v, got %v", nil, err)
	}
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("GetUserInfo() failed, expected %v, got %v", expected, got)
	}
}

func TestGetUserInfo_WhenUserInfoDataStoreError_ThenFailure(t *testing.T) {
	req := &pb.UserInfoRequest{
		Id: 1,
	}

	var uids datastore.UserInfoDataStore = &userInfoDataStoreMock{}
	findByID = func(ID int64) (*model.UserInfo, error) {
		return nil, errors.New("Could not get user information")
	}
	uigs := &grpcservice.UserInfoGrpcService{UIDS: uids}

	_, err := uigs.GetUserInfo(nil, req)

	if err == nil {
		t.Error("ReadConfig() failed, expected an error, got nil")
	}
}

var findByID func(ID int64) (*model.UserInfo, error)

type userInfoDataStoreMock struct{}

func (eruc *userInfoDataStoreMock) FindByID(ID int64) (*model.UserInfo, error) {
	return findByID(ID)
}
