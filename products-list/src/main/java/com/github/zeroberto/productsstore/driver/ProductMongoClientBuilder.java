package com.github.zeroberto.productsstore.driver;

import com.github.zeroberto.productsstore.config.AppConfig;
import com.mongodb.MongoClient;
import com.mongodb.MongoClientOptions;
import com.mongodb.MongoCredential;
import com.mongodb.ServerAddress;
import lombok.AllArgsConstructor;

import static com.github.zeroberto.productsstore.utils.AppConfigUtils.getAuthPass;
import static com.github.zeroberto.productsstore.utils.AppConfigUtils.getAuthUser;
import static java.util.Optional.ofNullable;

@AllArgsConstructor
public final class ProductMongoClientBuilder {

  private AppConfig.DSConfig dsConfig;

  public ProductMongoClientBuilder config(final AppConfig.DSConfig dsConfig) {
    this.dsConfig = dsConfig;
    return this;
  }

  public MongoClient build() {
    return new MongoClient(
      new ServerAddress(dsConfig.getHost()),
      getMongoCredentials(),
      MongoClientOptions.builder().build());
  }

  private MongoCredential getMongoCredentials() {
    final String username = getAuthUser(dsConfig.getAuth());
    final char[] password = ofNullable(getAuthPass(dsConfig.getAuth()))
      .map(String::toCharArray)
      .orElse(new char[]{});
    final String repository = dsConfig.getAuth().getRepo();

    return MongoCredential.createCredential(username, repository, password);
  }
}
