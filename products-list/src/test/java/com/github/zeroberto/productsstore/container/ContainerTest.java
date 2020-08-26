package com.github.zeroberto.productsstore.container;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertNotNull;

class ContainerTest {

  private static final String APP_CONFIG_FILE_NAME = "config.yml";

  private Container container;

  @BeforeEach
  void setUp() {
    container = new Container(APP_CONFIG_FILE_NAME);
  }

  @Test
  void givenGetAppConfig() {
    assertNotNull(container.getAppConfig());
  }

  @Test
  void givenGetProductMongoClientBuilder_whenFirstCall_thenNotNull() {
    assertNotNull(container.getProductMongoClientBuilder());
  }

  @Test
  void givenGetProductMongoClientBuilder_whenSecondCall_thenNotNull() {
    container.getDiscountGrpcClientBuilder();

    assertNotNull(container.getProductMongoClientBuilder());
  }

  @Test
  void givenGetDiscountGrpcClientBuilder_whenFirstCall_thenNotNull() {
    assertNotNull(container.getDiscountGrpcClientBuilder());
  }

  @Test
  void givenGetDiscountGrpcClientBuilder_whenSecondtCall_thenNotNull() {
    container.getDiscountGrpcClientBuilder();

    assertNotNull(container.getDiscountGrpcClientBuilder());
  }
}
