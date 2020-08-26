package com.github.zeroberto.productsstore.usecase;

import com.github.zeroberto.productsstore.model.DiscountedProduct;
import com.github.zeroberto.productsstore.model.Product;

import java.util.List;

public interface ProductUseCase {

  List<Product> listProducts();

  List<DiscountedProduct> listDiscountedProductsByUser(final long userId);
}
