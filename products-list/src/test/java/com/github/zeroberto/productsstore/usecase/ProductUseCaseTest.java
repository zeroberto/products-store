package com.github.zeroberto.productsstore.usecase;

import com.github.zeroberto.productsstore.datastore.ProductDataStore;
import com.github.zeroberto.productsstore.model.Discount;
import com.github.zeroberto.productsstore.model.DiscountedProduct;
import com.github.zeroberto.productsstore.model.Product;
import com.github.zeroberto.productsstore.usecase.impl.ProductUseCaseImpl;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import java.util.List;
import java.util.Optional;

import static org.junit.jupiter.api.Assertions.assertAll;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.mockito.ArgumentMatchers.anyLong;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

class ProductUseCaseTest {

  private ProductUseCase productUseCase;
  private ProductDataStore productDataStore;
  private DiscountUseCase discountUseCase;

  @BeforeEach
  public void setUp() {
    productDataStore = mock(ProductDataStore.class);
    discountUseCase = mock(DiscountUseCase.class);

    productUseCase = new ProductUseCaseImpl(productDataStore, discountUseCase);
  }

  @Test
  void listProductsWithDiscountByUser() {
    final int userId = 1;
    final Product product = Product.builder()
      .id("test1")
      .priceInCents(100)
      .title("test title")
      .description("test description")
      .build();
    final Discount discount = Discount.builder()
      .pct(10)
      .valueInCents(10)
      .build();

    when(productDataStore.findAll()).thenReturn(List.of(product));
    when(discountUseCase.calculateDiscount(anyString(), anyLong())).thenReturn(discount);

    final List<DiscountedProduct> discountedProducts = productUseCase.listDiscountedProductsByUser(userId);

    assert !discountedProducts.isEmpty();

    final Optional<DiscountedProduct> discountedProduct = discountedProducts.stream().findFirst();

    assertAll("testListProductsWithDiscountByUser",
      () -> assertEquals(product, discountedProduct.get().getProduct()),
      () -> assertEquals(discount, discountedProduct.get().getDiscount()));
  }

  @Test
  void listProducts() {
    final Product expected = Product.builder()
      .id("test1")
      .priceInCents(100)
      .title("test title")
      .description("test description")
      .build();

    when(productDataStore.findAll()).thenReturn(List.of(expected));

    final List<Product> products = productUseCase.listProducts();

    assert !products.isEmpty();

    final Optional<Product> got = products.stream().findFirst();

    assertAll("listProducts",
      () -> assertEquals(expected, got.get()));
  }
}
