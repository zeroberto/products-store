package com.github.zeroberto.productsstore.usecase.impl;

import com.github.zeroberto.productsstore.datastore.ProductDataStore;
import com.github.zeroberto.productsstore.model.DiscountedProduct;
import com.github.zeroberto.productsstore.model.Product;
import com.github.zeroberto.productsstore.usecase.DiscountUseCase;
import com.github.zeroberto.productsstore.usecase.ProductUseCase;
import lombok.RequiredArgsConstructor;

import java.util.List;

import static java.util.stream.Collectors.toList;

@RequiredArgsConstructor
public class ProductUseCaseImpl implements ProductUseCase {

  private final ProductDataStore productDataStore;
  private final DiscountUseCase discountUseCase;

  @Override
  public List<Product> listProducts() {
    return productDataStore.findAll();
  }

  @Override
  public List<DiscountedProduct> listDiscountedProductsByUser(long userId) {
    return productDataStore.findAll()
      .stream()
      .map(product -> DiscountedProduct.builder()
        .product(product)
        .discount(discountUseCase.calculateDiscount(product.getId(), userId))
        .build())
      .collect(toList());
  }
}
