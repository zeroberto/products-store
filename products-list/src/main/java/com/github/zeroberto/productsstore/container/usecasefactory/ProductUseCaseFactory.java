package com.github.zeroberto.productsstore.container.usecasefactory;

import com.github.zeroberto.productsstore.container.Container;
import com.github.zeroberto.productsstore.datastore.ProductDataStore;
import com.github.zeroberto.productsstore.usecase.DiscountUseCase;
import com.github.zeroberto.productsstore.usecase.ProductUseCase;
import com.github.zeroberto.productsstore.usecase.impl.ProductUseCaseImpl;
import lombok.NoArgsConstructor;

import javax.annotation.Nonnull;

import static com.github.zeroberto.productsstore.container.datastorefactory.ProductDataStoreFactory.makeProductDataStore;
import static com.github.zeroberto.productsstore.container.usecasefactory.DiscountUseCaseFactory.makeDiscountUseCase;
import static lombok.AccessLevel.PRIVATE;

@NoArgsConstructor(access = PRIVATE)
public final class ProductUseCaseFactory {

  @Nonnull
  public static ProductUseCase makeProductUseCase(final Container container) {
    final ProductDataStore productDataStore = makeProductDataStore(container);
    final DiscountUseCase discountUseCase = makeDiscountUseCase(container);
    return new ProductUseCaseImpl(productDataStore, discountUseCase);
  }
}
