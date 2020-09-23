package clientfactory

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/zeroberto/products-store/discount-calculator/container"
	"github.com/zeroberto/products-store/discount-calculator/pb/userinfo"
	"google.golang.org/grpc"
)

// ServiceClientFactory is responsible for building instances of gRPC service clients
type ServiceClientFactory struct{}

// MakeUserInfoServiceClient is responsible for build a UserInfoServiceClient instance
func (scf *ServiceClientFactory) MakeUserInfoServiceClient(c container.Container, host string, port uint) (userinfo.UserInfoClient, error) {
	conn, ok := c.Get(container.UserServiceConnKey)
	if !ok {
		uri := fmt.Sprintf("%s:%d", host, port)

		newConn, err := grpc.Dial(uri, grpc.WithInsecure())
		if err != nil {
			return nil, errors.Wrap(err, "did not connect to the gRPC server")
		}

		conn = newConn
		c.Put(container.UserServiceConnKey, conn)
	}
	return userinfo.NewUserInfoClient(conn.(grpc.ClientConnInterface)), nil
}
