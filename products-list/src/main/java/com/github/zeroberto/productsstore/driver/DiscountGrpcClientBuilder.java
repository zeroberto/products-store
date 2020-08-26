package com.github.zeroberto.productsstore.driver;

import com.github.zeroberto.productsstore.cmd.grpcclient.DiscountGrpcClient;
import lombok.AllArgsConstructor;

@AllArgsConstructor
public final class DiscountGrpcClientBuilder {

  private String host;

  public DiscountGrpcClientBuilder host(final String host) {
    this.host = host;
    return this;
  }

  public DiscountGrpcClient build() {
    return new DiscountGrpcClient(host);
  }
}
