package com.github.zeroberto.productsstore.cmd;

import com.github.zeroberto.productsstore.cmd.grpcserver.ProductsListGrpcServer;
import com.github.zeroberto.productsstore.cmd.grpcservice.ProductsListGrpcService;
import com.github.zeroberto.productsstore.container.Container;
import lombok.NoArgsConstructor;
import lombok.extern.java.Log;

import static java.lang.String.format;
import static java.util.Objects.nonNull;
import static java.util.logging.Level.SEVERE;
import static lombok.AccessLevel.PRIVATE;

@Log
@NoArgsConstructor(access = PRIVATE)
public final class Server {

  private static final String APP_CONFIG_FILE_NAME = "application%s.yml";

  public static void start(final String profile) {
    try {
      final Container container = new Container(
        format(APP_CONFIG_FILE_NAME, capitalizeProfile(profile)));

      final ProductsListGrpcService productsListGrpcService = new ProductsListGrpcService(container);

      final ProductsListGrpcServer productsListGrpcServer = new ProductsListGrpcServer(
        container.getAppConfig().getGrpcServerConfig().getProductConfig(),
        productsListGrpcService
      );

      productsListGrpcServer.start();
      productsListGrpcServer.blockUntilShutdown();
    } catch (Exception e) {
      log.log(SEVERE, format("An exception occurred when trying to start the server: %s", e.getMessage()), e);
    }
  }

  private static String capitalizeProfile(final String profileName) {
    if (nonNull(profileName) && !profileName.isBlank()) {
      return profileName.substring(0, 1).toUpperCase() + profileName.substring(1);
    }
    return "";
  }
}
