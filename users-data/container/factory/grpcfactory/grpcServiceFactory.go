package grpcfactory

import (
	"github.com/zeroberto/products-store/users-data/cmd/grpc/grpcservice"
	"github.com/zeroberto/products-store/users-data/container"
	"github.com/zeroberto/products-store/users-data/container/factory/dsfactory"
)

// GrpcServiceFactory is responsible for providing instances of grpc service structs
type GrpcServiceFactory struct{}

// MakeUserInfoGrpcService is responsible for providing an instance of UserInfoGrpcService
func (gsf *GrpcServiceFactory) MakeUserInfoGrpcService(c container.Container) *grpcservice.UserInfoGrpcService {
	dsf := &dsfactory.DataStoreFactory{}
	return &grpcservice.UserInfoGrpcService{UIDS: dsf.MakeUserInfoDataStore(c)}
}
