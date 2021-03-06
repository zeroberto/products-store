// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package productslist

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ProductsListClient is the client API for ProductsList service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductsListClient interface {
	// ListProducts is responsible for returning a list of all products with discounts
	// applied for a given user, if an identifier is passed
	ListProducts(ctx context.Context, in *ProductsListRequest, opts ...grpc.CallOption) (ProductsList_ListProductsClient, error)
}

type productsListClient struct {
	cc grpc.ClientConnInterface
}

func NewProductsListClient(cc grpc.ClientConnInterface) ProductsListClient {
	return &productsListClient{cc}
}

var productsListListProductsStreamDesc = &grpc.StreamDesc{
	StreamName:    "ListProducts",
	ServerStreams: true,
}

func (c *productsListClient) ListProducts(ctx context.Context, in *ProductsListRequest, opts ...grpc.CallOption) (ProductsList_ListProductsClient, error) {
	stream, err := c.cc.NewStream(ctx, productsListListProductsStreamDesc, "/productslist.ProductsList/ListProducts", opts...)
	if err != nil {
		return nil, err
	}
	x := &productsListListProductsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProductsList_ListProductsClient interface {
	Recv() (*ProductsListResponse, error)
	grpc.ClientStream
}

type productsListListProductsClient struct {
	grpc.ClientStream
}

func (x *productsListListProductsClient) Recv() (*ProductsListResponse, error) {
	m := new(ProductsListResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProductsListService is the service API for ProductsList service.
// Fields should be assigned to their respective handler implementations only before
// RegisterProductsListService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type ProductsListService struct {
	// ListProducts is responsible for returning a list of all products with discounts
	// applied for a given user, if an identifier is passed
	ListProducts func(*ProductsListRequest, ProductsList_ListProductsServer) error
}

func (s *ProductsListService) listProducts(_ interface{}, stream grpc.ServerStream) error {
	m := new(ProductsListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return s.ListProducts(m, &productsListListProductsServer{stream})
}

type ProductsList_ListProductsServer interface {
	Send(*ProductsListResponse) error
	grpc.ServerStream
}

type productsListListProductsServer struct {
	grpc.ServerStream
}

func (x *productsListListProductsServer) Send(m *ProductsListResponse) error {
	return x.ServerStream.SendMsg(m)
}

// RegisterProductsListService registers a service implementation with a gRPC server.
func RegisterProductsListService(s grpc.ServiceRegistrar, srv *ProductsListService) {
	srvCopy := *srv
	if srvCopy.ListProducts == nil {
		srvCopy.ListProducts = func(*ProductsListRequest, ProductsList_ListProductsServer) error {
			return status.Errorf(codes.Unimplemented, "method ListProducts not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "productslist.ProductsList",
		Methods:     []grpc.MethodDesc{},
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "ListProducts",
				Handler:       srvCopy.listProducts,
				ServerStreams: true,
			},
		},
		Metadata: "protobuf/proto/productslist/products_list.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewProductsListService creates a new ProductsListService containing the
// implemented methods of the ProductsList service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewProductsListService(s interface{}) *ProductsListService {
	ns := &ProductsListService{}
	if h, ok := s.(interface {
		ListProducts(*ProductsListRequest, ProductsList_ListProductsServer) error
	}); ok {
		ns.ListProducts = h.ListProducts
	}
	return ns
}

// UnstableProductsListService is the service API for ProductsList service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableProductsListService interface {
	// ListProducts is responsible for returning a list of all products with discounts
	// applied for a given user, if an identifier is passed
	ListProducts(*ProductsListRequest, ProductsList_ListProductsServer) error
}
