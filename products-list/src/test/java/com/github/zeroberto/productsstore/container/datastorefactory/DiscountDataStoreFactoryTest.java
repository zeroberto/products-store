package com.github.zeroberto.productsstore.container.datastorefactory;

import com.github.zeroberto.productsstore.config.AppConfig;
import com.github.zeroberto.productsstore.container.Container;
import com.github.zeroberto.productsstore.datastore.DiscountDataStore;
import com.github.zeroberto.productsstore.datastore.network.DiscountDataStoreGrpc;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static com.github.zeroberto.productsstore.container.configfactory.AppConfigFactory.makeAppConfig;
import static com.github.zeroberto.productsstore.container.datastorefactory.DiscountDataStoreFactory.makeDiscountDataStore;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertTrue;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

class DiscountDataStoreFactoryTest {

  private Container container;

  @BeforeEach
  void setUp() {
    final AppConfig appConfig = makeAppConfig("config.yml");

    container = mock(Container.class);

    when(container.getAppConfig()).thenReturn(appConfig);
  }

  @Test
  void givenMakeDiscountDataStore_thenNotNull() {
    final DiscountDataStore got = makeDiscountDataStore(container);

    assertNotNull(got);
  }

  @Test
  void givenMakeDiscountDataStore_thenReturnDiscountDataStoreGrpc() {
    final DiscountDataStore got = makeDiscountDataStore(container);

    assertTrue(got instanceof DiscountDataStoreGrpc);
  }
}
