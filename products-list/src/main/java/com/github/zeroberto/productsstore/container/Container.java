package com.github.zeroberto.productsstore.container;

import com.github.zeroberto.productsstore.config.AppConfig;
import com.github.zeroberto.productsstore.driver.DiscountGrpcClientBuilder;
import com.github.zeroberto.productsstore.driver.ProductMongoClientBuilder;
import lombok.Getter;

import static com.github.zeroberto.productsstore.container.configfactory.AppConfigFactory.makeAppConfig;
import static java.util.Objects.isNull;

public final class Container {

  @Getter
  private final AppConfig appConfig;
  private ProductMongoClientBuilder productMongoClientBuilder;
  private DiscountGrpcClientBuilder discountGrpcClientBuilder;

  public Container(final String appConfigFileName) {
    appConfig = makeAppConfig(appConfigFileName);
  }

  public ProductMongoClientBuilder getProductMongoClientBuilder() {
    if (isNull(productMongoClientBuilder)) {
      productMongoClientBuilder = new ProductMongoClientBuilder(appConfig
        .getUseCaseConfig()
        .getProductConfig()
        .getDsConfig()
        .getHost());
    }
    return productMongoClientBuilder;
  }

  public DiscountGrpcClientBuilder getDiscountGrpcClientBuilder() {
    if (isNull(discountGrpcClientBuilder)) {
      discountGrpcClientBuilder = new DiscountGrpcClientBuilder(appConfig
        .getUseCaseConfig()
        .getDiscountConfig()
        .getDsConfig()
        .getHost());
    }
    return discountGrpcClientBuilder;
  }
}
