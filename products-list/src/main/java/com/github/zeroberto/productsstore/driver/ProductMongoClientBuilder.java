package com.github.zeroberto.productsstore.driver;

import com.mongodb.MongoClient;
import lombok.AllArgsConstructor;

@AllArgsConstructor
public final class ProductMongoClientBuilder {

  private String host;

  public ProductMongoClientBuilder host(final String host) {
    this.host = host;
    return this;
  }

  public MongoClient build() {
    return new MongoClient(host);
  }
}
