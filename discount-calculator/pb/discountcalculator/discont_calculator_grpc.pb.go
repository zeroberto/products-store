// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package discountcalculator

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// DiscountCalculatorClient is the client API for DiscountCalculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiscountCalculatorClient interface {
	// CalculateDiscount is responsible for calculating the discount for a product based
	// on an informed user
	CalculateDiscount(ctx context.Context, in *DiscountRequest, opts ...grpc.CallOption) (*DiscountResponse, error)
}

type discountCalculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscountCalculatorClient(cc grpc.ClientConnInterface) DiscountCalculatorClient {
	return &discountCalculatorClient{cc}
}

var discountCalculatorCalculateDiscountStreamDesc = &grpc.StreamDesc{
	StreamName: "CalculateDiscount",
}

func (c *discountCalculatorClient) CalculateDiscount(ctx context.Context, in *DiscountRequest, opts ...grpc.CallOption) (*DiscountResponse, error) {
	out := new(DiscountResponse)
	err := c.cc.Invoke(ctx, "/discountcalculator.DiscountCalculator/CalculateDiscount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscountCalculatorService is the service API for DiscountCalculator service.
// Fields should be assigned to their respective handler implementations only before
// RegisterDiscountCalculatorService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type DiscountCalculatorService struct {
	// CalculateDiscount is responsible for calculating the discount for a product based
	// on an informed user
	CalculateDiscount func(context.Context, *DiscountRequest) (*DiscountResponse, error)
}

func (s *DiscountCalculatorService) calculateDiscount(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.CalculateDiscount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/discountcalculator.DiscountCalculator/CalculateDiscount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.CalculateDiscount(ctx, req.(*DiscountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterDiscountCalculatorService registers a service implementation with a gRPC server.
func RegisterDiscountCalculatorService(s grpc.ServiceRegistrar, srv *DiscountCalculatorService) {
	srvCopy := *srv
	if srvCopy.CalculateDiscount == nil {
		srvCopy.CalculateDiscount = func(context.Context, *DiscountRequest) (*DiscountResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method CalculateDiscount not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "discountcalculator.DiscountCalculator",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "CalculateDiscount",
				Handler:    srvCopy.calculateDiscount,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "protobuf/proto/discountcalculator/discont_calculator.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewDiscountCalculatorService creates a new DiscountCalculatorService containing the
// implemented methods of the DiscountCalculator service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewDiscountCalculatorService(s interface{}) *DiscountCalculatorService {
	ns := &DiscountCalculatorService{}
	if h, ok := s.(interface {
		CalculateDiscount(context.Context, *DiscountRequest) (*DiscountResponse, error)
	}); ok {
		ns.CalculateDiscount = h.CalculateDiscount
	}
	return ns
}

// UnstableDiscountCalculatorService is the service API for DiscountCalculator service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableDiscountCalculatorService interface {
	// CalculateDiscount is responsible for calculating the discount for a product based
	// on an informed user
	CalculateDiscount(context.Context, *DiscountRequest) (*DiscountResponse, error)
}
