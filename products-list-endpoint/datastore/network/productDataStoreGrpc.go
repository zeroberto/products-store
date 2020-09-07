package network

import (
	"context"
	"io"
	"time"

	"github.com/pkg/errors"
	"github.com/zeroberto/products-store/products-list-endpoint/model"
	"github.com/zeroberto/products-store/products-list-endpoint/pb/productslist"
)

// ProductDataStoreGrpc is responsible for obtaining product data
// through a gRPC client
type ProductDataStoreGrpc struct {
	Client productslist.ProductsListClient
}

// FindAllWithDiscountByUserID is responsible for returning a list of discounted
// products per user from a given repository
func (pds *ProductDataStoreGrpc) FindAllWithDiscountByUserID(userID int64) ([]model.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	plreq := &productslist.ProductsListRequest{}
	plreq.UserId = 1

	stream, err := pds.Client.ListProducts(ctx, plreq)
	if err != nil {
		return nil, errors.Wrap(err, "Could not get product list")
	}

	products := []model.Product{}

	for {
		plresp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.Wrap(err, "Could not get product list")
		}
		products = append(products, toProduct(plresp))
	}
	return products, nil
}

func toProduct(plresp *productslist.ProductsListResponse) model.Product {
	return model.Product{
		ID:           plresp.GetId(),
		Title:        plresp.GetTilte(),
		Description:  plresp.GetDescription(),
		PriceInCents: plresp.GetPriceInCents(),
		Discount:     toDiscount(plresp.GetDiscount()),
	}
}

func toDiscount(d *productslist.ProductsListResponse_Discount) *model.Discount {
	if d != nil {
		return &model.Discount{
			Percentage:   d.GetPct(),
			ValueInCents: d.GetValueInCents(),
		}
	}
	return nil
}
