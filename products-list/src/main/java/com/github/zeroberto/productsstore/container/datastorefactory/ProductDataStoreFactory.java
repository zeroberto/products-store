package com.github.zeroberto.productsstore.container.datastorefactory;

import com.github.zeroberto.productsstore.container.Container;
import com.github.zeroberto.productsstore.datastore.ProductDataStore;
import com.github.zeroberto.productsstore.datastore.database.ProductDataStoreMongoDB;
import lombok.NoArgsConstructor;

import javax.annotation.Nonnull;

import static lombok.AccessLevel.PRIVATE;

@NoArgsConstructor(access = PRIVATE)
public final class ProductDataStoreFactory {

  @Nonnull
  public static ProductDataStore makeProductDataStore(final Container container) {
    return new ProductDataStoreMongoDB(
      container.getProductMongoClientBuilder(),
      container
        .getAppConfig()
        .getUseCaseConfig()
        .getProductConfig()
        .getDsConfig()
        .getDatabase());
  }
}
