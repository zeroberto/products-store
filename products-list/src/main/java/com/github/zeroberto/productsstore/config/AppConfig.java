package com.github.zeroberto.productsstore.config;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import static lombok.AccessLevel.PRIVATE;

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor(access = PRIVATE)
@Builder
public class AppConfig {

  private GrpcServerConfig grpcServerConfig;
  private UseCaseConfig useCaseConfig;

  @Getter
  @Setter
  @NoArgsConstructor
  public static class UseCaseConfig {
    private ProductUseCaseConfig productConfig;
    private DiscountUseCaseConfig discountConfig;
  }

  @Getter
  @Setter
  @NoArgsConstructor
  public static class GrpcServerConfig {
    private ServerConfig productConfig;
  }

  @Getter
  @Setter
  @NoArgsConstructor
  public static class ProductUseCaseConfig {
    private DSConfig dsConfig;
  }

  @Getter
  @Setter
  @NoArgsConstructor
  public static class DiscountUseCaseConfig {
    private DSConfig dsConfig;
  }

  @Getter
  @Setter
  @NoArgsConstructor
  public static class DSConfig {
    private String type;
    private String host;
    private String database;
  }

  @Getter
  @Setter
  @NoArgsConstructor
  public static class ServerConfig {
    private int port;
  }
}
