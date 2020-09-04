package clientfactory

import (
	"github.com/pkg/errors"
	"github.com/zeroberto/products-store/products-list-endpoint/container"
	"github.com/zeroberto/products-store/products-list-endpoint/pb/productslist"
	"google.golang.org/grpc"
)

// ServiceClientFactory is responsible for building instances of gRPC service clients
type ServiceClientFactory struct{}

// MakeProductsListServiceClient is responsible for build a ProductsListServiceClient instance
func (scf *ServiceClientFactory) MakeProductsListServiceClient(c container.Container) (productslist.ProductsListServiceClient, error) {
	conn, ok := c.Get(container.ProductsListServiceConnKey)
	if !ok {
		newConn, err := grpc.Dial(c.GetAppConfig().GrpcClientConfig.ProductsListConfig.GetURI(), grpc.WithInsecure())
		if err != nil {
			return nil, errors.Wrap(err, "did not connect to the gRPC server")
		}
		conn = newConn
		c.Put(container.ProductsListServiceConnKey, conn)
	}
	return productslist.NewProductsListServiceClient(conn.(grpc.ClientConnInterface)), nil
}
