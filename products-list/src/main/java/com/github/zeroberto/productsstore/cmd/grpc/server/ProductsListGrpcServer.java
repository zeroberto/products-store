package com.github.zeroberto.productsstore.cmd.grpc.server;

import com.github.zeroberto.productsstore.cmd.grpc.service.ProductsListGrpcService;
import com.github.zeroberto.productsstore.config.AppConfig;
import io.grpc.Server;
import io.grpc.ServerBuilder;
import lombok.extern.java.Log;

import java.io.IOException;

import static java.lang.String.format;
import static java.util.concurrent.TimeUnit.SECONDS;
import static java.util.logging.Level.SEVERE;

@Log
public final class ProductsListGrpcServer {

  private final AppConfig.ServerConfig serverConfig;
  private final Server server;

  public ProductsListGrpcServer(
    final AppConfig.ServerConfig serverConfig,
    final ProductsListGrpcService productsListGrpcService) {
    this(productsListGrpcService, serverConfig, ServerBuilder.forPort(serverConfig.getPort()));
  }

  public ProductsListGrpcServer(
    final ProductsListGrpcService productsListGrpcService,
    final AppConfig.ServerConfig serverConfig,
    final ServerBuilder<?> serverBuilder
  ) {
    this.serverConfig = serverConfig;
    server = serverBuilder
      .addService(productsListGrpcService)
      .build();
  }

  public void start() throws IOException {
    server.start();
    log.info("Server started, listening on " + serverConfig.getPort());
    Runtime.getRuntime().addShutdownHook(new Thread(() -> {
      log.info("*** shutting down ProductsListGrpcServer server since JVM is shutting down");
      try {
        ProductsListGrpcServer.this.stop();
      } catch (InterruptedException e) {
        log.log(SEVERE, format("m=addShutdownHook, msg=%s", e.getMessage()), e);
        Thread.currentThread().interrupt();
      }
      log.info("*** server shut down");
    }));
  }

  public void blockUntilShutdown() throws InterruptedException {
    if (server != null) {
      server.awaitTermination();
    }
  }

  private void stop() throws InterruptedException {
    if (server != null) {
      server.shutdown().awaitTermination(30, SECONDS);
    }
  }
}
