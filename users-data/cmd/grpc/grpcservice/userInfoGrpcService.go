package grpcservice

import (
	"context"
	"errors"

	"github.com/zeroberto/products-store/users-data/datastore"
	"github.com/zeroberto/products-store/users-data/model"
	pb "github.com/zeroberto/products-store/users-data/pb/userinfo"
)

// UserInfoGrpcService is responsible for providing communication with the grpc service
type UserInfoGrpcService struct {
	UIDS datastore.UserInfoDataStore
}

// GetUserInfo is responsible for obtaining information from a user
func (uigs *UserInfoGrpcService) GetUserInfo(ctx context.Context, uireq *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	ui, err := uigs.UIDS.FindByID(uireq.GetId())
	if err != nil {
		return nil, &datastore.Error{Cause: err}
	}
	if ui == nil {
		return nil, &datastore.Error{Cause: errors.New("User not found for this id")}
	}
	return toUserInfoResponse(ui), nil
}

func toUserInfoResponse(ui *model.UserInfo) *pb.UserInfoResponse {
	return &pb.UserInfoResponse{
		Id:          ui.ID,
		FirstName:   ui.FirstName,
		LastName:    ui.LastName,
		DateOfBirth: ui.DateOfBirth.UnixNano(),
		CreatedAt:   ui.CreatedAt.UnixNano(),
	}
}

// Error is responsible for encapsulating errors generated by operations in grpc service
type Error struct {
	Cause error
}

func (err *Error) Error() string {
	return err.Cause.Error()
}
