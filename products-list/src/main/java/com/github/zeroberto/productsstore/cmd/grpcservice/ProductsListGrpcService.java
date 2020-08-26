package com.github.zeroberto.productsstore.cmd.grpcservice;

import com.github.zeroberto.productsstore.container.Container;
import com.github.zeroberto.productsstore.model.DiscountedProduct;
import com.github.zeroberto.productsstore.model.Product;
import com.github.zeroberto.productsstore.productslist.ProductsListRequest;
import com.github.zeroberto.productsstore.productslist.ProductsListResponse;
import com.github.zeroberto.productsstore.productslist.ProductsListServiceGrpc;
import com.github.zeroberto.productsstore.usecase.ProductUseCase;
import io.grpc.stub.StreamObserver;
import lombok.RequiredArgsConstructor;
import lombok.extern.java.Log;

import java.util.List;

import static com.github.zeroberto.productsstore.container.usecasefactory.ProductUseCaseFactory.makeProductUseCase;

@Log
@RequiredArgsConstructor
public class ProductsListGrpcService extends ProductsListServiceGrpc.ProductsListServiceImplBase {

  private final Container container;

  @Override
  public void listProducts(
    final ProductsListRequest request,
    final StreamObserver<ProductsListResponse> responseObserver
  ) {
    final ProductUseCase productUseCase = makeProductUseCase(container);

    if (request.getUserId() != 0) {
      processRequest(
        productUseCase.listDiscountedProductsByUser(request.getUserId()),
        responseObserver,
        this::toProductsListResponse
      );
    } else {
      processRequest(
        productUseCase.listProducts(),
        responseObserver,
        this::toProductsListResponse
      );
    }
  }

  private <T> void processRequest(
    final List<T> products,
    final StreamObserver<ProductsListResponse> responseObserver,
    final ProductListResponseMapper<T> mapper
  ) {
    products.forEach(p -> responseObserver.onNext(mapper.map(p)));
    responseObserver.onCompleted();
  }

  private ProductsListResponse toProductsListResponse(final Product product) {
    return ProductsListResponse.newBuilder()
      .setId(product.getId())
      .setPriceInCents(product.getPriceInCents())
      .setTilte(product.getTitle())
      .setDescription(product.getDescription())
      .build();
  }

  private ProductsListResponse toProductsListResponse(final DiscountedProduct discountedProduct) {
    return ProductsListResponse.newBuilder()
      .setId(discountedProduct.getProduct().getId())
      .setPriceInCents(discountedProduct.getProduct().getPriceInCents())
      .setTilte(discountedProduct.getProduct().getTitle())
      .setDescription(discountedProduct.getProduct().getDescription())
      .setDiscount(
        ProductsListResponse.Discount.newBuilder()
          .setPct(discountedProduct.getDiscount().getPct())
          .setValueInCents(discountedProduct.getDiscount().getValueInCents()))
      .build();
  }

  private interface ProductListResponseMapper<T> {
    ProductsListResponse map(T from);
  }
}
