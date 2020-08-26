package com.github.zeroberto.productsstore.container.usecasefactory;

import com.github.zeroberto.productsstore.container.Container;
import com.github.zeroberto.productsstore.datastore.DiscountDataStore;
import com.github.zeroberto.productsstore.usecase.DiscountUseCase;
import com.github.zeroberto.productsstore.usecase.impl.DiscountUseCaseImpl;
import lombok.NoArgsConstructor;

import javax.annotation.Nonnull;

import static com.github.zeroberto.productsstore.container.datastorefactory.DiscountDataStoreFactory.makeDiscountDataStore;
import static lombok.AccessLevel.PRIVATE;

@NoArgsConstructor(access = PRIVATE)
public final class DiscountUseCaseFactory {

  @Nonnull
  public static DiscountUseCase makeDiscountUseCase(final Container container) {
    final DiscountDataStore discountDataStore = makeDiscountDataStore(container);
    return new DiscountUseCaseImpl(discountDataStore);
  }
}
