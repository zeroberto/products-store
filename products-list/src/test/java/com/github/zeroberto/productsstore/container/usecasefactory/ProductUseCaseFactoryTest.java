package com.github.zeroberto.productsstore.container.usecasefactory;

import com.github.zeroberto.productsstore.config.AppConfig;
import com.github.zeroberto.productsstore.container.Container;
import com.github.zeroberto.productsstore.usecase.ProductUseCase;
import com.github.zeroberto.productsstore.usecase.impl.ProductUseCaseImpl;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static com.github.zeroberto.productsstore.container.configfactory.AppConfigFactory.makeAppConfig;
import static com.github.zeroberto.productsstore.container.usecasefactory.ProductUseCaseFactory.makeProductUseCase;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertTrue;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

class ProductUseCaseFactoryTest {

  private Container container;

  @BeforeEach
  void setUp() {
    final AppConfig appConfig = makeAppConfig("config.yml");

    container = mock(Container.class);

    when(container.getAppConfig()).thenReturn(appConfig);
  }

  @Test
  void givenMakeProductUseCase_thenNotNull() {
    final ProductUseCase got = makeProductUseCase(container);

    assertNotNull(got);
  }

  @Test
  void givenMakeProductUseCase_thenReturnProductUseCaseImpl() {
    final ProductUseCase got = makeProductUseCase(container);

    assertTrue(got instanceof ProductUseCaseImpl);
  }
}
