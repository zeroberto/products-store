package com.github.zeroberto.productsstore.container.datastorefactory;

import com.github.zeroberto.productsstore.container.Container;
import com.github.zeroberto.productsstore.datastore.DiscountDataStore;
import com.github.zeroberto.productsstore.datastore.network.DiscountDataStoreGrpc;
import lombok.NoArgsConstructor;

import javax.annotation.Nonnull;

import static lombok.AccessLevel.PRIVATE;

@NoArgsConstructor(access = PRIVATE)
public final class DiscountDataStoreFactory {

  @Nonnull
  public static DiscountDataStore makeDiscountDataStore(final Container container) {
    return new DiscountDataStoreGrpc(container.getDiscountGrpcClientBuilder());
  }
}
