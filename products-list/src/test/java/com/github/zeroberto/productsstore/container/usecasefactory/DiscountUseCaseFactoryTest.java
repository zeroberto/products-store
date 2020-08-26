package com.github.zeroberto.productsstore.container.usecasefactory;

import com.github.zeroberto.productsstore.config.AppConfig;
import com.github.zeroberto.productsstore.container.Container;
import com.github.zeroberto.productsstore.usecase.DiscountUseCase;
import com.github.zeroberto.productsstore.usecase.impl.DiscountUseCaseImpl;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static com.github.zeroberto.productsstore.container.configfactory.AppConfigFactory.makeAppConfig;
import static com.github.zeroberto.productsstore.container.usecasefactory.DiscountUseCaseFactory.makeDiscountUseCase;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertTrue;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

class DiscountUseCaseFactoryTest {

  private Container container;

  @BeforeEach
  void setUp() {
    final AppConfig appConfig = makeAppConfig("config.yml");

    container = mock(Container.class);

    when(container.getAppConfig()).thenReturn(appConfig);
  }

  @Test
  void givenMakeDiscountUseCase_thenNotNull() {
    final DiscountUseCase got = makeDiscountUseCase(container);

    assertNotNull(got);
  }

  @Test
  void givenMakeDiscountUseCase_thenReturnDiscountUseCaseImpl() {
    final DiscountUseCase got = makeDiscountUseCase(container);

    assertTrue(got instanceof DiscountUseCaseImpl);
  }
}
