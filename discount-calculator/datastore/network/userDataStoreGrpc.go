package network

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/zeroberto/products-store/discount-calculator/chrono"
	"github.com/zeroberto/products-store/discount-calculator/model"
	"github.com/zeroberto/products-store/discount-calculator/pb/userinfo"
)

// UserDataStoreGrpc is responsible for implementing the UserDataStore interface
// using grpc
type UserDataStoreGrpc struct {
	TS     chrono.TimeStamp
	Client func() (userinfo.UserInfoClient, error)
}

// FindByID is responsible for obtaining an user according to the given identifier
func (uds *UserDataStoreGrpc) FindByID(ID int64) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	uireq := &userinfo.UserInfoRequest{}
	uireq.Id = 1

	client, err := uds.Client()
	if err != nil {
		return nil, errors.Wrap(err, "Could not connect to user service")
	}

	r, err := client.GetUserInfo(ctx, uireq)
	if err != nil {
		return nil, errors.Wrap(err, "Could not retrieve user")
	}

	return uds.toUser(r), nil
}

func (uds *UserDataStoreGrpc) toUser(uiresp *userinfo.UserInfoResponse) *model.User {
	return &model.User{
		ID:          uiresp.GetId(),
		DateOfBirth: uds.TS.GetTimeByNanoSeconds(uiresp.GetDateOfBirth()),
	}
}
