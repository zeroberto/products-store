package com.github.zeroberto.productsstore.container.datastorefactory;

import com.github.zeroberto.productsstore.config.AppConfig;
import com.github.zeroberto.productsstore.container.Container;
import com.github.zeroberto.productsstore.datastore.ProductDataStore;
import com.github.zeroberto.productsstore.datastore.database.ProductDataStoreMongoDB;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static com.github.zeroberto.productsstore.container.configfactory.AppConfigFactory.makeAppConfig;
import static com.github.zeroberto.productsstore.container.datastorefactory.ProductDataStoreFactory.makeProductDataStore;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertTrue;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

class ProductDataStoreFactoryTest {

  private Container container;

  @BeforeEach
  void setUp() {
    final AppConfig appConfig = makeAppConfig("config.yml");

    container = mock(Container.class);

    when(container.getAppConfig()).thenReturn(appConfig);
  }

  @Test
  void givenMakeProductDataStore_thenNotNull() {
    final ProductDataStore got = makeProductDataStore(container);

    assertNotNull(got);
  }

  @Test
  void givenMakeProductDataStore_thenReturnProductDataStoreMongoDB() {
    final ProductDataStore got = makeProductDataStore(container);

    assertTrue(got instanceof ProductDataStoreMongoDB);
  }
}
