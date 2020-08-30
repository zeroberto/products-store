package com.github.zeroberto.productsstore.cmd;

import com.github.zeroberto.productsstore.cmd.grpc.server.ProductsListGrpcServer;
import com.github.zeroberto.productsstore.cmd.grpc.service.ProductsListGrpcService;
import com.github.zeroberto.productsstore.container.Container;
import lombok.NoArgsConstructor;
import lombok.extern.java.Log;

import java.io.IOException;

import static java.lang.String.format;
import static java.util.Objects.nonNull;
import static java.util.logging.Level.SEVERE;
import static lombok.AccessLevel.PRIVATE;

@Log
@NoArgsConstructor(access = PRIVATE)
public final class Server {

  private static final String APP_CONFIG_FILE_NAME = "application%s.yml";
  private static final String DEFAULT_PROFILE = "";

  public static Server newServer() {
    return new Server();
  }

  public void start(final String profile) {
    try {
      final String formattedProfile = capitalizeProfile(profile);

      final Container container = new Container(
        format(APP_CONFIG_FILE_NAME, formattedProfile));

      log.info(format("Application starting with the %s profile",
        DEFAULT_PROFILE.equals(formattedProfile) ? "default" : profile
      ));

      startGrpcServer(container);
    } catch (Exception e) {
      log.log(SEVERE, format("An exception occurred when trying to start the server: %s", e.getMessage()), e);
    }
  }

  private void startGrpcServer(final Container container) throws IOException, InterruptedException {
    final ProductsListGrpcService productsListGrpcService = new ProductsListGrpcService(container);

    final ProductsListGrpcServer productsListGrpcServer = new ProductsListGrpcServer(
      container.getAppConfig().getGrpcServerConfig().getProductConfig(),
      productsListGrpcService
    );

    productsListGrpcServer.start();
    productsListGrpcServer.blockUntilShutdown();
  }

  private String capitalizeProfile(final String profileName) {
    if (nonNull(profileName) && !profileName.isBlank()) {
      return profileName.substring(0, 1).toUpperCase() + profileName.substring(1);
    }
    return DEFAULT_PROFILE;
  }
}
