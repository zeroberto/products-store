// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package userinfo

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// UserInfoClient is the client API for UserInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserInfoClient interface {
	// GetUserInfo is responsible for obtaining the information for a particular user
	GetUserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
}

type userInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewUserInfoClient(cc grpc.ClientConnInterface) UserInfoClient {
	return &userInfoClient{cc}
}

var userInfoGetUserInfoStreamDesc = &grpc.StreamDesc{
	StreamName: "GetUserInfo",
}

func (c *userInfoClient) GetUserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, "/userinfo.UserInfo/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInfoService is the service API for UserInfo service.
// Fields should be assigned to their respective handler implementations only before
// RegisterUserInfoService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type UserInfoService struct {
	// GetUserInfo is responsible for obtaining the information for a particular user
	GetUserInfo func(context.Context, *UserInfoRequest) (*UserInfoResponse, error)
}

func (s *UserInfoService) getUserInfo(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/userinfo.UserInfo/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetUserInfo(ctx, req.(*UserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterUserInfoService registers a service implementation with a gRPC server.
func RegisterUserInfoService(s grpc.ServiceRegistrar, srv *UserInfoService) {
	srvCopy := *srv
	if srvCopy.GetUserInfo == nil {
		srvCopy.GetUserInfo = func(context.Context, *UserInfoRequest) (*UserInfoResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "userinfo.UserInfo",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "GetUserInfo",
				Handler:    srvCopy.getUserInfo,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "protobuf/proto/userinfo/user_info.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewUserInfoService creates a new UserInfoService containing the
// implemented methods of the UserInfo service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewUserInfoService(s interface{}) *UserInfoService {
	ns := &UserInfoService{}
	if h, ok := s.(interface {
		GetUserInfo(context.Context, *UserInfoRequest) (*UserInfoResponse, error)
	}); ok {
		ns.GetUserInfo = h.GetUserInfo
	}
	return ns
}

// UnstableUserInfoService is the service API for UserInfo service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableUserInfoService interface {
	// GetUserInfo is responsible for obtaining the information for a particular user
	GetUserInfo(context.Context, *UserInfoRequest) (*UserInfoResponse, error)
}
