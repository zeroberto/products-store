package grpcfactory

import (
	"github.com/pkg/errors"
	"github.com/zeroberto/products-store/users-data/cmd/grpc/grpcservice"
	"github.com/zeroberto/products-store/users-data/container"
	"github.com/zeroberto/products-store/users-data/container/factory/dsfactory"
)

// GrpcServiceFactory is responsible for providing instances of grpc service structs
type GrpcServiceFactory struct{}

// MakeUserInfoGrpcService is responsible for providing an instance of UserInfoGrpcService
func (gsf *GrpcServiceFactory) MakeUserInfoGrpcService(c container.Container, dsf *dsfactory.DataStoreFactory) (*grpcservice.UserInfoGrpcService, error) {
	uids, err := dsf.MakeUserInfoDataStore(c)
	if err != nil {
		return nil, errors.Wrap(err, "Can not create a UserInfoDataStore instance")
	}

	return &grpcservice.UserInfoGrpcService{UIDS: uids}, nil
}
