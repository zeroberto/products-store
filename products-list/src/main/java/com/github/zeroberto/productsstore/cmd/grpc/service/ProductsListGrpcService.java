package com.github.zeroberto.productsstore.cmd.grpc.service;

import com.github.zeroberto.productsstore.container.Container;
import com.github.zeroberto.productsstore.exceptions.ProductUseCaseException;
import com.github.zeroberto.productsstore.model.Discount;
import com.github.zeroberto.productsstore.model.DiscountedProduct;
import com.github.zeroberto.productsstore.model.Product;
import com.github.zeroberto.productsstore.productslist.ProductsListGrpc;
import com.github.zeroberto.productsstore.productslist.ProductsListRequest;
import com.github.zeroberto.productsstore.productslist.ProductsListResponse;
import com.github.zeroberto.productsstore.usecase.ProductUseCase;
import io.grpc.stub.StreamObserver;
import lombok.RequiredArgsConstructor;
import lombok.extern.java.Log;

import java.util.List;

import static com.github.zeroberto.productsstore.container.usecasefactory.ProductUseCaseFactory.makeProductUseCase;
import static java.util.Objects.nonNull;
import static java.util.logging.Level.SEVERE;

@Log
@RequiredArgsConstructor
public class ProductsListGrpcService extends ProductsListGrpc.ProductsListImplBase {

  private final Container container;

  @Override
  public void listProducts(
    final ProductsListRequest request,
    final StreamObserver<ProductsListResponse> responseObserver
  ) {
    final ProductUseCase productUseCase = makeProductUseCase(container);

    try {
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
    } catch (ProductUseCaseException e) {
      log.log(SEVERE, e.getMessage(), e);
      responseObserver.onError(e);
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
    final var builder = ProductsListResponse.newBuilder()
      .setId(discountedProduct.getProduct().getId())
      .setPriceInCents(discountedProduct.getProduct().getPriceInCents())
      .setTilte(discountedProduct.getProduct().getTitle())
      .setDescription(discountedProduct.getProduct().getDescription());

    if (nonNull(discountedProduct.getDiscount())) {
      builder.setDiscount(toDiscount(discountedProduct.getDiscount()));
    }

    return builder.build();
  }

  private ProductsListResponse.Discount toDiscount(final Discount discount) {
    return ProductsListResponse.Discount.newBuilder()
      .setPct(discount.getPct())
      .setValueInCents(discount.getValueInCents())
      .build();
  }

  private interface ProductListResponseMapper<T> {
    ProductsListResponse map(T from);
  }
}
